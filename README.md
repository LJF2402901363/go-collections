# 1.简介

>├── README.md
>├── array_list.go							   # ArrayList的实现.go文件
>├── array_list_test.go
>├── collections.go     					# 公共接口，HashSet接口，ArrayList接口所在.go文件
>├── go.mod
>├── hash_set.go                           # HashSet的实现.go文件
>└── hash_set_test.go

​	本项目旨在通过使用go模拟Java的集合框架实现HashSet和ArrayList，从而提供更强大的功能给用户使用。

# 2.支持的方法接口

## 2.1 HashSet的接口

>E 为comparable接口限定的泛型类型



| Add(e E) bool                                              | adds element to collection                                   |
| ---------------------------------------------------------- | ------------------------------------------------------------ |
| AddAll(c Collection[E]) bool                               | Adds all of the elements in the specified collection to this collection (optional operation). |
| AddSlice(elements ...E) bool                               | Adds all of the elements in the specified collection to this collection (optional operation). |
| Clear() bool                                               | Removes all of the elements from this collection (optional operation). |
| Contains(e E) bool                                         | Returns true if this collection contains the specified element. |
| ContainsAll(c Collection[E]) bool                          | Returns true if this collection contains all of the elements in the specified collection. |
| Equals(e Collection[E]) bool                               | Compares the specified object with this collection for equality. When E implements the EqualsWith interface, the comparison is first made using EqualsWith 's IsEquals() |
| EqualsWithFn(e Collection[E], fn func(o1, o2 E) bool) bool | use the custom method fn to compare whether the elements in the collection are equal.if fn is nil, the default Equals() will be used for comparison |
| IsEmpty() bool                                             | Returns true if this collection contains no elements.        |
| Remove(e E) bool                                           | Removes a single instance of the specified element from this collection, if it is present (optional operation). |
| RemoveAll(c Collection[E]) bool                            | Removes all of this collection's elements that are also contained in the specified collection (optional operation). |
| RetainAll(c Collection[E]) bool                            | Retains only the elements in this collection that are contained in the specified collection (optional operation). |
| Size() int                                                 | Returns the number of elements in this collection.           |
| ToSlice() []E                                              | Returns an array containing all of the elements in this collection. |
| Clone() HashSet[E]                                         | Returns a shallow copy of this HashSet instance: the elements themselves are not cloned. |



## 2.2 ArrayList的接口

>E 为comparable接口限定的泛型类型



