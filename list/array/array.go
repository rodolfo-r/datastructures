package array

// Array represents an unsorted array list.
type Array struct {
	list []interface{}
}

const growthFactor = 2
const shrinkFactor = 4

// New creates a dynamic array.
func New() *Array {
	return &Array{}
}

// Insert adds value to the array list.
func (a *Array) Insert(value interface{}) {
	if len(a.list)+1 > cap(a.list) {
		a.growWith(len(a.list)*growthFactor, value)
	} else {
		a.list = append(a.list, value)
	}
}

// Search looks for value in the list, and returns it's index,
func (a *Array) Search(value interface{}) (index int, ok bool) {
	for i := range a.list {
		if a.list[i] == value {
			return i, true
		}
	}
	return 0, false
}

// Get retrieves the value at provided index,
// or an error if the index is out of bounds.
func (a *Array) Get(index int) (value interface{}, ok bool) {
	if index >= len(a.list) {
		return nil, false
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

func (a *Array) growWith(newCap int, value interface{}) {
	if len(a.list) == 0 {
		a.list = []interface{}{value}
		return
	}
	newList := make([]interface{}, len(a.list)+1, newCap)
	for i := range a.list {
		newList[i] = a.list[i]
	}
	newList[len(a.list)] = value

	a.list = newList
}

// shrinkWithout reassigns the Array's list to a new list
// with cap shrunkCap, and without the element at index.
func (a *Array) shrinkWithout(newCap, index int) {
	newList := make([]interface{}, len(a.list)-1, newCap)
	for i := 0; i < index; i++ {
		newList[i] = a.list[i]
	}

	for i := index + 1; i < len(a.list); i++ {
		newList[i-1] = a.list[i]
	}

	a.list = newList
}
