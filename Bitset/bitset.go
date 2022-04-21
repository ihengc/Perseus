package Bitset

import (
	"errors"
	"sync"
)

/********************************************************
* @author: Ihc
* @date: 2022/4/19 0019 17:24
* @version: 1.0
* @description: 位图
*********************************************************/

// IBitset 接口
type IBitset interface {
	And(bitset IBitset)
	AndNot(bitset Bitset)
	Clear()
	ClearRange(start uint64, end uint64)
	Get(index uint64) (bool, error)
	Set(index uint64) (bool, error)
	SetStatus(index uint64, status bool) (bool, error)
	SetRangeStatus(start uint64, end uint64, status bool) (bool, error)
	IsEmpty() bool
	Len() uint64
	Size() uint64
	Or(bitset Bitset)
	Xor(bitset Bitset)
}

// Options 初始化选项
type Options struct {
	Locker          sync.Locker // Locker 设置Bitset是否为线程安全
	MaxCapacity     int         // BitSetMaxCapacity 指定Bitset的最大容量;单位为字节
	InitialCapacity int         // InitialCapacity 初始化容量;单位为字节
}

// Option 配置初始化选项
type Option func(options *Options)

var (
	defaultMaxCapacity     = 1024 * 1024 * 512
	defaultLocker          = &sync.Mutex{}
	defaultInitialCapacity = 8
	IndexOutOfRangeErr     = errors.New("index out of range")
)

// Bitset 位图
type Bitset struct {
	locker      sync.Locker
	data        []byte
	maxCapacity uint64
}

// NewBitset 创建位图实例
func NewBitset(opts ...Option) *Bitset {
	options := Options{
		InitialCapacity: defaultInitialCapacity,
		Locker:          defaultLocker,
		MaxCapacity:     defaultMaxCapacity,
	}
	for _, opt := range opts {
		opt(&options)
	}
	if options.InitialCapacity <= 0 {
		options.InitialCapacity = defaultInitialCapacity
	}
	if options.MaxCapacity <= 0 {
		options.MaxCapacity = defaultMaxCapacity
	}
	if options.InitialCapacity > options.MaxCapacity {
		options.InitialCapacity = options.MaxCapacity
	}
	return &Bitset{
		data:        make([]byte, options.InitialCapacity),
		maxCapacity: uint64(options.MaxCapacity),
		locker:      options.Locker,
	}
}

func (b *Bitset) Set(index uint64) (bool, error) {
	b.locker.Lock()
	defer b.locker.Unlock()
	if index < 0 || index > b.maxCapacity*8 {
		return false, IndexOutOfRangeErr
	}
	realIndex := index % 8
	b.data[realIndex] = 1
	return true, nil
}

func (b *Bitset) Get(index uint64) (bool, error) {
	b.locker.Lock()
	defer b.locker.Unlock()
	if index < 0 || index > b.maxCapacity*8 {
		return false, IndexOutOfRangeErr
	}
	realIndex := index % 8
	if b.data[realIndex] == 0 {
		return false, nil
	}
	if b.data[realIndex] == 1 {
		return true, nil
	}
	return false, nil
}
