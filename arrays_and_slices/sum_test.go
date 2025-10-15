package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, numbers)
		}
	})
}

func BenchmarkSum(b *testing.B) {
	numbers := make([]int, 1000)
	for i := 0; i < len(numbers); i++ {
		numbers[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}
