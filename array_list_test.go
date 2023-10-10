/*
@author: 陌意随影
@since: 2023/9/29 00:01:09
@desc:
*/
package collections

import (
	"fmt"
	"testing"
)

func Test_arrayList_Add(t *testing.T) {
	intList := NewArrayListWithCap[int](2)
	intList.Add(1)
	intList.Add(2)
	intList.Add(1)
	intList1 := NewArrayListWithCap[string](2)
	intList1.Add("1")
	intList1.Add("2")
	intList1.Add("3")
	intList2 := NewArrayListWithCap[User](2)
	intList2.Add(User{
		Name:    "1",
		Age:     1,
		Address: "a",
	})
	intList2.Add(User{
		Name:    "2",
		Age:     2,
		Address: "a",
	})
	intList2.Add(User{
		Name:    "1",
		Age:     1,
		Address: "a",
	})

}

func Test_arrayList_AddAll(t *testing.T) {
	intList := NewArrayListWithCap[int](2)
	intList.Add(1)
	intList.Add(2)
	intList.Add(1)

	intList2 := NewArrayListWithCap[int](2)
	intList2.Add(1)
	intList2.Add(32)
	intList2.Add(323)
	intList.AddAll(intList2)

	intList3 := NewArrayListWithCap[int](2)
	intList.AddAll(intList3)
	intList3.Add(332)
	intList3.AddSlice(1, 2, 3, 4, 5)
	intList.AddSlice(intList3.ToSlice()...)

}

func Test_arrayList_AddWithIndex(t *testing.T) {
	intList := NewArrayListWithCap[int](2)
	intList.Add(1)
	intList.Add(2)
	intList.Add(1)
	intList.AddWithIndex(0, 109)
	intList.AddWithIndex(1, 1000)
	//intList.AddWithIndex(-1, 99)
	intList.AddWithIndex(intList.Size(), 400)
}

func Test_arrayList_Clear(t *testing.T) {
	intList := NewArrayListWithCap[int](2)
	intList.Add(1)
	intList.Add(2)
	intList.Add(1)
	intList.Clear()
	intList.Add(1)
}

func Test_arrayList_Clone(t *testing.T) {
	intList := NewArrayListWithCap[int](2)
	intList.Add(1)
	intList.Add(2)
	intList.Add(1)
	clone := intList.Clone()
	intList.Add(3233)
	clone.Add(11111111111)
}

func Test_arrayList_Contains(t *testing.T) {
	intList := NewArrayListWithCap[int](2)
	intList.Add(1)
	intList.Add(2)
	intList.Add(1)
	fmt.Println(intList.Contains(1))
	fmt.Println(intList.Contains(11111111))
}

func Test_arrayList_ContainsAll(t *testing.T) {
	intList := NewArrayListWithCap[int](2)
	intList.Add(1)
	intList.Add(2)
	intList.Add(1)
	fmt.Println(intList.ContainsAll(intList))
	clone := intList.Clone()
	fmt.Println(intList.ContainsAll(clone))
	fmt.Println(clone.ContainsAll(intList))
	clone.AddSlice(11, 213)
	fmt.Println(intList.ContainsAll(clone))
	fmt.Println(clone.ContainsAll(intList))
}

func Test_arrayList_Equals(t *testing.T) {
	intList := NewArrayListWithCap[int](2)
	intList.Add(1)
	intList.Add(2)
	intList.Add(1)
	fmt.Println(intList.Equals(intList))

	intList1 := NewArrayListWithCap[int](2)
	intList1.Add(1)
	intList1.Add(2)
	fmt.Println(intList.Equals(intList1))
	fmt.Println(intList1.Equals(intList))

	intList1.Add(1)
	fmt.Println(intList.Equals(intList1))
	fmt.Println(intList1.Equals(intList))

}

