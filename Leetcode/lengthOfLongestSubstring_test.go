package Leetcode

import "testing"

/********************************************************
* @author: Ihc
* @date: 2022/4/24 0024 16:15
* @version: 1.0
* @description:
*********************************************************/

func Test_lengthOfLongestSubstring(t *testing.T) {
	size := lengthOfLongestSubstring("dvdf") // wpwake //dvd
	if size != 3 {
		t.Fatal("dvdf")
	}
	size = lengthOfLongestSubstring("wpwake")
	if size != 5 {
		t.Fatal("wpwake")
	}
	size = lengthOfLongestSubstring("abcabcbbb")
	if size != 3 {
		t.Fatal("abcabcbbb")
	}
	size = lengthOfLongestSubstring(" ")
	if size != 1 {
		t.Fatal(" ")
	}
	size = lengthOfLongestSubstring("")
	if size != 0 {
		t.Fatal("")
	}
}
