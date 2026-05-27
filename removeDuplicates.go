package main

func removeDuplicates(nums []int) int {
	var cur int = 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			cur++
			nums[cur] = nums[i]
		}
	}
	return cur + 1
}
