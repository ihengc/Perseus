package SkipList

import (
	"sync"
	"testing"
)

/********************************************************
* @author: Ihc
* @date: 2022/4/20 0020 11:14
* @version: 1.0
* @description:
*********************************************************/

func OptionFunc(option *Options) {
	option.probability = 0.5
	option.locker = &sync.RWMutex{}
	option.comparator = BuiltinTypeComparator
	option.maxLevel = 16
}

var skipList = NewSkipListWithOption(OptionFunc)

func TestSkipList(t *testing.T) {
	if skipList.Len() != 0 {
		t.Fatal("len error")
	}
	for i := 0; i < 10; i++ {
		skipList.Put(i, i+1)
		if skipList.Len() != i+1 {
			t.Fatal("put error")
		}
	}
	if skipList.Len() != 10 {
		t.Fatal("len error")
	}
	if skipList.Get(0) != 1 {
		t.Fatal("get error")
	}
	if skipList.Get(9) != 10 {
		t.Fatal("get error")
	}
	if skipList.Del(2) != true {
		t.Fatal("del error")
	}
	if skipList.Len() != 9 {
		t.Fatal("len error")
	}
	if skipList.Del(0) != true {
		t.Fatal("del error")
	}
	if skipList.Len() != 8 {
		t.Fatal("len error")
	}

}
