// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ebnf

import (
	"bytes"
	"testing"
)

var goodGrammars = []string{
	`Program = .`,

	`Program = foo .
	 foo = "foo" .`,

	`Program = "a" | "b" "c" .`,

	`Program = "a" … "z" .`,

	`Program = Song .
	 Song = { Note } .
	 Note = Do | (Re | Mi | Fa | So | La) | Ti .
	 Do = "c" .
	 Re = "d" .
	 Mi = "e" .
	 Fa = "f" .
	 So = "g" .
	 La = "a" .
	 Ti = ti .
	 ti = "b" .`,

	"Program = `\"` .",

	// Comments using the form "(* ... *)" should be accepted anywhere whitespace is allowed.
	`Program = (* top-level comment *) .`,
	`Program = "a" (* inline comment *) "b" .`,
	`(* file comment *) Program = "a" .`,

	// Additional tests exercising comments in many positions and with complex expressions
	`Program = ( "a" (* inside group *) "b" | "c" ) .`,
	`Program (*between name and eq*) = "x" .`,
	`Program = (* multi
line
comment *) "a" .`,
	`Program = "a" (*c*) … "z" .`,
	`Program = { (*c*) Note } .
	Note = "a" .`,
	`Program = ( "a" | "b" (*alt*) | "c" ) .`,
	`Program = "(* not a comment *)" .`,
	`Program = ( (*c*) "a" "b" ) .`,
	`Program = Song .
	 Song = { Note } .
	 Note = Do | (Re (*c*) | Mi | Fa | So | La) | Ti .
	 Do = "c" .
	 Re = "d" .
	 Mi = "e" .
	 Fa = "f" .
	 So = "g" .
	 La = "a" .
	 Ti = ti .
	 ti = "b" .`,

	// Ensure '(*' and '*)' sequences inside string literals are not treated as comments
	`Program = "contains (*" .`,
	`Program = "contains *)" .`,
	`Program = "x(*y" "z*)w" .`,
	`Program = "(*" ")*" .`,
	"Program = `(* raw inside backquotes *)` .",
}

var badGrammars = []string{
	`Program = | .`,
	`Program = | b .`,
	`Program = a … b .`,
	`Program = "a" … .`,
	`Program = … "b" .`,
	`Program = () .`,
	`Program = [] .`,
	`Program = {} .`,

	// Unterminated comment cases should fail
	`Program = (* unterminated comment .`,
	`Program = "a" (* missing end .`,
	`(* file comment without close Program = "a" .`,
}

func checkGood(t *testing.T, src string) {
	grammar, err := Parse("", bytes.NewBuffer([]byte(src)))
	if err != nil {
		t.Errorf("Parse(%s) failed: %v", src, err)
		return
	}
	if err = Verify(grammar, "Program"); err != nil {
		t.Errorf("Verify(%s) failed: %v", src, err)
	}
}

func checkBad(t *testing.T, src string) {
	_, err := Parse("", bytes.NewBuffer([]byte(src)))
	if err == nil {
		t.Errorf("Parse(%s) should have failed", src)
	}
}

func TestGrammars(t *testing.T) {
	for _, src := range goodGrammars {
		checkGood(t, src)
	}
	for _, src := range badGrammars {
		checkBad(t, src)
	}
}
