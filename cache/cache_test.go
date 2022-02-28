package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestSet(t *testing.T) {
	Set("name", "wjl", 30)
	for {
		time.Sleep(time.Duration(1) * time.Second)
		fmt.Println(Get("name"))
	}
}

func TestGet(t *testing.T) {
	if Get("name") == nil {
		t.Error(fmt.Errorf("Not found cache"))
	}
}
