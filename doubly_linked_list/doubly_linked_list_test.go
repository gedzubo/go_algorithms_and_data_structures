package doubly_linked_list

import "testing"

func TestPush(t *testing.T) {
	dll := NewDoublyLinkedList("A")
	dll.Push("B")
	dll.Push("C")

	want := "A -> B -> C"
	got := dll.Print()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestPrintInReverse(t *testing.T) {
	dll := NewDoublyLinkedList("A")
	dll.Push("B")
	dll.Push("C")

	want := "C -> B -> A"
	got := dll.PrintInReverse()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
