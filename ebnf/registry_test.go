package ebnf

import (
	"bytes"
	"testing"
)

func TestGetKnownDialect(t *testing.T) {
	d, ok := GetDialect("go")
	if !ok {
		t.Fatalf("expected 'go' dialect to be registered")
	}
	g, err := d.Parse("", bytes.NewBufferString(`Program = "a" .`))
	if err != nil {
		t.Fatalf("parse failed for go dialect: %v", err)
	}
	if g == nil {
		t.Fatalf("expected grammar, got nil")
	}
}

func TestGetUnknownDialect(t *testing.T) {
	if _, ok := GetDialect("nonexistent"); ok {
		t.Fatalf("unexpectedly found nonexistent dialect")
	}
}

func TestDefaultAlias(t *testing.T) {
	d, ok := GetDialect("default")
	if !ok {
		t.Fatalf("expected 'default' dialect to be registered")
	}
	if d.Name != "w3c" {
		t.Fatalf("default alias expected to refer to 'w3c', got %q", d.Name)
	}
}
