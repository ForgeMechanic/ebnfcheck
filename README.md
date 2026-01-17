# ebnfcheck
A simple tool to validate EBNF grammars via the CLI. ✅

## Dialects
This tool supports multiple EBNF dialects via the `-dialect` (or `-d`) flag. The following built-in dialects are available:

- **w3c** (default) — the W3C-style EBNF dialect (chosen as the default dialect).
- **iso** — ISO-style dialect.
- **go** — the Go `golang.org/x/exp/ebnf` compatible dialect (useful for compatibility testing).

You can select a dialect when validating a file:

```sh
# Use W3C dialect (default)
ebnfcheck -d w3c mygrammar.ebnf

# Use ISO dialect
ebnfcheck -d iso mygrammar.ebnf

# Use Go dialect
ebnfcheck -d go mygrammar.ebnf
```

## Building per-dialect binaries
If you'd like to produce separate downloadable binaries for each dialect (e.g., for distribution), you can build binaries whose default dialect is baked in at link time using `-ldflags`.

A convenient Makefile is provided with targets:

```sh
# build three binaries: ebnfcheck-go, ebnfcheck-iso, ebnfcheck-w3c
make build-all
```

Or build a single dialect binary:

```sh
make build-go
make build-iso
make build-w3c
```

These binaries will have their default `-dialect` value set at build time, so running `./ebnfcheck-w3c file.ebnf` will use the W3C dialect by default.

## Tests & Coverage
Run the test suite with coverage:

```sh
go test ./... -coverprofile=coverage
go tool cover -func=coverage
go tool cover -html=coverage -o coverage.html
```

## Extensibility
The registry design allows adding additional dialects later on (for example, enabling dialect-specific comment rules or verification behavior). Contributions welcome.
