/*
Copyright © 2024 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/

package permute_test

import (
	"reflect"
	"testing"

	"github.com/rudifa/goutil/permute"
)

func TestPermute(t *testing.T) {

	expected := [][]int{
		{0, 1, 2, 3},
		{0, 1, 3, 2},
		{0, 2, 1, 3},
		{0, 2, 3, 1},
		{0, 3, 2, 1},
		{0, 3, 1, 2},
		{1, 0, 2, 3},
		{1, 0, 3, 2},
		{1, 2, 0, 3},
		{1, 2, 3, 0},
		{1, 3, 2, 0},
		{1, 3, 0, 2},
		{2, 1, 0, 3},
		{2, 1, 3, 0},
		{2, 0, 1, 3},
		{2, 0, 3, 1},
		{2, 3, 0, 1},
		{2, 3, 1, 0},
		{3, 1, 2, 0},
		{3, 1, 0, 2},
		{3, 2, 1, 0},
		{3, 2, 0, 1},
		{3, 0, 2, 1},
		{3, 0, 1, 2},
	}

	result := permute.Permute(permute.RangeOf(4))
	permute.SortPermutations(result)
	permute.SortPermutations(expected)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("permute4(4) = %v; want %v", result, expected)
	}
}

func TestPermuteExcept(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	excluding := []int{2, 4}
	expected := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 5, 4, 3},
		{3, 2, 1, 4, 5},
		{3, 2, 5, 4, 1},
		{5, 2, 1, 4, 3},
		{5, 2, 3, 4, 1},
	}

	result := permute.PermuteExcept(nums, excluding)
	permute.SortPermutations(result)
	permute.SortPermutations(expected)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("permutationsOf(%v, %v) = %v; want %v", nums, excluding, result, expected)
	}
}

func TestPermuteExcept2(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	excluding := []int{}
	expected := [][]int{
		{1, 2, 3, 4},
		{1, 2, 4, 3},
		{1, 3, 2, 4},
		{1, 3, 4, 2},
		{1, 4, 3, 2},
		{1, 4, 2, 3},
		{2, 1, 3, 4},
		{2, 1, 4, 3},
		{2, 3, 1, 4},
		{2, 3, 4, 1},
		{2, 4, 3, 1},
		{2, 4, 1, 3},
		{3, 2, 1, 4},
		{3, 2, 4, 1},
		{3, 1, 2, 4},
		{3, 1, 4, 2},
		{3, 4, 1, 2},
		{3, 4, 2, 1},
		{4, 2, 3, 1},
		{4, 2, 1, 3},
		{4, 3, 2, 1},
		{4, 3, 1, 2},
		{4, 1, 3, 2},
		{4, 1, 2, 3},
	}

	result := permute.PermuteExcept(nums, excluding)
	permute.SortPermutations(result)
	permute.SortPermutations(expected)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("permutationsOf(%v, %v) = %v; want %v", nums, excluding, result, expected)
	}
}
