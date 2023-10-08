/*
@author: 陌意随影
@since: 2023/9/24 20:54:27
@desc: 集合框架的通用接口
*/

package collections

type EqualsWith[E comparable] interface {
	// IsEquals Compares the specified object with this collection for equality.
	IsEquals(e E) bool
}
type CompareWith[E comparable] interface {
	// CompareTo Compares this object with the specified object for order. Returns a negative integer, zero, or a positive integer as this object is less than, equal to, or greater than the specified object.
	//The implementor must ensure signum(x.compareTo(y)) == -signum(y.compareTo(x)) for all x and y. (This implies that x.compareTo(y) must throw an exception if and only if y.compareTo(x) throws an exception.)
	//The implementor must also ensure that the relation is transitive: (x.compareTo(y) > 0 && y.compareTo(z) > 0) implies x.compareTo(z) > 0.
	//Finally, the implementor must ensure that x.compareTo(y)==0 implies that signum(x.compareTo(z)) == signum(y.compareTo(z)), for all z.
	//Params:
	//e – the object to be compared.
	//Returns:
	//a negative integer, zero, or a positive integer as this object is less than, equal to, or greater than the specified object.
	CompareTo(e E) int
}
type Comparator[E comparable] interface {
	// Compare Compares its two arguments for order. Returns a negative integer, zero, or a positive integer as the first argument is less than, equal to, or greater than the second.
	//In the foregoing description, the notation sgn(expression) designates the mathematical signum function, which is defined to return one of -1, 0, or 1 according to whether the value of expression is negative, zero or positive.
	//The implementor must ensure that sgn(compare(x, y)) == -sgn(compare(y, x)) for all x and y. (This implies that compare(x, y) must throw an exception if and only if compare(y, x) throws an exception.)
	//The implementor must also ensure that the relation is transitive: ((compare(x, y)>0) && (compare(y, z)>0)) implies compare(x, z)>0.
	//Finally, the implementor must ensure that compare(x, y)==0 implies that sgn(compare(x, z))==sgn(compare(y, z)) for all z.
	//It is generally the case, but not strictly required that (compare(x, y)==0) == (x.equals(y)). Generally speaking, any comparator that violates this condition should clearly indicate this fact. The recommended language is "Note: this comparator imposes orderings that are inconsistent with equals."
	// Parameters:
	// o1 - the first object to be compared.
	// o2 - the second object to be compared.
	// Returns:
	//a negative integer, zero, or a positive integer as the first argument is less than, equal to, or greater than the second.
	Compare(o1, o2 E) int
}

// Collection 集合通用接口
type Collection[E comparable] interface {
	// Add  adds element to collection
	Add(e E) bool
	// AddAll Adds all of the elements in the specified collection to this collection (optional operation).
	AddAll(c Collection[E]) bool
	// AddSlice Adds all of the elements in the specified collection to this collection (optional operation).
	AddSlice(elements ...E) bool
	// Clear Removes all of the elements from this collection (optional operation).
	Clear() bool
	// Contains Returns true if this collection contains the specified element.
	Contains(e E) bool
	// ContainsAll Returns true if this collection contains all of the elements in the specified collection.
	ContainsAll(c Collection[E]) bool
	// Equals  Compares the specified object with this collection for equality. When E implements the EqualsWith interface, the comparison is first made using EqualsWith 's IsEquals()
	Equals(e Collection[E]) bool
	// EqualsWithFn use the custom method fn to compare whether the elements in the collection are equal.if fn is nil, the default Equals() will be used for comparison
	EqualsWithFn(e Collection[E], fn func(o1, o2 E) bool) bool
	// IsEmpty Returns true if this collection contains no elements.
	IsEmpty() bool
	// Remove Removes a single instance of the specified element from this collection, if it is present (optional operation).
	Remove(e E) bool
	// RemoveAll Removes all of this collection's elements that are also contained in the specified collection (optional operation).
	RemoveAll(c Collection[E]) bool
	// RetainAll Retains only the elements in this collection that are contained in the specified collection (optional operation).
	RetainAll(c Collection[E]) bool
	// Size Returns the number of elements in this collection.
	Size() int
	// ToSlice Returns an array containing all of the elements in this collection.
	ToSlice() []E
}

type Set[E comparable] interface {
	Collection[E]
}
type SortSet[E comparable] interface {
	Set[E]
	// Comparator Returns the comparator used to order the elements in this set, or null if this set uses the natural ordering of its elements.
	Comparator() Comparator[E]
	// First Returns the first (lowest) element currently in this set.
	First() E
	// Last Returns the first (lowest) element currently in this set.
	Last() E
	// HeadSet Returns a view of the portion of this set whose elements are strictly less than toElement.
	HeadSet(toElement E) SortSet[E]
	// SubSet Returns a view of the portion of this set whose elements range from fromElement, inclusive, to toElement, exclusive.
	SubSet(fromElement, toElement E) SortSet[E]
	// TailSet Returns a view of the portion of this set whose elements are greater than or equal to fromElement.
	TailSet(fromElement E) SortSet[E]
}

