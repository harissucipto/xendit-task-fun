package util

import (
	"strings"
	"testing"
)


func TestRandomInt(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := RandomInt(0, 100)
		if n < 0 || n > 100 {
			t.Errorf("RandomInt(0, 100) = %d, want 0 <= n <= 100", n)
		}
	}
}

func TestRandomEmail(t *testing.T) {
	for i := 0; i < 100; i++ {
		email := RandomEmail()
		if !strings.Contains(email, "@") {
			t.Errorf("RandomEmail() = %s, want @", email)
		}
	}
}


func TestRandomComment(t *testing.T) {
	for i := 0; i < 100; i++ {
		comment := RandomComment()
		if len(comment) < 10 {
			t.Errorf("RandomComment() = %s, want 10 chars", comment)
		}
	}
}