package ebnf

import (
	"errors"
	"strings"
	"testing"
	"text/scanner"
)

func TestPosMethods(t *testing.T) {
	p1 := scanner.Position{Filename: "file.ebnf", Offset: 10}
	p2 := scanner.Position{Filename: "file.ebnf", Offset: 20}

	n := &Name{StringPos: p1, String: "N"}
	if got := n.Pos(); got != p1 {
		t.Fatalf("Name.Pos() = %v, want %v", got, p1)
	}

	tok := &Token{StringPos: p2, String: "t"}
	if got := tok.Pos(); got != p2 {
		t.Fatalf("Token.Pos() = %v, want %v", got, p2)
	}

	r := &Range{Begin: tok, End: &Token{StringPos: p2}}
	if got := r.Pos(); got != tok.Pos() {
		t.Fatalf("Range.Pos() = %v, want %v", got, tok.Pos())
	}

	g := &Group{Lparen: p1, Body: nil}
	if got := g.Pos(); got != p1 {
		t.Fatalf("Group.Pos() = %v, want %v", got, p1)
	}

	o := &Option{Lbrack: p1, Body: nil}
	if got := o.Pos(); got != p1 {
		t.Fatalf("Option.Pos() = %v, want %v", got, p1)
	}

	rp := &Repetition{Lbrace: p1, Body: nil}
	if got := rp.Pos(); got != p1 {
		t.Fatalf("Repetition.Pos() = %v, want %v", got, p1)
	}

	prod := &Production{Name: n, Expr: nil}
	if got := prod.Pos(); got != n.Pos() {
		t.Fatalf("Production.Pos() = %v, want %v", got, n.Pos())
	}

	b := &Bad{TokPos: p2, Error: "err"}
	if got := b.Pos(); got != p2 {
		t.Fatalf("Bad.Pos() = %v, want %v", got, p2)
	}
}

func TestErrorListError(t *testing.T) {
	var list errorList
	if list.Err() != nil {
		t.Fatal("expected Err() to return nil for empty list")
	}

	list = append(list, errors.New("first"))
	if list.Err() == nil {
		t.Fatal("expected non-nil Err() for non-empty list")
	}
	if s := list.Error(); s != "first" {
		t.Fatalf("Error() = %q, want %q", s, "first")
	}

	list = append(list, errors.New("second"))
	s := list.Error()
	if !strings.Contains(s, "and 1 more errors") {
		t.Fatalf("Error() = %q, want suffix containing %q", s, "and 1 more errors")
	}
}
