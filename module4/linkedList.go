package main

type LinkedList struct {
	head  *node
	tail  *node
	count int
}

type node struct {
	value    string
	next     *node
	previous *node
}

type ForwardIterator struct {
	list    *LinkedList
	current *node
}

type BackwardIterator struct {
	list    *LinkedList
	current *node
}

func NewLinkedList() *LinkedList {
	list := new(LinkedList)
	return list
}

func (list *LinkedList) Add(value string) {
	if list.count == 0 {
		n := new(node)
		n.value = value
		list.head = n
		list.tail = n
	} else {
		n := new(node)
		n.value = value

		n.previous = list.tail
		list.tail.next = n

		list.tail = n
	}
	list.count = list.count + 1
}

func (list LinkedList) Size() int {
	return list.count
}

func (list LinkedList) Contains(value string) bool {

	for i := 0; i < list.count; i++ {
		if list.head.value == value {
			return true
		} else {
			list.head = list.head.next
		}

	}
	return false
}

func (list *LinkedList) Remove(value string) bool {

	iterator := list.NewForwardIterator()
	for i := 0; i < list.count; i++ {
		if iterator.current.value == value {

			iterator.current.previous.next = iterator.current.next
			iterator.current.next.previous = iterator.current.previous
			iterator.current = nil
			return true

		} else {
			iterator.Next()
		}
	}
	return false

}

func (list *LinkedList) Clear() {

	list.count = 0
	list.head = nil
	list.tail = nil

}

//**********************************************
//	The Linked List forward Iterator
//**********************************************

func (list *LinkedList) NewForwardIterator() *ForwardIterator {
	return &ForwardIterator{
		current: list.head,
		list:    list,
	}

}

func (iterator *ForwardIterator) HasNext() bool {

	if iterator.current.next != nil {
		return true
	}
	return false

}

func (iterator *ForwardIterator) Next() string {

	if iterator.HasNext() == true {
		iterator.current = iterator.current.next
	} else {
		iterator.current = iterator.list.head
	}
	return iterator.current.value
}

func (iterator *ForwardIterator) Insert(value string) {

	n := new(node)
	n.value = value
	n.previous = iterator.current
	n.next = iterator.current.next
	iterator.current.next = n
	iterator.Next()
	iterator.current.next.previous = iterator.current

}

func (iterator *ForwardIterator) Set(value string) {

	iterator.current.value = value

}

//**********************************************
//	The Linked List backward Iterator
//**********************************************

func (list *LinkedList) NewBackwardIterator() *BackwardIterator {
	return &BackwardIterator{
		current: list.tail,
		list:    list,
	}
}

func (iterator *BackwardIterator) HasPrevious() bool {

	if iterator.current.previous != nil {
		return true
	}

	return false
}

func (iterator *BackwardIterator) Previous() string {

	if iterator.HasPrevious() == true {
		iterator.current = iterator.current.previous
	} else {
		iterator.current = iterator.list.tail
	}
	return iterator.current.value
}

func (iterator *BackwardIterator) Insert(value string) {

	n := new(node)
	n.value = value
	n.next = iterator.current
	n.previous = iterator.current.previous
	iterator.current.previous = n
	iterator.Previous()
	iterator.current.previous.next = iterator.current

}

func (iterator *BackwardIterator) Set(value string) {
	iterator.current.value = value
}
