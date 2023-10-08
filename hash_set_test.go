/*
@author: 陌意随影
@since: 2023/9/24 22:25:25
@desc:
*/
package collections

import (
	"fmt"
	"testing"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func (u *User) CompareTo(e User) int {
	return u.Age - e.Age
}

func (u *User) IsEquals(e User) bool {
	if u.Name == e.Name {
		return true
	}
	return false
}

func Test_hashSetImpl_Add(t *testing.T) {
	intSet := NewHashSet[int]()
	fmt.Println(intSet.Add(1))
	fmt.Println(intSet.Add(2))
	fmt.Println(intSet.Add(3))
	fmt.Println(intSet.Add(1))
}

func Test_hashSetImpl_AddAll(t *testing.T) {
	set1 := NewHashSet[int]()
	set1.Add(1)
	set1.Add(2)
	set2 := NewHashSet[int]()
	set2.Add(1)
	set2.Add(2)
	set2.Add(3)
	set2.Add(4)
	fmt.Println(set1.AddAll(set2))

}

func Test_hashSetImpl_Clear(t *testing.T) {
	set1 := NewHashSet[int]()
	set1.Add(1)
	set1.Add(2)
	fmt.Println(set1.Clear())
}

func Test_hashSetImpl_Clone(t *testing.T) {
	set1 := NewHashSet[int]()
	set1.Add(1)
	set1.Add(2)
	set1.Add(323)
	clone := set1.Clone()
	slice := clone.ToSlice()
	for _, item := range slice {
		fmt.Println(item)
	}
}

func Test_hashSetImpl_Contains(t *testing.T) {
	set1 := NewHashSet[int]()
	set1.Add(1)
	set1.Add(2)
	set1.Add(323)
	fmt.Println(set1.Contains(1))
	fmt.Println(set1.Contains(5))

}

func Test_hashSetImpl_ContainsAll(t *testing.T) {
	set1 := NewHashSet[int]()
	set1.Add(1)
	set1.Add(2)
	set1.Add(323)

	set2 := NewHashSet[int]()
	set2.Add(1)
	set2.Add(2)
	set2.Add(323)

	set3 := NewHashSet[int]()
	set3.Add(1)
	set3.Add(2)
	set3.Add(5)
	fmt.Println(set1.ContainsAll(set1))
	fmt.Println(set1.ContainsAll(set2))
	fmt.Println(set1.ContainsAll(set3))

}

func Test_hashSetImpl_Equals(t *testing.T) {
	set1 := NewHashSet[int]()
	set1.Add(1)
	set1.Add(2)
	set1.Add(323)
	fmt.Println(set1.Equals(set1))
	set2 := NewHashSet[int]()
	set2.Add(1)
	set2.Add(2)
	set2.Add(5)
	fmt.Println(set1.Equals(set2))
	fmt.Println(set2.Equals(set1))
}

func Test_hashSetImpl_Equals1(t *testing.T) {
	set1 := NewHashSet[User]()
	set1.Add(User{
		Name:    "a",
		Age:     1,
		Address: "aadsfad",
	})
	set1.Add(User{
		Name:    "b",
		Age:     1,
		Address: "aadsfad",
	})
	set1.Add(User{
		Name:    "c",
		Age:     1,
		Address: "aadsfad",
	})
	fmt.Println(set1.Equals(set1))
	set2 := NewHashSet[User]()
	set2.Add(User{
		Name:    "a",
		Age:     1,
		Address: "aadsfad",
	})
	set2.Add(User{
		Name:    "b",
		Age:     32,
		Address: "a",
	})
	set2.Add(User{
		Name:    "c",
		Age:     515,
		Address: "aadsfad",
	})
	fmt.Println(set1.Equals(set2))

	set3 := NewHashSet[User]()
	set3.Add(User{
		Name:    "a",
		Age:     1,
		Address: "aadsfad",
	})
	set3.Add(User{
		Name:    "b",
		Age:     1,
		Address: "aadsfad",
	})
	set3.Add(User{
		Name:    "c",
		Age:     1,
		Address: "aadsfad",
	})
	fmt.Println(set2.Equals(set3))
}
func Test_hashSetImpl_IsEmpty(t *testing.T) {
	set1 := NewHashSet[int]()
	set1.Add(1)
	set1.Add(2)
	set1.Add(323)
	fmt.Println(set1.IsEmpty())
}

func Test_hashSetImpl_Remove(t *testing.T) {
	set1 := NewHashSet[int]()
	set1.Add(1)
	set1.Add(2)
	set1.Add(323)
	set1.Remove(1)
	set1.Remove(5)
}

func Test_hashSetImpl_RemoveAll(t *testing.T) {
	set1 := NewHashSet[int]()
	set1.Add(1)
	set1.Add(2)
	set1.Add(323)

	set2 := NewHashSet[int]()
	set2.Add(1)
	set2.Add(2)
	set2.Add(323)
	set2.Add(3)
	fmt.Println(set1.RemoveAll(set2))

	set3 := NewHashSet[int]()
	set3.Add(1)
	set3.Add(2)
	set3.Add(323)
	set3.Add(3)
	set3.Add(23)
	fmt.Println(set2.Equals(set3))
}

func Test_hashSetImpl_RetainAll(t *testing.T) {
	set1 := NewHashSet[int]()
	set1.Add(1)
	set1.Add(2)
	set1.Add(323)
	fmt.Println(set1.ContainsAll(set1))
	set2 := NewHashSet[int]()
	set2.Add(1)
	set2.Add(2)
	set2.Add(323)
	set2.Add(3)
	fmt.Println(set1.ContainsAll(set2))

	set3 := NewHashSet[int]()
	set3.Add(1)
	set3.Add(2)
	fmt.Println(set1.ContainsAll(set3))
}

func Test_hashSetImpl_Size(t *testing.T) {
	set1 := NewHashSet[int]()
	set1.Add(1)
	set1.Add(2)
	set1.Add(323)
	fmt.Println(set1.Size())
}

func Test_hashSetImpl_ToSlice(t *testing.T) {
	set1 := NewHashSet[int]()
	set1.Add(1)
	set1.Add(2)
	set1.Add(323)
	for _, item := range set1.ToSlice() {
		fmt.Println(item)
	}
}

func TestTypes(t *testing.T) {
	//hashSet := NewHashSet[int]()
	//typeOf := reflect.TypeOf(hashSet)
	//var c Collection[int]
	//cType := reflect.TypeOf((*c)(nil)).Elem()
	//fmt.Println(typeOf.Implements((*Collection)nil))
	//fmt.Println(typeOf.Elem())
	//valueOf := reflect.ValueOf(hashSet)
	//fmt.Println(valueOf)
	//fmt.Println(valueOf.Elem())
	//fmt.Println(valueOf)

	set := NewHashSet[any]()
	set.Add(1)
	set.Add("a")
	set.Add(User{
		Name:    "1",
		Age:     109,
		Address: "ada",
	})
	slice := set.ToSlice()
	for i := range slice {
		fmt.Println(slice[i])
	}
}
