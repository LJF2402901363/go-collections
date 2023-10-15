/*
@author: 陌意随影
@since: 2023/9/28 00:29:56
@desc: 使用 go的数组实现一个可变集合List
*/
package collections

import (
	"fmt"
	"reflect"
	"sort"
)

type arrayList[E comparable] struct {
	elements []E
}

func NewArrayList[E comparable]() ArrayList[E] {
	arr := arrayList[E]{}
	arr.elements = make([]E, 0)
	return &arr
}
func (a *arrayList[E]) Add(e E) bool {
	a.elements = append(a.elements, e)
	return true
}

func (a *arrayList[E]) AddAll(c Collection[E]) bool {
	if c == nil {
		return true
	}
	slice := c.ToSlice()
	for _, e := range slice {
		a.Add(e)
	}
	return true
}

func (a *arrayList[E]) AddSlice(elements ...E) bool {
	for _, e := range elements {
		a.Add(e)
	}
	return true
}

func (a *arrayList[E]) Clear() bool {
	a.elements = nil
	return true
}

func (a *arrayList[E]) Contains(e E) bool {
	return a.IndexOf(e) >= 0
}

func (a *arrayList[E]) ContainsAll(c Collection[E]) bool {
	if a == c {
		return true
	}
	if c == nil {
		return false
	}
	slice := c.ToSlice()
	for _, e := range slice {
		if !a.Contains(e) {
			return false
		}
	}
	return true
}

func (a *arrayList[E]) Equals(e Collection[E]) bool {
	if a == e {
		return true
	}

	if a != nil && e == nil {
		return false
	}
	if a == nil && e != nil {
		return false
	}
	if a.Size() != e.Size() {
		return false
	}
	return a.ContainsAll(e) && e.ContainsAll(a)
}

func (a *arrayList[E]) EqualsWithFn(e Collection[E], fn func(o1 E, o2 E) bool) bool {
	if fn == nil {
		return a.Equals(e)
	}
	if a == nil && e != nil {
		return false
	}
	if a != nil && e == nil {
		return false
	}
	if a.Size() != e.Size() {
		return false
	}

	rightSlice := e.ToSlice()
	for i := 0; i < len(a.elements); i++ {
		fla := false
		for _, rightItem := range rightSlice {
			fla = fn(a.elements[i], rightItem)
			if fla {
				break
			}
		}
		if !fla {
			return false
		}
	}

	for _, element := range rightSlice {
		fla := false
		for i := 0; i < len(a.elements); i++ {
			fla = fn(element, a.elements[i])
			if fla {
				break
			}
		}
		if !fla {
			return false
		}
	}
	return true
}

func (a *arrayList[E]) IsEmpty() bool {
	return a.Size() == 0
}

func (a *arrayList[E]) Remove(e E) bool {
	indexOf := a.IndexOf(e)
	if indexOf < 0 {
		return false
	}
	ele := a.RemoveWithIndex(indexOf)
	return e == ele
}

func (a *arrayList[E]) RemoveAll(c Collection[E]) bool {
	if c == nil || c.IsEmpty() {
		return false
	}
	count := 0
	for i := 0; i < a.Size(); i++ {
		if c.Contains(a.elements[i]) {
			count++
		}
	}
	arr := make([]E, a.Size()-count)
	index := 0
	for i := 0; i < a.Size(); i++ {
		e := a.elements[i]
		if c.Contains(e) {
			continue
		}
		arr[index] = e
		index++
	}
	a.elements = arr
	return true
}

func (a *arrayList[E]) RetainAll(c Collection[E]) bool {
	if c == nil || c.Size() == 0 {
		a.elements = make([]E, 0)
		return true
	}
	var newElements []E
	for _, element := range a.elements {
		if c.Contains(element) {
			newElements = append(newElements, element)
		}
	}
	a.elements = newElements
	return true
}

func (a *arrayList[E]) Size() int {
	return len(a.elements)
}

func (a *arrayList[E]) ToSlice() []E {
	arr := make([]E, a.Size())
	copy(arr, a.elements)
	return arr
}