func Test_arrayList_EqualsWithFn(t *testing.T) {
	fn := func(a, b int) bool {
		return a == b
	}
	intList := NewArrayListWithCap[int](2)
	intList.Add(1)
	intList.Add(2)
	intList.Add(1)
	fmt.Println(intList.EqualsWithFn(intList, fn))

	intList1 := NewArrayListWithCap[int](2)
	intList1.Add(1)
	intList1.Add(2)
	fmt.Println(intList.EqualsWithFn(intList1, fn))
	fmt.Println(intList1.EqualsWithFn(intList, fn))

	intList1.Add(1)
	fmt.Println(intList.EqualsWithFn(intList1, fn))
	fmt.Println(intList1.EqualsWithFn(intList, fn))
}

func Test_arrayList_Get(t *testing.T) {
	intList := NewArrayListWithCap[int](2)
	intList.Add(1)
	intList.Add(2)
	intList.Add(3)
	fmt.Println(intList.Get(0))
	fmt.Println(intList.Get(1))
	fmt.Println(intList.Get(2))
	fmt.Println(intList.Get(3))

}

func Test_arrayList_IndexOf(t *testing.T) {
	intList := NewArrayListWithCap[int](2)
	intList.Add(1)
	intList.Add(2)
	intList.Add(3)
	fmt.Println(intList.IndexOf(0))
	fmt.Println(intList.IndexOf(1))
	fmt.Println(intList.IndexOf(2))
	fmt.Println(intList.IndexOf(3))
}

func Test_arrayList_IsEmpty(t *testing.T) {
	intList := NewArrayListWithCap[int](2)
	fmt.Println(intList.IsEmpty())
	intList.Add(1)
	intList.Add(2)
	intList.Add(3)
	fmt.Println(intList.IsEmpty())

}

func Test_arrayList_LastIndexOf(t *testing.T) {
	intList := NewArrayListWithCap[int](2)
	intList.Add(1)
	intList.Add(2)
	intList.Add(3)
	intList.Add(3)
	fmt.Println(intList.LastIndexOf(0))
	fmt.Println(intList.LastIndexOf(1))
	fmt.Println(intList.LastIndexOf(2))
	fmt.Println(intList.LastIndexOf(3))

	strList := NewArrayListWithCap[string](2)
	strList.Add("a")
	strList.Add("b")
	strList.Add("c")
	strList.Add("c")
	strList.Add("d")
	fmt.Println(strList.LastIndexOf("0"))
	fmt.Println(strList.LastIndexOf("a"))
	fmt.Println(strList.LastIndexOf("c"))
	fmt.Println(strList.LastIndexOf("d"))
}

func Test_arrayList_Remove(t *testing.T) {
	intList := NewArrayListWithCap[int](2)
	intList.Add(1)
	intList.Add(2)
	intList.Add(3)
	intList.Add(3)
	fmt.Println(intList.Remove(3))
	fmt.Println(intList.ToSlice())
}

func Test_arrayList_RemoveAll(t *testing.T) {
	intList := NewArrayListWithCap[int](2)
	intList.Add(1)
	intList.Add(2)
	intList.Add(3)
	intList.Add(3)

	intList1 := NewArrayListWithCap[int](2)
	intList1.Add(2)
	intList1.Add(3)
	intList1.Add(1)
	intList1.Add(6)

	intList.RemoveAll(intList1)
	fmt.Println(intList.ToSlice())
}

func Test_arrayList_RemoveWithIndex(t *testing.T) {
	intList := NewArrayListWithCap[int](2)
	intList.Add(1)
	intList.Add(2)
	intList.Add(3)
	intList.Add(3)
	fmt.Println(intList.RemoveWithIndex(0))
	fmt.Println(intList.RemoveWithIndex(2))
	fmt.Println(intList.ToSlice())
	fmt.Println(intList.ToDataSlice())
}

