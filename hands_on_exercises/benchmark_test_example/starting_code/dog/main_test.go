package dog

import "testing"

func TestYears(t *testing.T) {
	result := Years(6)
	expected := 6 * 7

	if result != expected {
		t.Errorf("Expected: %v\tGot:%v\n", expected, result)
	}
}

func TestYearsTwo(t *testing.T) {
	result := YearsTwo(2)
	expected := 14

	if result != expected {
		t.Errorf("Expected: %v\tGot:%v\n", expected, result)
	}
}

func BenchmarkYears(b *testing.B) {
	for i:= 0; i < b.N; i++{
		Years(4)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i:= 0; i < b.N; i++{
		YearsTwo(4)
	}
}