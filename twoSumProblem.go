package main

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		targetLeft := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == targetLeft {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}
