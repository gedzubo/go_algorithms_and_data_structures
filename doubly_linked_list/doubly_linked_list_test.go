package doubly_linked_list

import "testing"

func TestPush(t *testing.T) {
	dll := setupDoublyLinkedList()

	want := "A -> B -> C -> D"
	got := dll.Print()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestPrint(t *testing.T) {
	dll := setupDoublyLinkedList()

	want := "A -> B -> C -> D"
	got := dll.Print()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestPrintInReverse(t *testing.T) {
	dll := setupDoublyLinkedList()

	want := "D -> C -> B -> A"
	got := dll.PrintInReverse()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestPop(t *testing.T) {
	dll := setupDoublyLinkedList()

	t.Run("Returns correct node value", func(t *testing.T) {
		want := "D"
		got := dll.Pop()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Updates the list correctly", func(t *testing.T) {
		want := "A -> B -> C"
		got := dll.Print()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestUnshift(t *testing.T) {
	dll := NewDoublyLinkedList("B")
	dll.Unshift("A")

	want := "A -> B"
	got := dll.Print()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestShift(t *testing.T) {
	t.Run("Returns correct node value", func(t *testing.T) {
		dll := setupDoublyLinkedList()

		want := "A"
		got := dll.Shift()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Updates the list correctly", func(t *testing.T) {
		dll := setupDoublyLinkedList()

		dll.Shift()

		want := "B -> C -> D"
		got := dll.Print()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestGet(t *testing.T) {
	dll := setupDoublyLinkedList()

	t.Run("Returns correct node value", func(t *testing.T) {
		want := "C"
		got := dll.Get(2)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Returns empty string if index is out of bounds", func(t *testing.T) {
		want := ""
		got := dll.Get(5)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestSet(t *testing.T) {
	dll := setupDoublyLinkedList()

	t.Run("Returns true if node is set", func(t *testing.T) {
		want := true
		got := dll.Set(2, "E")

		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})

	t.Run("Updates the list correctly", func(t *testing.T) {
		dll.Set(2, "E")

		want := "A -> B -> E -> D"
		got := dll.Print()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Returns false if index is out of bounds", func(t *testing.T) {
		want := false
		got := dll.Set(5, "F")

		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})
}

func TestInsert(t *testing.T) {
	t.Run("Returns true if node is inserted", func(t *testing.T) {
		dll := setupDoublyLinkedList()

		want := true
		got := dll.Insert(2, "E")

		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})

	t.Run("Updates the list correctly", func(t *testing.T) {
		dll := setupDoublyLinkedList()

		dll.Insert(2, "E")

		want := "A -> B -> E -> C -> D"
		got := dll.Print()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Returns false if index is out of bounds", func(t *testing.T) {
		dll := setupDoublyLinkedList()

		want := false
		got := dll.Insert(5, "F")

		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})
}

func TestDeleteByIndex(t *testing.T) {
	t.Run("Returns node value if node is deleted", func(t *testing.T) {
		dll := setupDoublyLinkedList()

		want := "C"
		got := dll.DeleteByIndex(2)

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("Updates the list correctly", func(t *testing.T) {
		dll := setupDoublyLinkedList()

		dll.DeleteByIndex(2)

		want := "A -> B -> D"
		got := dll.Print()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Returns empty string if index is out of bounds", func(t *testing.T) {
		dll := setupDoublyLinkedList()

		want := ""
		got := dll.DeleteByIndex(5)

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

func TestReverse(t *testing.T) {
	ll := setupDoublyLinkedList()

	ll.Reverse()

	want := "D -> C -> B -> A"
	got := ll.Print()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func setupDoublyLinkedList() *DoublyLinkedList {
	dll := NewDoublyLinkedList("A")
	dll.Push("B")
	dll.Push("C")
	dll.Push("D")

	return dll
}
