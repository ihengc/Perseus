package Bitset

import "testing"

/********************************************************
* @author: Ihc
* @date: 2022/4/21 0021 11:35
* @version: 1.0
* @description:
*********************************************************/

func TestBitset(t *testing.T) {
	bitset, err := NewBitSetWithNBit(64)
	if err != nil {
		t.Fatal(err)
	}
	if bitset.Cap() != 2 {
		t.Fatal("Cap must be 2")
	}
	ret := bitset.Set(128)
	if !ret {
		t.Fatal("Set 128")
	}
	if bitset.Cap() != 4 {
		t.Fatal("Cap must be 4")
	}
	ret = bitset.Set(64)
	if !ret {
		t.Fatal("Set 64")
	}
	ret = bitset.Get(128)
	if !ret {
		t.Fatal("Get 128")
	}
	ret = bitset.Del(64)
	if !ret {
		t.Fatal("Del 64")
	}
	ret = bitset.Get(64)
	if ret {
		t.Fatal("Get 64")
	}
}
