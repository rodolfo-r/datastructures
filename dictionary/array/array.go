package array

import (
	"fmt"

	"github.com/techmexdev/algos/dictionary"
)

// Array represents an array list.
type Array struct {
	list []int
	dictionary.Dictionary
}

const growthFactor = 2
const shrinkFactor = 4

// New creates a dynamic array.
func New() *Array {
	return &Array{}
}

// Insert adds num to the array list.
func (a *Array) Insert(num int) {
	if len(a.list)+1 > cap(a.list) {
		a.growWith(len(a.list)*growthFactor, num)
	} else {
		a.list = append(a.list, num)
	}
}

// Search looks for num in the list, and returns it's index,
// or an error if it doesn't exist.
func (a *Array) Search(num int) (int, error) {
	for i := range a.list {
		if a.list[i] == num {
			return i, nil
		}
	}
	return -1, fmt.Errorf("num %v not in list", num)
}

// Get retrieves the number at provided index,
// or an error if the index is out of bounds.
func (a *Array) Get(index int) (int, error) {
	if index >= len(a.list) {
		return 0, fmt.Errorf("index %v out of bounds", index)
	}
	return a.list[index], nil
}

// Delete removes the element at provided index,
// and returns an error if index is out of bounds.
func (a *Array) Delete(index int) error {
	if index >= len(a.list) || index < 0 {
		return fmt.Errorf("index %v out of bounds", index)
	}
	if shrunkCap := cap(a.list) / shrinkFactor; len(a.list)-1 <= shrunkCap {
		a.shrinkWithout(shrunkCap, index)
	} else {
		a.list = append(a.list[:index], a.list[index+1:]...)
	}
	return nil
}

func (a *Array) growWith(newCap, num int) {
	if len(a.list) == 0 {
		a.list = []int{num}
		return
	}
	newList := make([]int, len(a.list)+1, newCap)
	for i := range a.list {
		newList[i] = a.list[i]
	}
	newList[len(a.list)] = num

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
