package slice_functions

import (
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {
	t.Run("Divides a slice into chunks of a given size", func(t *testing.T) {
		s := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
		chunkSize := 3

		want := [][]string{
			{"A", "B", "C"},
			{"D", "E", "F"},
			{"G", "H"},
		}
		got := Chunk(s, chunkSize)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got[0], want[0])
		}
	})

	t.Run("Handles slices that don't divide evenly", func(t *testing.T) {
		s := []string{"A", "B", "C", "D", "E", "F", "G"}
		chunkSize := 3

		want := [][]string{
			{"A", "B", "C"},
			{"D", "E", "F"},
			{"G"},
		}
		got := Chunk(s, chunkSize)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
