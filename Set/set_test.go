package Set

import "testing"

/********************************************************
* @author: Ihc
* @date: 2022/4/21 0021 09:46
* @version: 1.0
* @description:
*********************************************************/

var s ISet = NewSet()

func TestSet(t *testing.T) {
	s.Add(1)
	if s.Len() != 1 {
		t.Fatal("len err")
	}
	if s.In(1) != true {
		t.Fatal("in err")
	}
	s.Del(1)
	if s.Len() != 0 {
		t.Fatal("del err")
	}
}
