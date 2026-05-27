package main

func removeElement(nums []int, val int) int {
	var curIdx = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			if curIdx != i {
				nums[curIdx] = nums[i]
			}
			curIdx++
		}
	}
	return curIdx
}
