package util

var utilVerbose bool = false

// SetVerbose sets the verbose flag
func SetVerbose(verbose bool) {
	utilVerbose = verbose
}

// Verbose returns the verbose flag
func Verbose() bool {
	return utilVerbose
}
