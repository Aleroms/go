package word

import (
	"testing"
)

func TestUseCount(t *testing.T) {
	s := "A string of words"
	result := UseCount(s)
	expect := map[string]int {
		"A": 1,
		"string": 1,
		"of": 1,
		"words": 1,
	}
	for k, v := range result{
		if expect[k] != v {
			t.Error("result:",result,"expected:", expect)
		}
	}
	
}

func TestCount(t *testing.T) {
	s := "there are five words here"
	expected := 5
	result := Count(s)

	if result != expected {
		t.Errorf("Count failed. Expected: %v but got: %v",expected, result)
	}
}