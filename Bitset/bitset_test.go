package Bitset

import "testing"

/********************************************************
* @author: Ihc
* @date: 2022/4/21 0021 11:35
* @version: 1.0
* @description:
*********************************************************/

var bitset *Bitset

func TestNewBitSet(t *testing.T) {
	var err error
	bitset, err = NewBitSetWithNBit(64)
	if err != nil {
		t.Fatal(err)
	}
	if cap(bitset.data) != 2 {
		t.Fatal("NewBitSetWithNBit Err")
	}
}

func TestBitset_Set(t *testing.T) {
	ok, err := bitset.Set(64)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("Set Err")
	}
}

func TestBitset_Get(t *testing.T) {
	ok, err := bitset.Get(64)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("Get Err")
	}
	ok, err = bitset.Get(1213)
	if err != nil {
		t.Fatal("Get Err")
	}
	if ok {
		t.Fatal("Get Err")
	}
}
