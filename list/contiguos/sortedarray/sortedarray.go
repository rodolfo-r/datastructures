package sortedarray

// Array represents a sorted array list.
type Array struct {
	list []int
}

const growthFactor = 2
const shrinkFactor = 4

// New creates a dynamic array.
func New() *Array {
	return &Array{}
}

// Insert adds value to the array list.
// O(n)
func (a *Array) Insert(value int) {
	if a.list[0] > value {
		a.list = append([]int{value}, a.list[:]...)
		return
	}
	for i := 1; i < len(a.list); i++ {
		if value > a.list[i] {
			a.list = append(append(a.list[:i-1], value), a.list[i-1:]...)
			return
		}
	}
	a.list = append(a.list, value)
}

// Search applies a binary search.
// O(log n)
func (a *Array) Search(value int) (index int, ok bool) {
	return a.search(value, 0, len(a.list)-1)
}

func (a *Array) search(target, start, end int) (index int, ok bool) {
	if start >= end {
		return 0, false
	}
	if mid := (start + end) / 2; a.list[mid] == target {
		return mid, true
	} else if a.list[mid] < target {
		return a.search(target, start, mid-1)
	} else {
		return a.search(target, mid+1, end)
	}
}

// Get retrieves the value at provided index,
// or an error if the index is out of bounds.
func (a *Array) Get(index int) (value int, ok bool) {
	if index >= len(a.list) {
		return 0, false
	}
	return a.list[index], true
}

// Delete removes the element at provided index,
// and returns an error if index is out of bounds.
func (a *Array) Delete(index int) (ok bool) {
	if index >= len(a.list) || index < 0 {
		return false
	}
	if shrunkCap := cap(a.list) / shrinkFactor; len(a.list)-1 <= shrunkCap {
		a.shrinkWithout(shrunkCap, index)
	} else {
		a.list = append(a.list[:index], a.list[index+1:]...)
	}
	return true
}

func (a *Array) growWith(newCap int, value int) {
	if len(a.list) == 0 {
		a.list = []int{value}
		return
	}
	newList := make([]int, len(a.list)+1, newCap)
	for i := range a.list {
		newList[i] = a.list[i]
	}
	newList[len(a.list)] = value

	a.list = newList
}

// shrinkWithout reassigns the Array's list to a new list
// with cap shrunkCap, and without the element at index.
func (a *Array) shrinkWithout(newCap, index int) {
	newList := make([]int, len(a.list)-1, newCap)
	for i := 0; i < index; i++ {
		newList[i] = a.list[i]
	}

	for i := index + 1; i < len(a.list); i++ {
		newList[i-1] = a.list[i]
	}

	a.list = newList
}