func Test_arrayList_RetainAll(t *testing.T) {
	intList := NewArrayListWithCap[int](2)
	intList.Add(1)
	intList.Add(2)
	intList.Add(3)
	intList.Add(3)
	fmt.Println(intList.RetainAll(intList))

	intList1 := NewArrayListWithCap[int](2)
	intList1.Add(1)
	//fmt.Println(intList.RetainAll(intList1))
	intList1.Add(2)
	fmt.Println(intList.RetainAll(intList1))
	intList1.Add(3)
	//fmt.Println(intList.RetainAll(intList1))
	intList1.Add(4)
	//fmt.Println(intList.RetainAll(intList1))
}

func Test_arrayList_Set(t *testing.T) {
	intList := NewArrayListWithCap[int](2)
	intList.Add(1)
	intList.Add(2)
	intList.Add(3)
	intList.Add(3)
	intList.Set(0, 1)
	fmt.Println(intList.ToDataSlice())
	intList.Set(0, 100)
	fmt.Println(intList.ToDataSlice())
	intList.Set(5, 10)
	intList.Set(3, 333)

}

func Test_arrayList_Size(t *testing.T) {
	intList := NewArrayListWithCap[*int](2)
	e := 1
	intList.Add(&e)
	b := 2
	intList.Add(&b)
	//intList.Add(2)
	//intList.Add(32332)
	//intList.Add(32)
	//intList.Add(32)
	//fmt.Println(intList.ToDataSlice())
	//strList := NewArrayListWithCap[string](2)
	//strList.Add("a")
	//strList.Add("c")
	//strList.Add("b")
	//strList.Add("e")
	//strList.Add("d")
	//strList.Sort()
	//fmt.Println(strList.ToDataSlice())

}

type DefaultUserComparator struct {
}

func (d *DefaultUserComparator) Compare(o1, o2 User) int {
	return o1.Age - o2.Age
}

func Test_arrayList_Sort(t *testing.T) {
	intList2 := NewArrayListWithCap[User](2)
	intList2.Add(User{
		Name:    "1",
		Age:     1,
		Address: "a",
	})
	intList2.Add(User{
		Name:    "2",
		Age:     2,
		Address: "a",
	})
	intList2.Add(User{
		Name:    "1",
		Age:     1,
		Address: "a",
	})
	intList2.Sort()
	fmt.Println(intList2.ToSlice())
}

func Test_arrayList_SortWithComparator(t *testing.T) {
	intList2 := NewArrayListWithCap[User](2)
	intList2.Add(User{
		Name:    "1",
		Age:     1,
		Address: "a",
	})
	intList2.Add(User{
		Name:    "2",
		Age:     2,
		Address: "a",
	})
	intList2.Add(User{
		Name:    "1",
		Age:     1,
		Address: "a",
	})
	d := new(DefaultUserComparator)
	intList2.SortWithComparator(d)
	fmt.Println(intList2.ToSlice())
}

func Test_arrayList_SortWithFn(t *testing.T) {
	intList2 := NewArrayListWithCap[User](2)
	intList2.Add(User{
		Name:    "1",
		Age:     1,
		Address: "a",
	})
	intList2.Add(User{
		Name:    "2",
		Age:     2,
		Address: "a",
	})
	intList2.Add(User{
		Name:    "1",
		Age:     1,
		Address: "a",
	})
	intList2.SortWithFn(func(o1, o2 User) bool {
		return o1.Age < o2.Age
	})
	fmt.Println(intList2.ToSlice())
}

func Test_arrayList_SubList(t *testing.T) {
	intList := NewArrayList[int]()
	intList.Add(1)
	intList.Add(2)
	intList.Add(3)
	intList.Add(3)
	fmt.Println(intList.SubList(0, 3).ToSlice())
	fmt.Println(intList.SubList(1, 2).ToSlice())
	fmt.Println(intList.SubList(3, 3).ToSlice())
	fmt.Println(intList.SubList(0, 0).ToSlice())

}

func Test_arrayList_ToSlice(t *testing.T) {

}

func Test_arrayList_growth(t *testing.T) {

}
