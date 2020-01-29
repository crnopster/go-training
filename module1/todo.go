package main

import "fmt"

// DynamicArray .
type DynamicArray struct {
	items []string
	count int
}

func newDynamicArray(capacity int) *DynamicArray {
	array := new(DynamicArray)
	array.items = make([]string, capacity)
	return array
}

func (array *DynamicArray) Add(item string) {
	array.items = append(array.items, item)
	array.count = array.count + 1
}

func (array DynamicArray) Capacity() int {
	i := cap(array.items)

	return i
}

func (array DynamicArray) Size() int {
	i := len(array.items)

	return i
}

func (array DynamicArray) Get(index int) string {
	i := array.items[index]
	return i
}

func (array *DynamicArray) GrowArray() {
	// append doubles the capacity of the slice when it fills up.
	c := cap(array.items)
	if array.count == c {
		arr := make([]string, 0, c*2)
		copy(arr, array.items)
	}
}

func (array *DynamicArray) Remove(index int) {
	array.items = append(array.items[:index], array.items[index+1:]...)
	array.count = array.count - 1
}

func (array *DynamicArray) Insert(index int, item string) {
	if index < 0 {
		fmt.Println("Index can't be negative")
	} else if index <= len(array.items) {
		array.items = append(array.items, "")
		copy(array.items[index+1:], array.items[index:])
		array.items[index] = item
		array.count = array.count + 1
	} else {
		fmt.Printf("Index %v is out of slice range\n", index)
	}

}

func (array DynamicArray) Contains(item string) bool {

	for _, s := range array.items {
		if item == s {
			return true
		}
	}

	return false
}

func (array DynamicArray) IndexOf(item string) int {

	for i, s := range array.items {
		if item == s {
			return i
		}
	}

	return -1
}

func (array *DynamicArray) Clear() {
	array.items = append(array.items[:0])
	array.count = 0
}
