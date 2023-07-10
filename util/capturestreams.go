// freely adapted by @rudifa Rudolf Farkas
// from the code https://github.com/eliben/code-for-blog/tree/master/2020/go-fake-stdio
// Basic package for faking Stdio.
//
// Eli Bendersky [https://eli.thegreenplace.net]
// This code is in the public domain.
// see also the article https://eli.thegreenplace.net/2020/faking-stdin-and-stdout-in-go/

package util

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

// CaptureStdio can be used to fake stdin and capture stdout.
// Between creating a new CaptureStdio and calling ReadAndRestore on it,
// code reading os.Stdin will get the contents of stdinText passed to New.
// Output to os.Stdout will be captured and returned from ReadAndRestore.
// CaptureStdio is not reusable; don't attempt to use it after calling
// ReadAndRestore, but it should be safe to create a new CaptureStdio.
type CaptureStdio struct {
	origStdout *os.File
	origStderr *os.File

	stdoutCh chan []byte
	stderrCh chan []byte

	stdoutReader *os.File
	stderrReader *os.File
	origStdin    *os.File
	stdinWriter  *os.File
}

func New(stdinText string) (*CaptureStdio, error) {
	// Pipe for stdin.
	//
	//                 ======
	//  w ------------->||||------> r
	// (stdinWriter)   ======      (os.Stdin)
	stdinReader, stdinWriter, err := os.Pipe()
	if err != nil {
		return nil, err
	}

	// Pipe for stdout.
	//
	//               ======
	//  w ----------->||||------> r
	// (os.Stdout)   ======      (stdoutReader)
	stdoutReader, stdoutWriter, err := os.Pipe()
	if err != nil {
		return nil, err
	}

	// Pipe for stderr.
	//
	//               ======
	//  w ----------->||||------> r
	// (os.Stderr)   ======      (stderrReader)
	stderrReader, stderrWriter, err := os.Pipe()
	if err != nil {
		return nil, err
	}

	origStdin := os.Stdin
	os.Stdin = stdinReader

	_, err = stdinWriter.Write([]byte(stdinText))
	if err != nil {
		stdinWriter.Close()
		os.Stdin = origStdin
		return nil, err
	}

	origStdout := os.Stdout
	os.Stdout = stdoutWriter

	stdoutCh := make(chan []byte)

	// This goroutine reads stdout into a buffer in the background.
	go func() {
		var b bytes.Buffer
		if _, err := io.Copy(&b, stdoutReader); err != nil {
			log.Println(err)
		}
		stdoutCh <- b.Bytes()
	}()

	origStderr := os.Stderr
	os.Stderr = stderrWriter

	stderrCh := make(chan []byte)

	// This goroutine reads stderr into a buffer in the background.
	go func() {
		var b bytes.Buffer
		if _, err := io.Copy(&b, stderrReader); err != nil {
			log.Println(err)
		}
		stderrCh <- b.Bytes()
	}()

	return &CaptureStdio{
		origStdout:   origStdout,
		origStderr:   origStderr,
		stdoutReader: stdoutReader,
		stderrReader: stderrReader,
		stdoutCh:     stdoutCh,
		stderrCh:     stderrCh,
		origStdin:    origStdin,
		stdinWriter:  stdinWriter,
	}, nil
}

// CloseStdin closes the fake stdin. This may be necessary if the process has
// logic for reading stdin until EOF; otherwise such code would block forever.
func (sf *CaptureStdio) CloseStdin() {
	if sf.stdinWriter != nil {
		sf.stdinWriter.Close()
		sf.stdinWriter = nil
	}
}

// ReadAndRestore collects all captured stdout and returns it; it also restores
// os.Stdin and os.Stdout to their original values.
func (sf *CaptureStdio) ReadAndRestore() ([]byte, []byte, error) {
	if sf.stdoutReader == nil {
		return nil, nil, fmt.Errorf("ReadAndRestore from closed FakeStdio stdoutReader")
	}
	if sf.stderrReader == nil {
		return nil, nil, fmt.Errorf("ReadAndRestore from closed FakeStdio stderrReader")
	}

	// Close the writer side of the faked stdout pipe. This signals to the
	// background goroutine that it should exit.
	os.Stdout.Close()
	stdout := <-sf.stdoutCh

	os.Stderr.Close()
	stderr := <-sf.stderrCh

	os.Stdout = sf.origStdout
	os.Stderr = sf.origStderr
	os.Stdin = sf.origStdin

	if sf.stdoutReader != nil {
		sf.stdoutReader.Close()
		sf.stdoutReader = nil
	}

	if sf.stderrReader != nil {
		sf.stderrReader.Close()
		sf.stderrReader = nil
	}

	if sf.stdinWriter != nil {
		sf.stdinWriter.Close()
		sf.stdinWriter = nil
	}

	return stdout, stderr, nil
}

// ExecuteFnAndCaptureStdoutAndStderr runs fn with its args, captures its stdout and stder, and returns them as strings
func ExecuteFnAndCaptureStdoutAndStderr(fn func(string, ...string) error, cmd string, args ...string) (string, string, error) {

	fs, err := New("")
	if err != nil {
		return "", "", err
	}

	// Call the provided function
	err = fn(cmd, args...)

	stdout, stderr, err := fs.ReadAndRestore()
	if err != nil {
		return "", "", err
	}

	return string(stdout), string(stderr), err
}
