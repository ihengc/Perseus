package BinarySearch

/********************************************************
* @author: Ihc
* @date: 2022/4/26 0026 12:18
* @version: 1.0
* @description:
*********************************************************/

// BinarySearch 二分查找
// 查询dst在nums中的位置,若未查到,则返回要插入的位置
func BinarySearch(dst int, nums []int) int {
	rightIndex := len(nums) - 1
	if rightIndex < 0 {
		return 0
	}
	leftIndex := 0
	midIndex := 0
	for leftIndex <= rightIndex {
		midIndex = (rightIndex + leftIndex) / 2
		if dst == nums[midIndex] {
			return midIndex
		} else if dst > nums[midIndex] {
			leftIndex = midIndex + 1
		} else {
			rightIndex = midIndex - 1
		}
	}
	if dst > nums[midIndex] {
		return midIndex + 1
	}
	return midIndex
}
