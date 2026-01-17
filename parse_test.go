package main

import (
	"bytes"
	"testing"

	"github.com/forgemechanic/ebnfcheck/ebnf"
)

func TestParseWithDialect_Known(t *testing.T) {
	src := []byte(`Program = "a" .`)
	g, err := parseWithDialect("go", "", bytes.NewBuffer(src))
	if err != nil {
		t.Fatalf("expected parse to succeed: %v", err)
	}
	if g == nil {
		t.Fatalf("expected grammar, got nil")
	}
	if err := ebnf.Verify(g, "Program"); err != nil {
		t.Fatalf("verify failed: %v", err)
	}
}

func TestParseWithDialect_Unknown(t *testing.T) {
	src := []byte(`Program = "a" .`)
	if _, err := parseWithDialect("no-such-dialect", "", bytes.NewBuffer(src)); err == nil {
		t.Fatalf("expected error for unknown dialect")
	}
}
