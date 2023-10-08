/*
@author: 陌意随影
@since: 2023/9/24 21:29:36
@desc:  使用 map实现一个自定义的hash set
*/

package collections

import (
	"fmt"
)

const (
	// DefaultHashSetCapacity 默认的元素集合长度
	DefaultHashSetCapacity = 10
)

type hashSetImpl[E comparable] map[E]struct{}

// NewHashSet 初始化一个默认的 HashSet
func NewHashSet[E comparable]() HashSet[E] {
	h := make(hashSetImpl[E], DefaultHashSetCapacity)
	return &h
}

// NewHashSetWithCapacity 根据指定的容量capacity初始化一个 HashSet
func NewHashSetWithCapacity[E comparable](capacity int) HashSet[E] {
	h := make(hashSetImpl[E], capacity)
	return &h
}

func (h *hashSetImpl[E]) Add(e E) bool {
	if h.Contains(e) {
		// the element already exists, returns false
		return false
	}
	(*h)[e] = struct{}{}
	return true
}

func (h *hashSetImpl[E]) AddAll(c Collection[E]) bool {
	slice := c.ToSlice()
	fla := true
	for _, item := range slice {
		isAdded := h.Add(item)
		fla = fla && isAdded
	}
	return true
}

func (h *hashSetImpl[E]) AddSlice(list ...E) bool {
	fla := true
	for _, item := range list {
		fla = fla && h.Add(item)
	}
	return true
}

func (h *hashSetImpl[E]) Clear() bool {
	for key := range *h {
		delete(*h, key)
	}
	return true
}

func (h *hashSetImpl[E]) Contains(e E) bool {
	_, ok := (*h)[e]
	return ok
}

func (h *hashSetImpl[E]) ContainsAll(c Collection[E]) bool {
	if h == c {
		return true
	}
	if h == nil && c != nil {
		return false
	}
	if h != nil && c == nil {
		return false
	}
	slice := c.ToSlice()
	for _, item := range slice {
		if !h.Contains(item) {
			return false
		}
	}
	return true
}

func (h *hashSetImpl[E]) EqualsWithFn(e Collection[E], fn func(o1, o2 E) bool) bool {
	if fn == nil {
		panic(fmt.Errorf("fn can not be nil"))
	}
	if h == e {
		return true
	}
	if (h == nil && e != nil) || (h != nil && e == nil) {
		return false
	}
	if h.Size() == 0 && e.Size() == 0 {
		return true
	}
	if h.Size() != e.Size() {
		return false
	}
	srcSlice := h.ToSlice()
	desSlice := e.ToSlice()
	for _, srcItem := range srcSlice {
		fla := false
		for _, desItem := range desSlice {
			if fn(srcItem, desItem) {
				fla = true
				continue
			}
		}
		if !fla {
			return false
		}

	}
	return true
}
func (h *hashSetImpl[E]) Equals(e Collection[E]) bool {
	if h == e {
		return true
	}
	if (h == nil && e != nil) || (h != nil && e == nil) {
		return false
	}
	if h.Size() == 0 && e.Size() == 0 {
		return true
	}
	if h.Size() != e.Size() {
		return false
	}
	return h.ContainsAll(e)
}

func (h *hashSetImpl[E]) IsEmpty() bool {
	return len(*h) == 0
}

func (h *hashSetImpl[E]) Remove(e E) bool {
	if !h.Contains(e) {
		// the element does not exist, the deletion failed
		return false
	}
	delete((*h), e)
	return true
}

func (h *hashSetImpl[E]) RemoveAll(c Collection[E]) bool {
	slice := c.ToSlice()
	fla := true
	for _, e := range slice {
		fla = fla && h.Remove(e)
	}
	return fla
}

func (h *hashSetImpl[E]) RetainAll(c Collection[E]) bool {
	if h == c {
		return true
	}
	var deleteList []E
	for item := range *h {
		if !c.Contains(item) {
			deleteList = append(deleteList, item)
		}
	}
	for _, item := range deleteList {
		delete(*h, item)
	}
	return true
}

func (h *hashSetImpl[E]) Size() int {
	return len(*h)
}

func (h *hashSetImpl[E]) ToSlice() []E {
	var resultList []E
	for item, _ := range *h {
		resultList = append(resultList, item)
	}
	return resultList
}

func (h *hashSetImpl[E]) Clone() HashSet[E] {
	newHashSet := NewHashSetWithCapacity[E](h.Size())
	for item := range *h {
		newHashSet.Add(item)
	}
	return newHashSet
}
