package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/forgemechanic/ebnfcheck/ebnf"
)

var dialectFlag string

func init() {
	flag.StringVar(&dialectFlag, "dialect", "default", "EBNF dialect to use (eg 'go')")
	flag.StringVar(&dialectFlag, "d", "default", "EBNF dialect to use (shorthand)")
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <path> [-start <rule>]\n", os.Args[0])
		flag.PrintDefaults()
	}

	startRule := flag.String("start", "", "optional start rule to verify")
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	filePath := args[0]

	if err := validateEBNF(filePath, *startRule); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println("OK")
}

func parseWithDialect(dialect string, filename string, src io.Reader) (ebnf.Grammar, error) {
	d, ok := ebnf.GetDialect(dialect)
	if !ok {
		return nil, fmt.Errorf("unknown dialect: %s", dialect)
	}
	g, err := d.Parse(filename, src)
	return g, err
}

func validateEBNF(filePath string, startRule string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	grammar, err := parseWithDialect(dialectFlag, filePath, file)
	if err != nil {
		return err
	}

	if startRule != "" {
		if _, exists := grammar[startRule]; !exists {
			return fmt.Errorf("start rule '%s' not found in grammar", startRule)
		}
	}

	return nil
}
