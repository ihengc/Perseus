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

func TestNewSkipListWithOption(t *testing.T) {
	skipList := NewSkipListWithOption(OptionFunc)
	t.Log(skipList.randomLevel())
	t.Log(skipList.Len())
	skipList.Put(1, 2)
	t.Log(skipList.Len())
	t.Log(skipList.Get(1))
	t.Log(skipList.Del(1))
	t.Log(skipList.Get(1))
}
