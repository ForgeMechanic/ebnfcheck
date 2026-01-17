package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestValidateEBNF_ValidFile(t *testing.T) {
	err := validateEBNF("testdata/valid.ebnf", "")
	if err != nil {
		t.Errorf("Expected no error for valid EBNF file, got: %v", err)
	}
}

func TestValidateEBNF_SimpleFile(t *testing.T) {
	err := validateEBNF("testdata/simple.ebnf", "")
	if err != nil {
		t.Errorf("Expected no error for simple EBNF file, got: %v", err)
	}
}

func TestValidateEBNF_InvalidFile(t *testing.T) {
	err := validateEBNF("testdata/invalid.ebnf", "")
	if err == nil {
		t.Error("Expected error for invalid EBNF file, got nil")
	}
}

func TestValidateEBNF_NonExistentFile(t *testing.T) {
	err := validateEBNF("testdata/nonexistent.ebnf", "")
	if err == nil {
		t.Error("Expected error for non-existent file, got nil")
	}
}

func TestValidateEBNF_WithValidStartRule(t *testing.T) {
	err := validateEBNF("testdata/valid.ebnf", "Program")
	if err != nil {
		t.Errorf("Expected no error for valid start rule, got: %v", err)
	}
}

func TestValidateEBNF_WithInvalidStartRule(t *testing.T) {
	err := validateEBNF("testdata/valid.ebnf", "NonExistent")
	if err == nil {
		t.Error("Expected error for invalid start rule, got nil")
	}
	expectedMsg := "start rule 'NonExistent' not found in grammar"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestValidateEBNF_WithStartRuleOnSimpleGrammar(t *testing.T) {
	err := validateEBNF("testdata/simple.ebnf", "SimpleGrammar")
	if err != nil {
		t.Errorf("Expected no error for valid start rule in simple grammar, got: %v", err)
	}
}

func TestValidateEBNF_EmptyStartRule(t *testing.T) {
	err := validateEBNF("testdata/valid.ebnf", "")
	if err != nil {
		t.Errorf("Expected no error when start rule is empty, got: %v", err)
	}
}

func TestValidateEBNF_WithTemporaryFile(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.ebnf")
	
	content := []byte("TestRule = \"test\" .\n")
	if err := os.WriteFile(tmpFile, content, 0644); err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	
	err := validateEBNF(tmpFile, "")
	if err != nil {
		t.Errorf("Expected no error for temporary valid EBNF file, got: %v", err)
	}
	
	err = validateEBNF(tmpFile, "TestRule")
	if err != nil {
		t.Errorf("Expected no error with valid start rule, got: %v", err)
	}
}
