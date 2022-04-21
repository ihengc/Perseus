package Set

/********************************************************
* @author: Ihc
* @date: 2022/4/20 0020 12:00
* @version: 1.0
* @description: 无序不重复集合(使用golang的map实现)
map实现意味着无序集合的元素类型只能是有效的map键类型
可以使用自己实现的哈希表来实现无序不重复集合
*********************************************************/

// ISet 无序集合接口
type ISet interface {
	// Add 添加一个元素到集合中
	Add(elem interface{})
	// Del 删除指定元素到集合中
	Del(elem interface{})
	// In 判断给定的元素是否在集合中
	In(elem interface{}) bool
	// Elements 返回集合中的全部元素
	Elements() []interface{}
	// Len 返回集合中元素个数
	Len() int
	// IsSubsetOf 当前集合是否为给定集合的子集
	IsSubsetOf(superSet ISet) bool
	// IsSuperSetOf 当前集合是否为给定集合的父集
	IsSuperSetOf(subSet ISet) bool
	// Union 并集
	Union(set ISet) ISet
	// Diff 差集
	Diff(set ISet) ISet
	// Intersection 交集
	Intersection(set ISet) ISet
}

type set struct {
	elements map[interface{}]bool
}

// Add 添加一个元素到集合中
func (s *set) Add(elem interface{}) {
	s.elements[elem] = true
}

// Del 删除指定元素到集合中
func (s *set) Del(elem interface{}) {
	delete(s.elements, elem)
}

// In 判断给定的元素是否在集合中
func (s *set) In(elem interface{}) bool {
	if _, ok := s.elements[elem]; ok {
		return true
	}
	return false
}

// Elements 返回集合中的全部元素
func (s *set) Elements() []interface{} {
	elements := make([]interface{}, 0, len(s.elements))
	for elem := range s.elements {
		elements = append(elements, elem)
	}
	return elements
}

// IsSubsetOf 当前集合是否为给定集合的子集
func (s *set) IsSubsetOf(superSet ISet) bool {
	if s.Len() > superSet.Len() {
		return false
	}
	for _, elem := range s.Elements() {
		if !superSet.In(elem) {
			return false
		}
	}
	return true
}

// IsSuperSetOf 当前集合是否为给定集合的父集
func (s *set) IsSuperSetOf(subSet ISet) bool {
	return subSet.IsSubsetOf(s)
}

// Union 并集
func (s *set) Union(set ISet) ISet {
	unionSet := NewSet()
	for _, elem := range s.Elements() {
		unionSet.Add(elem)
	}
	for _, elem := range set.Elements() {
		unionSet.Add(elem)
	}
	return unionSet
}

// Intersection 交集
func (s *set) Intersection(set ISet) ISet {
	intersectionSet := NewSet()
	var minSet, maxSet ISet
	if s.Len() > set.Len() {
		minSet = set
		maxSet = s
	} else {
		minSet = s
		maxSet = set
	}
	for _, elem := range minSet.Elements() {
		if maxSet.In(elem) {
			intersectionSet.Add(elem)
		}
	}
	return intersectionSet
}

// Diff 差集
func (s *set) Diff(set ISet) ISet {
	diffSet := NewSet()
	for _, elem := range s.Elements() {
		if !set.In(elem) {
			diffSet.Add(elem)
		}
	}
	return diffSet
}

// Len 返回集合中元素个数
func (s *set) Len() int {
	return len(s.elements)
}

// NewSet 创建一个集合实例
func NewSet() ISet {
	return &set{elements: make(map[interface{}]bool)}
}
