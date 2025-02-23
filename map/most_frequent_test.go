package map_functions

import "testing"

func TestMostFrequent(t *testing.T) {
	t.Run("Returns the most frequent element", func(t *testing.T) {
		s := []string{"A", "B", "D", "A", "C", "B", "A", "D", "D", "D"}
		want := "D"
		got := MostFrequent(s)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Returns the first element if all elements have the same frequency", func(t *testing.T) {
		s := []string{"A", "B", "C", "D"}
		want := "A"
		got := MostFrequent(s)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
