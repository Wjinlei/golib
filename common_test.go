package golib

import (
	"testing"
)

func TestContains(t *testing.T) {
	t.Log(Contains([]int{1, 2, 3}, 3))

	t.Log(Contains([]int{1, 2, 3}, 4))

	t.Log(Contains([]string{"aa", "hello", "bb"}, "bb"))

	t.Log(Contains([]string{"aa", "hello", "bb"}, "ss"))

	t.Log(Contains([]float64{33.22, 23.54, 7.8}, 7.8))

	t.Log(Contains([]float64{33.22, 23.54, 7.8}, 7.9))
}
