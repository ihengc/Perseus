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
	t.Log(len(skipList.head.next))
}
