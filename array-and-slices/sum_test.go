package slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	ErrorMessage := func(t *testing.T, got, want int, numbers []int) {
		t.Helper()
		if got != want {
			t.Errorf("[FAILED]\ngot => %d\nwant => %d\ngiven => %v", got, want, numbers)
		}
	}

	t.Run("Sum a collection of numbers from an array", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		ErrorMessage(t, got, want, numbers)
	})

	t.Run("Sum a collection of numbers from a slice", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		ErrorMessage(t, got, want, numbers)
	})
}

func TestSumAll(t *testing.T) {
	ErrorMessage := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("[FAILED]\ngot => %v\nwant => %v", got, want)
		}
	}

	t.Run("SumAll slices and return a new slice", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		ErrorMessage(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {

	ErrorMessage := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("[FAILED]\ngot => %v\nwant => %v", got, want)
		}
	}

	t.Run("SumAllTails slice last element and returns a new slice", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		ErrorMessage(t, got, want)
	})

	t.Run("SumAllTails with diferent slice lenght", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		ErrorMessage(t, got, want)
	})
}

// func BenchmarkRepeat(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Repeat("a", 5)
// 	}
// }
