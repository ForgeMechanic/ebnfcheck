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
	// Alias "default" to "go" for backward compatibility.
	if d, ok := GetDialect("go"); ok {
		RegisterDialect("default", d)
	}
}
