package SkipList

import (
	"math/rand"
	"sync"
	"time"
)

/****************************************************************
 * @author: Ihc
 * @date: 2022/4/19 22:47
 * @description: 跳跃表
 ***************************************************************/

var (
	defaultComparator  = BuiltinTypeComparator
	defaultMaxLevel    = 16
	defaultLock        sync.RWMutex
	defaultProbability = 0.5
)

// Comparator 比较函数
type Comparator func(x, y interface{}) int

// Option 用于设置跳跃表的初始化选项
type Option func(option *Options)

// Options 跳跃表初始化选项
type Options struct {
	comparator  Comparator  // scoreCmp 分数比较大小规则
	maxLevel    int         // maxLevel 最大层高
	locker      sync.Locker // locker 线程安全
	probability float64     // probability 上升概率
}

// node 跳跃表节点
type node struct {
	next  []*node     // next 当前节点的后序节点
	score interface{} // score 分数
	data  interface{} // data 当前节点存放的数据
}

// SkipList 跳跃表
// 第一层起始位置为0,最高层位置为maxLevel-1
// 表中放置哨兵节点head
type SkipList struct {
	locker      sync.Locker // locker 并发控制
	head        *node       // head 头节点
	maxLevel    int         // maxLevel 最高层级
	len         int         // len 表中存放数据的个数
	comparator  Comparator  // comparator 比较函数
	probability float64     // probability 上升概率
	rand        *rand.Rand  // rand 随机参数
}

// NewSkipListWithOption 指定选项创建跳跃表
func NewSkipListWithOption(opts ...Option) *SkipList {
	defaultOptions := Options{
		comparator:  defaultComparator,
		maxLevel:    defaultMaxLevel,
		locker:      &defaultLock,
		probability: defaultProbability,
	}
	for _, opt := range opts {
		opt(&defaultOptions)
	}
	skl := &SkipList{
		locker:      defaultOptions.locker,
		maxLevel:    defaultOptions.maxLevel,
		comparator:  defaultOptions.comparator,
		probability: defaultOptions.probability,
		rand:        rand.New(rand.NewSource(time.Now().Unix())),
	}
	skl.head = &node{}
	skl.head.next = make([]*node, skl.maxLevel)
	return skl
}

// randomLevel 随机插入节点的层高
func (skl *SkipList) randomLevel() int {
	level := 1
	for skl.rand.Float64() < skl.probability && level < skl.maxLevel {
		level++
	}
	return level
}

// findPrevNodes 查询目的分数节点位置
func (skl *SkipList) findPrevNodes(score interface{}) []*node {
	prevNodes := make([]*node, skl.maxLevel)
	current := skl.head
	// 从顶层开始,逐层查找,查找每层的插入位置
	for i := skl.maxLevel - 1; i >= 0; i-- {
		if skl.head.next[i] != nil {
			// 在当前层中查询目标分数
			for next := current.next[i]; next != nil; next = next.next[i] {
				// 遇到比当前分数大的节点就停止
				if skl.comparator(next.score, score) >= 0 {
					break
				}
				current = current.next[i]
			}
		}
		prevNodes[i] = current
	}
	return prevNodes
}

// Put 放入数据到跳跃表
// score 节点分数; data 节点数据
func (skl *SkipList) Put(score interface{}, data interface{}) {
	skl.locker.Lock()
	defer skl.locker.Unlock()
	prevNodes := skl.findPrevNodes(score)
	if prevNodes[0].next[0] != nil && skl.comparator(prevNodes[0].next[0].score, score) == 0 {
		prevNodes[0].next[0].data = data
		return
	}
	level := skl.randomLevel()
	newNode := &node{
		score: score,
		data:  data,
		next:  make([]*node, level),
	}
	for i := range newNode.next {
		newNode.next[i] = prevNodes[i].next[i]
		prevNodes[i].next[i] = newNode
	}
	skl.len++
}

// Get 获取数据
// 若分数在表中不存在,则返回nil
// score 分数
func (skl *SkipList) Get(score interface{}) interface{} {
	prev := skl.head
	for i := skl.maxLevel - 1; i >= 0; i-- {
		for current := prev.next[i]; current != nil; current = current.next[i] {
			ret := skl.comparator(current.score, score)
			if ret == 0 {
				return current.data
			}
			if ret > 0 {
				break
			}
			prev = current.next[i]
		}
	}
	return nil
}

// Del 删除
// 成功删除返回true,否则返回false
func (skl *SkipList) Del(score interface{}) bool {
	skl.locker.Lock()
	defer skl.locker.Unlock()

	prevNodes := skl.findPrevNodes(score)
	node := prevNodes[0].next[0]
	if node == nil {
		return false
	}
	if node != nil && skl.comparator(node.score, score) != 0 {
		return false
	}
	for i, n := range node.next {
		prevNodes[i].next[i] = n
	}
	skl.len--
	return true
}

// Len 元素个数
func (skl *SkipList) Len() int {
	skl.locker.Lock()
	defer skl.locker.Unlock()
	return skl.len
}