func (a *arrayList[E]) Sort() {
	if a.IsEmpty() {
		return
	}
	sort.SliceStable(a.elements, func(i, j int) bool {
		xType := reflect.TypeOf(a.elements[i])
		if xType.Kind() == reflect.Pointer {
			xType = xType.Elem()
			switch xType.Kind() {
			case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				x := reflect.ValueOf(a.elements[i]).Elem().Uint()
				y := reflect.ValueOf(a.elements[j]).Elem().Uint()
				return x < y
			case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
				x := reflect.ValueOf(a.elements[i]).Elem().Int()
				y := reflect.ValueOf(a.elements[j]).Elem().Int()
				return x < y
			case reflect.String:
				x := reflect.ValueOf(a.elements[i]).Elem().String()
				y := reflect.ValueOf(a.elements[j]).Elem().String()
				return x < y
			case reflect.Float32, reflect.Float64:
				x := reflect.ValueOf(a.elements[i]).Elem().Float()
				y := reflect.ValueOf(a.elements[j]).Elem().Float()
				return x < y
			default:
				panic(fmt.Errorf("%s must implements collections.CompareWith   imterface", xType.String()))
			}
		} else {
			switch xType.Kind() {
			case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				x := reflect.ValueOf(a.elements[i]).Uint()
				y := reflect.ValueOf(a.elements[j]).Uint()
				return x < y
			case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
				x := reflect.ValueOf(a.elements[i]).Int()
				y := reflect.ValueOf(a.elements[j]).Int()
				return x < y
			case reflect.String:
				x := reflect.ValueOf(a.elements[i]).String()
				y := reflect.ValueOf(a.elements[j]).String()
				return x < y
			case reflect.Float32, reflect.Float64:
				x := reflect.ValueOf(a.elements[i]).Float()
				y := reflect.ValueOf(a.elements[j]).Float()
				return x < y
			default:
				panic(fmt.Errorf("just support to sort list contains elements int8,int16,int,int32,int64,uint8,uint16,uint,uint32,uint64,"+
					"float32,float64,string and these Type's pointer,and not support %s", xType.String()))
			}

		}
	})
}
func (a *arrayList[E]) SortWithComparator(comparator Comparator[E]) {
	if comparator == nil {
		panic(fmt.Errorf("comparator can not be nil"))
	}
	if a.Size() <= 1 {
		return
	}
	sort.SliceStable(a.elements, func(i, j int) bool {
		return comparator.Compare(a.elements[i], a.elements[j]) < 0
	})
}
func (a *arrayList[E]) SortWithFn(lessFn func(o1, o2 E) bool) {
	if lessFn == nil {
		panic(fmt.Errorf("lessFn can not be nil"))
	}
	sort.SliceStable(a.elements, func(i, j int) bool {
		return lessFn(a.elements[i], a.elements[j])
	})
}
func (a *arrayList[E]) Get(index int) E {
	if index >= a.Size() {
		panic(fmt.Errorf("toIndex:%d out of index bound：%d", index, a.Size()))
	}
	if index < 0 {
		panic(fmt.Errorf("fromIndex:%d out of index bound：%d", index, 0))
	}
	return a.elements[index]
}

func (a *arrayList[E]) Set(index int, e E) {
	if index >= a.Size() {
		panic(fmt.Errorf("toIndex:%d out of index bound：%d", index, a.Size()))
	}
	if index < 0 {
		panic(fmt.Errorf("fromIndex:%d out of index bound：%d", index, 0))
	}
	a.elements[index] = e
}

func (a *arrayList[E]) AddWithIndex(index int, e E) {
	if index > a.Size() {
		panic(fmt.Errorf("toIndex:%d out of index bound：%d", index, a.Size()))
	}
	if index < 0 {
		panic(fmt.Errorf("fromIndex:%d out of index bound：%d", index, 0))
	}
	if index == a.Size() {
		// 插入尾部
		a.elements = append(a.elements, e)
		return
	}
	a.elements = append(a.elements[:index+1], a.elements[index:]...)
	a.elements[index] = e
}

func (a *arrayList[E]) RemoveWithIndex(index int) E {
	if index >= a.Size() {
		panic(fmt.Errorf("toIndex:%d out of index bound：%d", index, a.Size()))
	}
	if index < 0 {
		panic(fmt.Errorf("fromIndex:%d out of index bound：%d", index, 0))
	}
	e := a.elements[index]
	// 删除最后一个元素
	a.elements = append(a.elements[:index], a.elements[index+1:]...)
	return e
}

func (a *arrayList[E]) IndexOf(e E) int {
	for i := 0; i < a.Size(); i++ {
		if a.elements[i] == e {
			return i
		}
	}
	return -1
}

func (a *arrayList[E]) LastIndexOf(e E) int {
	for i := a.Size() - 1; i >= 0; i-- {
		if a.elements[i] == e {
			return i
		}
	}
	return -1
}

func (a *arrayList[E]) SubList(fromIndex, toIndex int) List[E] {
	if fromIndex > toIndex {
		panic(fmt.Errorf("fromIndex:%d must not greater than toIndex:%d", fromIndex, toIndex))
	}
	if toIndex >= a.Size() {
		panic(fmt.Errorf("toIndex:%d out of index bound：%d", toIndex, a.Size()))
	}
	if fromIndex < 0 {
		panic(fmt.Errorf("fromIndex:%d out of index bound：%d", toIndex, 0))
	}
	eleSize := toIndex - fromIndex
	arr := arrayList[E]{
		elements: make([]E, eleSize),
	}
	copy(arr.elements, a.elements[fromIndex:toIndex])
	return &arr
}

func (a *arrayList[E]) Clone() ArrayList[E] {
	arr := arrayList[E]{}
	arr.elements = a.elements
	arr.elements = make([]E, a.Size())
	copy(arr.elements, a.elements)
	return &arr
}

func (a *arrayList[E]) ToDataSlice() []E {
	return a.elements
}
