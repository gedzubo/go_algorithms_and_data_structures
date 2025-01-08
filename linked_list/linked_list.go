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
		return &LinkedList{
			head:   &Node{value: initialValue},
			tail:   &Node{value: initialValue},
			length: 1,
		}
	}
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

func (ll *LinkedList) Print() string {
	var list []string

	currentNode := ll.head
	for currentNode != nil {
		list = append(list, currentNode.value)
		currentNode = currentNode.nextNode
	}

	return strings.Join(list, " -> ")
}
