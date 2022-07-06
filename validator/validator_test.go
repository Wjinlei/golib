package validator

import "testing"

func TestExists(t *testing.T) {
	exists := Exists("validator.go")
	if exists != true {
		t.Fatal()
	}
}

func TestHas(t *testing.T) {
	t.Log(Has([]int{1, 2, 3}, 3))

	t.Log(Has([]int{1, 2, 3}, 4))

	t.Log(Has([]string{"aa", "hello", "bb"}, "bb"))

	t.Log(Has([]string{"aa", "hello", "bb"}, "ss"))

	t.Log(Has([]float64{33.22, 23.54, 7.8}, 7.8))

	t.Log(Has([]float64{33.22, 23.54, 7.8}, 7.9))
}
