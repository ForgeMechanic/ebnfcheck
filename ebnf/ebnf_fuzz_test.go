package ebnf

import (
	"bytes"
	"testing"
)

// FuzzParseString fuzzes the parser using string inputs.
// The goal is to ensure Parse (and Verify, where applicable) never panic
// and that common comment/string edge cases are exercised.
func FuzzParseString(f *testing.F) {
	seeds := []string{
		`Program = .`,
		`Program = (* comment *) .`,
		`Program = "a" (* inline *) "b" .`,
		`(* file comment *) Program = "a" .`,
		`Program = "(* not comment *)" .`,
		`Program = "contains (*" .`,
		`Program = "contains *)" .`,
		`Program = (* unterminated comment .`,
		`Program = "a" (* missing end .`,
		`Program = "x(*y" "z*)w" .`,
		`Program = ( (*c*) "a" "b" ) .`,
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, src string) {
		g, _ := Parse("", bytes.NewBufferString(src))
		// If we successfully parsed something, exercise Verify as well.
		if g != nil {
			_ = Verify(g, "Program")
		}
	})
}

// FuzzParseBytes fuzzes the parser using arbitrary bytes. This ensures
// the parser's preprocessing (comment stripping) and scanning are robust
// against arbitrary input, including unterminated comments and random data.
func FuzzParseBytes(f *testing.F) {
	seeds := [][]byte{
		[]byte("Program = (* comment *) ."),
		[]byte("Program = \"a\" (* inline comment *) \"b\" ."),
		[]byte("(* file comment *) Program = \"a\" ."),
		[]byte("Program = (* unterminated comment ."),
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, data []byte) {
		_, _ = Parse("", bytes.NewBuffer(data))
	})
}
