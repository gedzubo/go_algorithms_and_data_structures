package linked_list

import (
	"strings"
)

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewLinkedList(initialValue string) *LinkedList {
	if initialValue == "" {
		return &LinkedList{
			head:   nil,
			tail:   nil,
			length: 0,
		}
	} else {
		newNode := &Node{value: initialValue}
		return &LinkedList{
			head:   newNode,
			tail:   newNode,
			length: 1,
		}
	}
}

func (ll *LinkedList) Print() string {
	var list []string

	currentNode := ll.head
	for currentNode != nil {
		list = append(list, currentNode.value)
		currentNode = currentNode.nextNode
	}

	return strings.Join(list, " -> ")
}

func (ll *LinkedList) Push(value string) {
	newNode := &Node{value: value}
	if ll.length == 0 {
		ll.head = newNode
		ll.tail = newNode
	} else {
		ll.tail.nextNode = newNode
		ll.tail = newNode
	}
	ll.length++
}

func (ll *LinkedList) Pop() string {
	if ll.length == 0 {
		return ""
	}

	currentNode := ll.head
	if ll.length == 1 {
		ll.head = nil
		ll.tail = nil
		ll.length--
		return currentNode.value
	}

	for currentNode.nextNode.nextNode != nil {
		currentNode = currentNode.nextNode
	}

	value := currentNode.nextNode.value
	currentNode.nextNode = nil
	ll.tail = currentNode
	ll.length--

	return value
}

func (ll *LinkedList) Unshift(value string) {
	newNode := &Node{value: value}
	if ll.length == 0 {
		ll.head = newNode
		ll.tail = newNode
	} else {
		newNode.nextNode = ll.head
		ll.head = newNode
	}
	ll.length++
}

func (ll *LinkedList) Shift() string {
	if ll.length == 0 {
		return ""
	}

	tmp := ll.head
	ll.head = ll.head.nextNode
	tmp.nextNode = nil
	ll.length--

	return tmp.value
}

func (ll *LinkedList) Get(index int) string {
	if index < 0 || index >= ll.length {
		return ""
	}

	currentNode := ll.head
	for i := 0; i < index; i++ {
		currentNode = currentNode.nextNode
	}

	return currentNode.value
}

func (ll *LinkedList) Set(index int, value string) bool {
	if index < 0 || index >= ll.length {
		return false
	}

	currentNode := ll.head
	for i := 0; i < index; i++ {
		currentNode = currentNode.nextNode
	}

	currentNode.value = value

	return true
}

func (ll *LinkedList) Insert(index int, value string) bool {
	if index < 0 || index > ll.length {
		return false
	}

	if index == 0 {
		ll.Unshift(value)
		return true
	}

	if index == ll.length {
		ll.Push(value)
		return true
	}

	newNode := &Node{value: value}
	currentNode := ll.head
	for i := 0; i < index-1; i++ {
		currentNode = currentNode.nextNode
	}

	newNode.nextNode = currentNode.nextNode
	currentNode.nextNode = newNode
	ll.length++

	return true
}

func (ll *LinkedList) Delete(index int) string {
	if index < 0 || index >= ll.length {
		return ""
	}

	if index == 0 {
		return ll.Shift()
	}

	if index == ll.length-1 {
		return ll.Pop()
	}

	currentNode := ll.head
	for i := 0; i < index-1; i++ {
		currentNode = currentNode.nextNode
	}

	value := currentNode.nextNode.value
	currentNode.nextNode = currentNode.nextNode.nextNode
	ll.length--

	return value
}

func (ll *LinkedList) Reverse() {
	if ll.length == 0 {
		return
	}

	currentNode := ll.head
	ll.head = ll.tail
	ll.tail = currentNode

	var prevNode *Node
	var nextNode *Node
	for i := 0; i < ll.length; i++ {
		nextNode = currentNode.nextNode
		currentNode.nextNode = prevNode
		prevNode = currentNode
		currentNode = nextNode
	}
}

func (ll *LinkedList) FindMiddleNodeValue() string {
	if ll.length == 0 {
		return ""
	}

	slowPointer := ll.head
	fastPointer := ll.head

	for fastPointer != nil && fastPointer.nextNode != nil {
		slowPointer = slowPointer.nextNode
		fastPointer = fastPointer.nextNode.nextNode
	}

	return slowPointer.value
}
