package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	got := Sum(nums)
	want := 15

	if got != want {
		t.Errorf("got %q but want %q, given %v", got, want, nums)
	}
}

func TestSumAll(t *testing.T) {
	assertMessage := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but want %v", got, want)
		}
	}

	t.Run("sum all arrays", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{1, 2, 3, 4, 5})
		want := []int{6, 15}
		assertMessage(t, got, want)
	})

	t.Run("zero for empty array", func(t *testing.T) {
		got := SumAll([]int{}, []int{1, 2, 3, 4, 5})
		want := []int{0, 15}
		assertMessage(t, got, want)
	})
}

func TestSumAllTail(t *testing.T) {
	assertMessage := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but want %v", got, want)
		}
	}

	t.Run("sum all tails", func(t *testing.T) {
		got := SumAllTail([]int{1, 2, 3}, []int{1, 2, 3, 4, 5})
		want := []int{3, 5}
		assertMessage(t, got, want)
	})

	t.Run("zero for empty array", func(t *testing.T) {
		got := SumAllTail([]int{}, []int{1, 2, 3, 4, 5})
		want := []int{0, 5}
		assertMessage(t, got, want)
	})
}
