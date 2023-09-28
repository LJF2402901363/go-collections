/*
@author: 陌意随影
@since: 2023/9/28 00:29:56
@desc: 使用 go的数组实现一个可变集合List
*/
package collections

const (
	// DefaultArrayListCapacity 默认的元素集合长度
	DefaultArrayListCapacity = 10
)

type arrayList[E comparable] struct {
	capacity int
	eleSize  int
	elements []E
}

func NewArrayList[E comparable]() ArrayList[E] {
	arr := arrayList[E]{
		capacity: DefaultArrayListCapacity}
	arr.elements = make([]E, DefaultArrayListCapacity)
	return &arr
}
func (a *arrayList[E]) Add(e E) bool {
	//TODO implement me
	panic("implement me")
}

func (a *arrayList[E]) AddAll(c Collection[E]) bool {
	//TODO implement me
	panic("implement me")
}

func (a *arrayList[E]) AddSlice(elements []E) bool {
	//TODO implement me
	panic("implement me")
}

func (a *arrayList[E]) Clear() bool {
	//TODO implement me
	panic("implement me")
}

func (a *arrayList[E]) Contains(e E) bool {
	//TODO implement me
	panic("implement me")
}

func (a *arrayList[E]) ContainsAll(c Collection[E]) bool {
	//TODO implement me
	panic("implement me")
}

func (a *arrayList[E]) Equals(e Collection[E]) bool {
	//TODO implement me
	panic("implement me")
}

func (a *arrayList[E]) EqualsWithFn(e Collection[E], fn func(o1 E, o2 E) bool) bool {
	//TODO implement me
	panic("implement me")
}

func (a *arrayList[E]) IsEmpty() bool {
	//TODO implement me
	panic("implement me")
}

func (a *arrayList[E]) Remove(e E) bool {
	//TODO implement me
	panic("implement me")
}

func (a *arrayList[E]) RemoveAll(c Collection[E]) bool {
	//TODO implement me
	panic("implement me")
}

func (a *arrayList[E]) RetainAll(c Collection[E]) bool {
	//TODO implement me
	panic("implement me")
}

func (a *arrayList[E]) Size() int {
	//TODO implement me
	panic("implement me")
}

func (a *arrayList[E]) ToSlice() []E {
	//TODO implement me
	panic("implement me")
}

func (a *arrayList[E]) Sort(c Comparator[E]) {
	//TODO implement me
	panic("implement me")
}

func (a *arrayList[E]) Get(index int) E {
	//TODO implement me
	panic("implement me")
}

func (a *arrayList[E]) Set(index int, e E) {
	//TODO implement me
	panic("implement me")
}

func (a *arrayList[E]) AddWithIndex(index int, e E) {
	//TODO implement me
	panic("implement me")
}

func (a *arrayList[E]) RemoveWithIndex(index int) E {
	//TODO implement me
	panic("implement me")
}

func (a *arrayList[E]) IndexOf(e E) int {
	//TODO implement me
	panic("implement me")
}

func (a *arrayList[E]) LastIndexOf(e E) int {
	//TODO implement me
	panic("implement me")
}

func (a *arrayList[E]) SubList(fromIndex, toIndex int) List[E] {
	//TODO implement me
	panic("implement me")
}

func (a *arrayList[E]) Clone() ArrayList[E] {
	//TODO implement me
	panic("implement me")
}
