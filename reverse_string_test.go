package reverse_string

import (
	"strings"
	"testing"
)

func TestReverseString(t *testing.T) {
	msgExp := string([]rune{72, 101, 108, 108, 111, 32, 128506, 32, 33})
	msgFunc := ReverseString(ReverseString(msgExp))

	if !strings.EqualFold(msgFunc, msgExp) {
		t.Errorf("Unexpected result:\n\tExpected: %q\n\tGot: %q", msgExp, msgFunc)
	}
}
