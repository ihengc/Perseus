package Bitset

import (
	"fmt"
)

/********************************************************
* @author: Ihc
* @date: 2022/4/19 0019 17:24
* @version: 1.0
* @description: 位图
*********************************************************/

// Bitset 采用64位数组存放数据
// 存放0-63位数据
type Bitset struct {
	data []int64 // data 存放状态数据;以64bit为一个单位
}

// NewBitSetWithNBit 指定位数初始化 Bitset
func NewBitSetWithNBit(nbits int) (*Bitset, error) {
	// 处理nbits不合法的情况
	if nbits < 0 {
		return nil, fmt.Errorf("nbits < 0:%d", nbits)
	}
	// 64位个数 = 初始容量右移6位 + 1
	return &Bitset{data: make([]int64, nbits>>6+1, nbits>>6+1)}, nil
}

// index 获取带设置的位在数组中的索引位置
func (b *Bitset) index(bitIndex int) int {
	return bitIndex >> 6
}

// expand 扩容
func (b *Bitset) expand(index int) {
	oldCap := cap(b.data)
	if oldCap < index+1 {
		// 两倍于原容量进行扩容
		newCap := oldCap * 2
		newData := make([]int64, newCap, newCap)
		copy(newData[:oldCap], b.data)
		b.data = newData
	}
}

// Set 设置指定索引处的状态为1
func (b *Bitset) Set(bitIndex int) bool {
	if bitIndex < 0 {
		return false
	}
	index := b.index(bitIndex)
	b.expand(index)
	b.data[index] |= 1 << (bitIndex % 64)
	return true
}

// Get 获取指定索引处的状态
// 若获取成功,则返回true,反之返回false
func (b *Bitset) Get(bitIndex int) bool {
	if bitIndex < 0 {
		return false
	}
	index := b.index(bitIndex)
	return index < cap(b.data) && b.data[index]&(1<<(bitIndex%64)) == 1
}

// Del 將指定索引处的位设置为0
// 返回是否删除成功
func (b *Bitset) Del(bitIndex int) bool {
	if bitIndex < 0 {
		return false
	}
	index := b.index(bitIndex)
	if index > cap(b.data) {
		return false
	}
	if !b.Get(bitIndex) {
		return false
	}
	b.data[index] ^= 1 << (bitIndex % 64)
	return true
}

// Cap 获取容量
func (b *Bitset) Cap() int {
	return cap(b.data)
}
