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

func TestFillRow(t *testing.T) {
	g := NewGrid(2, 2)
	g.FillRow(0, "A")
	g.FillRow(1, "B")

	want := "A A \nB B \n"
	got := g.Print()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFillColumn(t *testing.T) {
	g := NewGrid(2, 2)
	g.FillColumn(0, "A")
	g.FillColumn(1, "B")

	want := "A B \nA B \n"
	got := g.Print()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFillAll(t *testing.T) {
	g := NewGrid(2, 2)
	g.FillAll("A")

	want := "A A \nA A \n"
	got := g.Print()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFillDiagonal(t *testing.T) {
	g := NewGrid(2, 2)
	g.FillDiagonal("A")

	want := "A  \n A \n"
	got := g.Print()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFillReverseDiagonal(t *testing.T) {
	g := NewGrid(2, 2)
	g.FillReverseDiagonal("A")

	want := " A \nA  \n"
	got := g.Print()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFillBorder(t *testing.T) {
	g := NewGrid(3, 3)
	g.FillAll("A")
	g.FillBorder("B")

	want := "B B B \nB A B \nB B B \n"
	got := g.Print()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestCheckIfValueIsPresent(t *testing.T) {
	g := NewGrid(3, 3)
	g.FillAll("A")

	t.Run("Returns true if value is present", func(t *testing.T) {
		want := true
		got := g.CheckIfValueIsPresent("A")

		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})

	t.Run("Returns false if value is not present", func(t *testing.T) {
		want := false
		got := g.CheckIfValueIsPresent("B")

		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})
}
