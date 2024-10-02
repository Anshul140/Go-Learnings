package arraysandslices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {

	got := SumAll([]int{3, 5}, []int{10, 5, 5})
	want := []int{8, 20}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	} else {
		fmt.Printf("got %v want %v", got, want)
	}
	fmt.Printf("\n")
}

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{2, 4, 6, 8, 1}

		got := Sum(numbers)
		expected := 21

		if got != expected {
			t.Errorf("got %d expected %d given: %v", got, expected, numbers)
		} else {
			fmt.Printf("got %d expected %d given: %v", got, expected, numbers)
		}
	})

	fmt.Printf("\n")

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{2, 4, 6}

		got := Sum(numbers)
		want := 12

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}
