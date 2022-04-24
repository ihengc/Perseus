package Leetcode

/********************************************************
* @author: Ihc
* @date: 2022/4/24 0024 11:49
* @version: 1.0
* @description:
*********************************************************/

// TowSum 两数之和
func TowSum(nums []int, target int) []int {
	var (
		// k存放target-nums[i]
		m map[int]int
		r []int
	)
	for i := 0; i < len(nums); i++ {
		if index, ok := m[nums[i]]; ok {
			return []int{r[nums[i]], index}
		} else {
			m[target-nums[i]] = i
		}
	}
	return r
}
