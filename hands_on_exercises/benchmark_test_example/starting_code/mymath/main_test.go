package mymath

import "testing"

func TestCenteredAvg(t *testing.T) {
	type test struct {
		data []int
		answer float64
	}

	tests := []test {
		{[]int{1, 4, 6, 8, 100}, 6},
		{[]int{0, 8, 10, 1000},9},
		{[]int{1, 12, 5, 8, 9},7.333333333333333},
	}
	
	for _, test := range tests {
		result := CenteredAvg(test.data)
		expected := test.answer

		if result != expected {
			t.Error("Expected", expected, "Got", result)
		}
	}

	
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i:=0; i < b.N; i++ {
		CenteredAvg([]int{1,12,5,8,9})
	}
}

