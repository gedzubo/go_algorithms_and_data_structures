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

func (dll *DoublyLinkedList) PrintInReverse() string {
	var list []string

	currentNode := dll.tail
	for currentNode != nil {
		list = append(list, currentNode.value)
		currentNode = currentNode.previousNode
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

func (dll *DoublyLinkedList) Pop() string {
	if dll.length == 0 {
		return ""
	}

	value := dll.tail.value
	if dll.length == 1 {
		dll.head = nil
		dll.tail = nil
	} else {
		dll.tail = dll.tail.previousNode
		dll.tail.nextNode = nil
	}
	dll.length--

	return value
}

func (dll *DoublyLinkedList) Shift() string {
	if dll.length == 0 {
		return ""
	}

	value := dll.head.value
	if dll.length == 1 {
		dll.head = nil
		dll.tail = nil
	} else {
		dll.head = dll.head.nextNode
		dll.head.previousNode = nil
	}
	dll.length--

	return value
}

func (dll *DoublyLinkedList) Unshift(value string) {
	newNode := &Node{value: value}
	if dll.length == 0 {
		dll.head = newNode
		dll.tail = newNode
	} else {
		dll.head.previousNode = newNode
		newNode.nextNode = dll.head
		dll.head = newNode
	}
	dll.length++
}

func (dll *DoublyLinkedList) Get(index int) string {
	if index < 0 || index >= dll.length {
		return ""
	}

	var currentNode *Node
	if index < dll.length/2 {
		currentNode = dll.head
		for i := 0; i < index; i++ {
			currentNode = currentNode.nextNode
		}
	} else {
		currentNode = dll.tail
		for i := dll.length - 1; i > index; i-- {
			currentNode = currentNode.previousNode
		}
	}

	return currentNode.value
}

func (dll *DoublyLinkedList) Set(index int, value string) bool {
	if index < 0 || index >= dll.length {
		return false
	}

	var currentNode *Node
	if index < dll.length/2 {
		currentNode = dll.head
		for i := 0; i < index; i++ {
			currentNode = currentNode.nextNode
		}
	} else {
		currentNode = dll.tail
		for i := dll.length - 1; i > index; i-- {
			currentNode = currentNode.previousNode
		}
	}

	currentNode.value = value

	return true
}

func (dll *DoublyLinkedList) Insert(index int, value string) bool {
	if index < 0 || index > dll.length {
		return false
	}

	if index == 0 {
		dll.Unshift(value)
		return true
	} else if index == dll.length {
		dll.Push(value)
		return true
	}

	newNode := &Node{value: value}
	var currentNode *Node
	if index < dll.length/2 {
		currentNode = dll.head
		for i := 0; i < index; i++ {
			currentNode = currentNode.nextNode
		}
	} else {
		currentNode = dll.tail
		for i := dll.length - 1; i > index; i-- {
			currentNode = currentNode.previousNode
		}
	}

	newNode.nextNode = currentNode
	newNode.previousNode = currentNode.previousNode
	currentNode.previousNode.nextNode = newNode
	currentNode.previousNode = newNode
	dll.length++

	return true
}

func (dll *DoublyLinkedList) DeleteByIndex(index int) string {
	if index < 0 || index >= dll.length {
		return ""
	}

	var currentNode *Node
	if index < dll.length/2 {
		currentNode = dll.head
		for i := 0; i < index; i++ {
			currentNode = currentNode.nextNode
		}
	} else {
		currentNode = dll.tail
		for i := dll.length - 1; i > index; i-- {
			currentNode = currentNode.previousNode
		}
	}

	value := currentNode.value
	if dll.length == 1 {
		dll.head = nil
		dll.tail = nil
	} else if currentNode == dll.head {
		dll.head = dll.head.nextNode
		dll.head.previousNode = nil
	} else if currentNode == dll.tail {
		dll.tail = dll.tail.previousNode
		dll.tail.nextNode = nil
	} else {
		currentNode.previousNode.nextNode = currentNode.nextNode
		currentNode.nextNode.previousNode = currentNode.previousNode
	}
	dll.length--

	return value
}

func (dll *DoublyLinkedList) Reverse() {
	if dll.length == 0 {
		return
	}

	currentNode := dll.head
	dll.head = dll.tail
	dll.tail = currentNode

	for currentNode != nil {
		prevNode := currentNode.previousNode
		currentNode.previousNode = currentNode.nextNode
		currentNode.nextNode = prevNode
		currentNode = currentNode.previousNode
	}
}
