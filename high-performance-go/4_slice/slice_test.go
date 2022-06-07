package __slice

import "testing"

func TestLastCharBySlice(t *testing.T) {
	testLastChars(t, lastNumsBySlice)
}

func TestLastCharByCopy(t *testing.T) {
	testLastChars(t, lastNumsByCopy)
}
