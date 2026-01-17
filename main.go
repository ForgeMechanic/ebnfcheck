package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/forgemechanic/ebnfcheck/ebnf"
)

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

func validateEBNF(filePath string, startRule string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	grammar, err := ebnf.Parse(filePath, file)
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
