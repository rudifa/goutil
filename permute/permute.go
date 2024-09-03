/*
Copyright © 2024 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/

package permute

import (
	"sort"
)

// Permute generates all permutations of nums
func Permute(nums []int) [][]int {
	var result [][]int
	var backtrack func(int)

	backtrack = func(start int) {
		if start == len(nums) {
			perm := make([]int, len(nums))
			copy(perm, nums)
			result = append(result, perm)
			return
		}

		for i := start; i < len(nums); i++ {
			nums[start], nums[i] = nums[i], nums[start]
			backtrack(start + 1)
			nums[start], nums[i] = nums[i], nums[start] // backtrack
		}
	}

	backtrack(0)
	return result
}

// PermuteExcept generates all permutations of nums except the values in `except`
func PermuteExcept(nums, except []int) [][]int {
	// Create a map of indices to exclude
	excludeMap := make(map[int]bool)
	for _, v := range except {
		for i, num := range nums {
			if num == v {
				excludeMap[i] = true
				break
			}
		}
	}

	// Create a slice of indices to permute
	var toPermute []int
	for i := range nums {
		if !excludeMap[i] {
			toPermute = append(toPermute, i)
		}
	}

	var result [][]int
	var backtrack func(int)

	backtrack = func(start int) {
		if start == len(toPermute) {
			perm := make([]int, len(nums))
			copy(perm, nums)
			result = append(result, perm)
			return
		}

		for i := start; i < len(toPermute); i++ {
			idx1, idx2 := toPermute[start], toPermute[i]
			nums[idx1], nums[idx2] = nums[idx2], nums[idx1]
			backtrack(start + 1)
			nums[idx1], nums[idx2] = nums[idx2], nums[idx1] // backtrack
		}
	}

	backtrack(0)
	return result
}

func RangeOf(n int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i
	}
	return nums
}

func SortPermutations(permutations [][]int) {
	sort.Slice(permutations, func(i, j int) bool {
		for x := range permutations[i] {
			if permutations[i][x] != permutations[j][x] {
				return permutations[i][x] < permutations[j][x]
			}
		}
		return false
	})
}
