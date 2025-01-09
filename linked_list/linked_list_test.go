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

func TestPop(t *testing.T) {
	ll := NewLinkedList("")
	ll.Push("A")
	ll.Push("B")
	ll.Push("C")
	ll.Push("D")

	t.Run("Returns correct node value", func(t *testing.T) {
		want := "D"
		got := ll.Pop()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Updates the list correctly", func(t *testing.T) {
		want := "A -> B -> C"
		got := ll.Print()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
