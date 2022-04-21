package Bitset

import "testing"

/********************************************************
* @author: Ihc
* @date: 2022/4/21 0021 11:35
* @version: 1.0
* @description:
*********************************************************/

var bitset = NewBitset()

func TestNewBitSet(t *testing.T) {
}

func TestBitset_Set(t *testing.T) {
	ok, err := bitset.Set(1)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("set error")
	}

	v, err := bitset.Get(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(v)
}
