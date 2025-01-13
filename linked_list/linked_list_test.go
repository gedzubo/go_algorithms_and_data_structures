package linked_list

import "testing"

func TestPush(t *testing.T) {
	ll := setupLinkedList()

	want := "A -> B -> C -> D"
	got := ll.Print()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestPop(t *testing.T) {
	ll := setupLinkedList()

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
		ll := setupLinkedList()

		want := "A"
		got := ll.Shift()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Updates the list correctly", func(t *testing.T) {
		ll := setupLinkedList()

		ll.Shift()

		want := "B -> C -> D"
		got := ll.Print()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestGet(t *testing.T) {
	ll := setupLinkedList()

	t.Run("Returns correct node value", func(t *testing.T) {
		want := "C"
		got := ll.Get(2)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Returns empty string if index is out of bounds", func(t *testing.T) {
		want := ""
		got := ll.Get(5)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestSet(t *testing.T) {
	ll := setupLinkedList()

	t.Run("Returns true if node is set", func(t *testing.T) {
		want := true
		got := ll.Set(2, "E")

		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})

	t.Run("Updates the list correctly", func(t *testing.T) {
		ll.Set(2, "E")

		want := "A -> B -> E -> D"
		got := ll.Print()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Returns false if index is out of bounds", func(t *testing.T) {
		want := false
		got := ll.Set(5, "F")

		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})
}

func TestInsert(t *testing.T) {
	t.Run("Returns true if node is inserted", func(t *testing.T) {
		ll := setupLinkedList()

		want := true
		got := ll.Insert(2, "E")

		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})

	t.Run("Updates the list correctly", func(t *testing.T) {
		ll := setupLinkedList()

		ll.Insert(2, "E")

		want := "A -> B -> E -> C -> D"
		got := ll.Print()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Returns false if index is out of bounds", func(t *testing.T) {
		ll := setupLinkedList()

		want := false
		got := ll.Insert(5, "F")

		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("Returns node value if node is deleted", func(t *testing.T) {
		ll := setupLinkedList()

		want := "C"
		got := ll.Delete(2)

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("Updates the list correctly", func(t *testing.T) {
		ll := setupLinkedList()

		ll.Delete(2)

		want := "A -> B -> D"
		got := ll.Print()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Returns empty string if index is out of bounds", func(t *testing.T) {
		ll := setupLinkedList()

		want := ""
		got := ll.Delete(5)

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

func setupLinkedList() *LinkedList {
	ll := NewLinkedList("A")
	ll.Push("B")
	ll.Push("C")
	ll.Push("D")

	return ll
}
