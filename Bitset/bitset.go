package Bitset

import (
	"fmt"
)

/********************************************************
* @author: Ihc
* @date: 2022/4/19 0019 17:24
* @version: 1.0
* @description: 位图

关于位操作
与操作:
	0: 置0的效果
	1: 保持原样
或操作：
	0：保持原样
	1：置1
异或:
	1: 取反
	0: 保持原样
*********************************************************/

// IBitset Bitset接口
type IBitset interface {
	// Set 将指定位的值置为1
	Set(bitIndex int) bool
	// SetTo 将指定位的值置为指定的值(0或1)
	SetTo(bitIndex int, val bool) bool
	// RangeSet 将指定范围内的值置为1
	RangeSet(startIndex int, endIndex int) bool
	// Del 将指定位置的值置为0
	Del(bitIndex int) bool
	// DelAll 将所有位置为9
	DelAll()
	// RangeDel 将指定范围内的值置为0
	RangeDel(startIndex int, endIndex int) bool
	// IsEmpty 表中全部为0
	IsEmpty() bool
	// Count 位的值为1的数量
	Count() int64
}

// Bitset 采用64位数组存放数据
// 8个字节表
type Bitset struct {
	data      []int64 // data 存放状态数据;以64bit为一个单位
	usedCount int64   // usedCount 统计有效位的数量
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
	b.usedCount++
	return true
}

// SetTo 设置指定位的值
func (b *Bitset) SetTo(bitIndex int, val bool) bool {
	if bitIndex < 0 {
		return false
	}
	index := b.index(bitIndex)
	b.expand(index)
	if val {
		b.data[index] |= 1 << (bitIndex % 64)
	} else {
		b.data[index] &= 0 << (bitIndex % 64)
	}
	b.usedCount++
	return true
}

// RangeSet 范围设置为1
// 若startIndex或endIndex为负数,则返回false,表示未设置成功
// 若startIndex >= endIndex,则返回false
// startIndex = endIndex,不应该使用此接口,应该使用 Set 接口
// 若设置成功,则返回true
func (b *Bitset) RangeSet(startIndex int, endIndex int) bool {
	if startIndex < 0 || endIndex < 0 || endIndex <= startIndex {
		return false
	}
	for bitIndex := startIndex; bitIndex <= endIndex; bitIndex++ {
		b.Set(bitIndex)
	}
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
	b.usedCount--
	return true
}

// RangeDel 将指定范围内的值置为0
func (b *Bitset) RangeDel(startIndex, endIndex int) bool {
	if startIndex < 0 || endIndex < 0 || startIndex >= endIndex {
		return false
	}
	// 指定的范围超过表的长度,不执行置0操作,(这里可以实现为全部置0,因为已经有DelAll接口
	// 所以实现全部置0
	// 这里时间复杂度为O(endIndex - startIndex)
	for bitIndex := startIndex; bitIndex <= endIndex; bitIndex++ {
		index := b.index(bitIndex)
		if index >= cap(b.data) {
			return false
		}
	}
	for bitIndex := startIndex; bitIndex <= endIndex; bitIndex++ {
		b.Del(bitIndex)

	}
	return true
}

// DelAll 将所有位置为0
func (b *Bitset) DelAll() {
	for i := 0; i < len(b.data); i++ {
		b.data[i] &= 0
	}
	b.usedCount = 0
}

// Flip 将指定索引处的位的值设置为其补码
func (b *Bitset) Flip(bitIndex int) bool {
	if bitIndex < 0 {
		return false
	}
	index := b.index(bitIndex)
	if index > cap(b.data) {
		return false
	}
	b.data[index] ^= 1 << (bitIndex % 64)
	return true
}

// RangeFlip 将指定范围内的位的值反转
func (b *Bitset) RangeFlip(startIndex, endIndex int) bool {
	if startIndex < 0 || endIndex < 0 || startIndex >= endIndex {
		return false
	}
	for bitIndex := startIndex; bitIndex <= endIndex; bitIndex++ {
		index := b.index(bitIndex)
		if index >= cap(b.data) {
			return false
		}
	}
	for bitIndex := startIndex; bitIndex <= endIndex; bitIndex++ {
		b.Flip(bitIndex)
	}
	return true
}

// Count 统计有效位的数量
func (b *Bitset) Count() int64 {
	return b.usedCount
}

// IsEmpty 是否有1的位
func (b *Bitset) IsEmpty() bool {
	return b.usedCount == 0
}

// Cap 获取容量
func (b *Bitset) Cap() int {
	return cap(b.data)
}
