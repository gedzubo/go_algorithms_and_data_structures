package linked_list

import "testing"

func TestPush(t *testing.T) {
	ll := NewLinkedList("")
	ll.Push("A")
	ll.Push("B")
	ll.Push("C")
	ll.Push("D")

	want := "A -> B -> C -> D"
	got := ll.Print()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