| Add(e E) bool                                              | adds element to collection                                   |
| ---------------------------------------------------------- | ------------------------------------------------------------ |
| AddAll(c Collection[E]) bool                               | Adds all of the elements in the specified collection to this collection (optional operation). |
| AddSlice(elements ...E) bool                               | Adds all of the elements in the specified collection to this collection (optional operation). |
| Clear() bool                                               | Removes all of the elements from this collection (optional operation). |
| Contains(e E) bool                                         | Returns true if this collection contains the specified element. |
| ContainsAll(c Collection[E]) bool                          | Returns true if this collection contains all of the elements in the specified collection. |
| Equals(e Collection[E]) bool                               | Compares the specified object with this collection for equality. When E implements the EqualsWith interface, the comparison is first made using EqualsWith 's IsEquals() |
| EqualsWithFn(e Collection[E], fn func(o1, o2 E) bool) bool | use the custom method fn to compare whether the elements in the collection are equal.if fn is nil, the default Equals() will be used for comparison |
| IsEmpty() bool                                             | Returns true if this collection contains no elements.        |
| Remove(e E) bool                                           | Removes a single instance of the specified element from this collection, if it is present (optional operation). |
| RemoveAll(c Collection[E]) bool                            | Removes all of this collection's elements that are also contained in the specified collection (optional operation). |
| RetainAll(c Collection[E]) bool                            | Retains only the elements in this collection that are contained in the specified collection (optional operation). |
| Size() int                                                 | Returns the number of elements in this collection.           |
| ToSlice() []E                                              | Returns an array containing all of the elements in this collection. |
| Clone() ArrayList[E]                                       | Returns a shallow copy of this ArrayList instance. (The elements themselves are not copied.) |
| Sort()                                                     | Sorts this list according to the order. The sort is stable: this method must not reorder equal elements. the element of the List must implement the CompareWith interface,otherwise, an exception will be thrown |
| SortWithComparator(c Comparator[E])                        | Sorts this list according to the order induced by the specified Comparator. The sort is stable: this method must not reorder equal elements.  All elements in this list must be mutually comparable using the specified comparator (that is, c.compare(e1, e2) must not throw a ClassCastException for any elements e1 and e2 in the list). |
| SortWithFn(lessFn func(o1, o2 E) bool)                     | Sorts this list according to the order induced by the  lessFn function. lessFn will return true if o1 less than o2.<br/> for example：<br/><br />fun (o1,o2, int) bool {<br/> if o1 < o2 {<br   return true<br/>}<br/> return false<br/>} |
| Get(index int) E                                           | Returns the element at the specified position in this list.<br/> Params:<br/> index – index of the element to return<br/> Returns:<br/> the element at the specified position in this list |
| Set(index int, e E)                                        | Replaces the element at the specified position in this list with the specified element (optional operation).<br/> Params:<br/> index – index of the element to replace element – element to be stored at the specified position<br/> Returns: the element previously at the specified position |
| AddWithIndex(index int, e E)                               | Inserts the specified element at the specified position in this list (optional operation). Shifts the element currently at that position (if any) and any subsequent elements to the right (adds one to their indices).<br/> Params:<br/> index – index at which the specified element is to be inserted element – element to be inserted |
| RemoveWithIndex(index int) E                               | Removes the element at the specified position in this list (optional operation). Shifts any subsequent elements to the left (subtracts one from their indices). Returns the element that was removed from the list.<br/> Params:<br/> index – the index of the element to be removed<br/> Returns: the element previously at the specified position |
| IndexOf(e E) int                                           | Returns the index of the first occurrence of the specified element in this list, or -1 if this list does not contain the element. More formally, returns the lowest index i such that Objects.equals(o, get(i)), or -1 if there is no such index.<br/> Params:<br/> e – element to search for<br/> Returns: the index of the first occurrence of the specified element in this list, or -1 if this list does not contain the element |
| LastIndexOf(e E) int                                       | Returns the index of the last occurrence of the specified element in this list, or -1 if this list does not contain the element. More formally, returns the highest index i such that Objects.equals(o, get(i)), or -1 if there is no such index.<br/> Params:<br/> e  – element to search for<br/> Returns: the index of the last occurrence of the specified element in this list, or -1 if this list does not contain the element |
| SubList(fromIndex, toIndex int) List[E]                    | SubList Returns a view of the portion of this list between the specified fromIndex, inclusive, and toIndex, exclusive. (If fromIndex and toIndex are equal, the returned list is empty.) The returned list is backed by this list, so non-structural changes in the returned list are reflected in this list, and vice-versa. The returned list supports all of the optional list operations supported by this list.<br/> This method eliminates the need for explicit range operations (of the sort that commonly exist for arrays). Any operation that expects a list can be used as a range operation by passing a subList view instead of a whole list. For example, the following idiom removes a range of elements from a list:<br/> list.subList(from, to).clear();<br/> Similar idioms may be constructed for indexOf and lastIndexOf, and all of the algorithms in the Collections class can be applied to a subList.<br/> The semantics of the list returned by this method become undefined if the backing list (i.e., this list) is structurally modified in any way other than via the returned list. (Structural modifications are those that change the size of this list, or otherwise perturb it in such a fashion that iterations in progress may yield incorrect results.)<br/> Params:<br/> fromIndex – low endpoint (inclusive) of the subList toIndex – high endpoint (exclusive) of the subList<br/> Returns: a view of the specified range within this list |

# 3.如何使用

## 3.1使用HashSet

see   [hash_set_test.go](https://github.com/LJF2402901363/go-collections/blob/main/hash_set_test.go)

## 3.2使用ArrayList

see   [array_list_test.go](https://github.com/LJF2402901363/go-collections/blob/main/array_list_test.go)

## 