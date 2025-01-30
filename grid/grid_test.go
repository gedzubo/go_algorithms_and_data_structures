package grid

import "testing"

func TestSet(t *testing.T) {
	g := NewGrid(2, 2)
	g.Set(0, 0, "A")
	g.Set(0, 1, "B")
	g.Set(1, 0, "C")
	g.Set(1, 1, "D")

	want := "A B \nC D \n"
	got := g.Print()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGet(t *testing.T) {
	g := NewGrid(2, 2)
	g.Set(0, 0, "A")
	g.Set(0, 1, "B")
	g.Set(1, 0, "C")
	g.Set(1, 1, "D")

	t.Run("Returns correct value", func(t *testing.T) {
		want := "B"
		got := g.Get(0, 1)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Returns empty string if row is out of bounds", func(t *testing.T) {
		want := ""
		got := g.Get(2, 1)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Returns empty string if column is out of bounds", func(t *testing.T) {
		want := ""
		got := g.Get(1, 2)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
