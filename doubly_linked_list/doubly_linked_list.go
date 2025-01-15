package doubly_linked_list

import "strings"

type DoublyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewDoublyLinkedList(initialValue string) *DoublyLinkedList {
	if initialValue == "" {
		return &DoublyLinkedList{
			head:   nil,
			tail:   nil,
			length: 0,
		}
	} else {
		newNode := &Node{value: initialValue}
		return &DoublyLinkedList{
			head:   newNode,
			tail:   newNode,
			length: 1,
		}
	}
}

func (dll *DoublyLinkedList) Print() string {
	var list []string

	currentNode := dll.head
	for currentNode != nil {
		list = append(list, currentNode.value)
		currentNode = currentNode.nextNode
	}

	return strings.Join(list, " -> ")
}

func (dll *DoublyLinkedList) Push(value string) {
	newNode := &Node{value: value}
	if dll.length == 0 {
		dll.head = newNode
		dll.tail = newNode
	} else {
		dll.tail.nextNode = newNode
		newNode.previousNode = dll.tail
		dll.tail = newNode
	}
	dll.length++
}
