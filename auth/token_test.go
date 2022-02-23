package store

import (
	"testing"
)

func TestGenerateToken(t *testing.T) {
	_, err := GenerateToken("wjl", "admin")
	if err != nil {
		t.Fatal(err)
	}
}

func TestParseToken(t *testing.T) {
	token, err := GenerateToken("wjl", "admin")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := ParseToken(token); err != nil {
		t.Fatal(err)
	}
}