type HashSet[E comparable] interface {
	Set[E]
	// Clone Returns a shallow copy of this HashSet instance: the elements themselves are not cloned.
	Clone() HashSet[E]
}

type List[E comparable] interface {
	Collection[E]
	// Sort Sorts this list according to the order. The sort is stable: this method must not reorder equal elements. the element of the List must implement the CompareWith interface,otherwise, an exception will be thrown
	Sort()
	// SortWithComparator Sorts this list according to the order induced by the specified Comparator. The sort is stable: this method must not reorder equal elements.
	//	//All elements in this list must be mutually comparable using the specified comparator (that is, c.compare(e1, e2) must not throw a ClassCastException for any elements e1 and e2 in the list).
	SortWithComparator(c Comparator[E])
	// SortWithFn Sorts this list according to the order induced by the  lessFn function. lessFn will return true if o1 less than o2.
	// for example：
	// fun (o1,o2, int) bool {
	//  if o1 < o2 {
	//   return true
	//}
	// return false
	//}
	SortWithFn(lessFn func(o1, o2 E) bool)
	// Get Returns the element at the specified position in this list.
	// Params:
	// index – index of the element to return
	// Returns:
	// the element at the specified position in this list
	Get(index int) E
	// Set Replaces the element at the specified position in this list with the specified element (optional operation).
	// Params:
	// index – index of the element to replace element – element to be stored at the specified position
	// Returns: the element previously at the specified position
	Set(index int, e E)
	// AddWithIndex Inserts the specified element at the specified position in this list (optional operation). Shifts the element currently at that position (if any) and any subsequent elements to the right (adds one to their indices).
	// Params:
	// index – index at which the specified element is to be inserted element – element to be inserted
	AddWithIndex(index int, e E)
	// RemoveWithIndex Removes the element at the specified position in this list (optional operation). Shifts any subsequent elements to the left (subtracts one from their indices). Returns the element that was removed from the list.
	// Params:
	// index – the index of the element to be removed
	// Returns: the element previously at the specified position
	RemoveWithIndex(index int) E
	// IndexOf Returns the index of the first occurrence of the specified element in this list, or -1 if this list does not contain the element. More formally, returns the lowest index i such that Objects.equals(o, get(i)), or -1 if there is no such index.
	// Params:
	// e – element to search for
	// Returns: the index of the first occurrence of the specified element in this list, or -1 if this list does not contain the element
	IndexOf(e E) int
	// LastIndexOf
	// Returns the index of the last occurrence of the specified element in this list, or -1 if this list does not contain the element. More formally, returns the highest index i such that Objects.equals(o, get(i)), or -1 if there is no such index.
	// Params:
	// e  – element to search for
	// Returns: the index of the last occurrence of the specified element in this list, or -1 if this list does not contain the element
	LastIndexOf(e E) int
	// SubList Returns a view of the portion of this list between the specified fromIndex, inclusive, and toIndex, exclusive. (If fromIndex and toIndex are equal, the returned list is empty.) The returned list is backed by this list, so non-structural changes in the returned list are reflected in this list, and vice-versa. The returned list supports all of the optional list operations supported by this list.
	// This method eliminates the need for explicit range operations (of the sort that commonly exist for arrays). Any operation that expects a list can be used as a range operation by passing a subList view instead of a whole list. For example, the following idiom removes a range of elements from a list:
	// list.subList(from, to).clear();
	// Similar idioms may be constructed for indexOf and lastIndexOf, and all of the algorithms in the Collections class can be applied to a subList.
	// The semantics of the list returned by this method become undefined if the backing list (i.e., this list) is structurally modified in any way other than via the returned list. (Structural modifications are those that change the size of this list, or otherwise perturb it in such a fashion that iterations in progress may yield incorrect results.)
	// Params:
	// fromIndex – low endpoint (inclusive) of the subList toIndex – high endpoint (exclusive) of the subList
	// Returns: a view of the specified range within this list
	SubList(fromIndex, toIndex int) List[E]
}

type ArrayList[E comparable] interface {
	List[E]
	// Clone Returns a shallow copy of this ArrayList instance. (The elements themselves are not copied.)
	// Returns: a clone of this ArrayList instance
	Clone() ArrayList[E]
	// ToDataSlice Returns the underlying array for traversing elements
	ToDataSlice() []E
}
