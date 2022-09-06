# goflagexample

Example of Golang CLI application with subcommands, positionals, and flags using standard library

#### Known issues:
Positional arguments should be last, it's a problem of package `flag` from standard library.

Documentation says: `Flag parsing stops just before the first non-flag argument ("-" is a non-flag argument) or after the terminator "--".`
