package Leetcode

/********************************************************
* @author: Ihc
* @date: 2022/4/24 0024 14:51
* @version: 1.0
* @description:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

*********************************************************/

// lengthOfLongestSubstring 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	var (
		start = 0 // 窗口起始位置
		cMap  = make(map[byte]int)
		size  = 0 // 窗口大小
	)
	// end 窗口末尾位置
	for end := 0; end < len(s); end++ {
		// 当前窗口中出现重复字符,将第一个重复字符向前移动一位,缩小窗口
		if k, ok := cMap[s[end]]; ok {
			if k >= start {
				start = k + 1
			}
		}
		cMap[s[end]] = end
		size = max(end-start+1, size)
	}
	return size
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
