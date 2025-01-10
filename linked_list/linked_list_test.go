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

func TestUnshift(t *testing.T) {
	ll := NewLinkedList("B")
	ll.Unshift("A")

	want := "A -> B"
	got := ll.Print()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestShift(t *testing.T) {
	t.Run("Returns correct node value", func(t *testing.T) {
		ll := NewLinkedList("A")
		ll.Push("B")
		ll.Push("C")
		ll.Push("D")

		want := "A"
		got := ll.Shift()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Updates the list correctly", func(t *testing.T) {
		ll := NewLinkedList("A")
		ll.Push("B")
		ll.Push("C")
		ll.Push("D")

		ll.Shift()

		want := "B -> C -> D"
		got := ll.Print()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
