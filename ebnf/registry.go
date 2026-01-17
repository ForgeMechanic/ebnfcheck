package ebnf

import "io"

// Dialect represents a named EBNF dialect provider. The Parse function
// should parse from the provided filename and reader and return a Grammar.
type Dialect struct {
	Name  string
	Parse func(filename string, src io.Reader) (Grammar, error)
}

var dialects = map[string]Dialect{}

// RegisterDialect registers a dialect under the given name.
func RegisterDialect(name string, d Dialect) {
	dialects[name] = d
}

// GetDialect retrieves a registered dialect by name.
func GetDialect(name string) (Dialect, bool) {
	d, ok := dialects[name]
	return d, ok
}

func init() {
	// Register the "go" dialect to use the existing Parse implementation.
	RegisterDialect("go", Dialect{
		Name:  "go",
		Parse: func(filename string, src io.Reader) (Grammar, error) { return Parse(filename, src) },
	})

	// Register "w3c" and "iso" dialects. For now they use the same parser
	// implementation, but this registry allows adding dialect-specific
	// parsing behavior later (comment styles, lexical rules, verification
	// differences, etc.).
	RegisterDialect("w3c", Dialect{
		Name:  "w3c",
		Parse: func(filename string, src io.Reader) (Grammar, error) { return Parse(filename, src) },
	})
	RegisterDialect("iso", Dialect{
		Name:  "iso",
		Parse: func(filename string, src io.Reader) (Grammar, error) { return Parse(filename, src) },
	})

	// Alias "default" to "w3c" (chosen as the default dialect).
	if d, ok := GetDialect("w3c"); ok {
		RegisterDialect("default", d)
	}
}
