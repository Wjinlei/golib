package system

import (
	"fmt"
	"testing"
)

func TestGetEnv(t *testing.T) {
	fmt.Println(GetEnv("tmp", "/tmp"))
}
