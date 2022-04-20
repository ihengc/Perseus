package SkipList

import "testing"

/********************************************************
* @author: Ihc
* @date: 2022/4/20 0020 14:22
* @version: 1.0
* @description:
*********************************************************/

func TestBuiltinTypeComparator(t *testing.T) {
	if BuiltinTypeComparator(1, 1) != 0 {
		t.Fatal("BuiltinTypeComparator error")
	}
}
