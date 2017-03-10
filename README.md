
# fasta-example

### Purpose
**fasta-example** is a demonstration repository to implement a sequence assembly tool in Go.  The input is a file consisting of multiple overlapping sequences in [FASTA format](https://en.wikipedia.org/wiki/FASTA_format "FASTA format"), which are assembled into one complete sequence.  The sequences are not assumed to be in order, nor necessarily solvable.  A sequence set that cannot be completely assembled is considered an error, and reported as such.

## Build

A `Makefile` is provided that will build `fasta-example`, run tests, and produce a test coverage report.  `fasta-example` uses only Go stdlib functions, and requires no additional resources beyond a functioning `make` and a Go 1.7+ environment.

### Tests
A suite of unit tests is provided, with a coverage report.  This is expected to exercise all the components, with particular focus on the matching algorithm itself.

In addition, a set of blackbox live tests is included in the `scripts/test-fasta.sh` file, which tests the resulting binary for correct response in a variety of cases.  Live tests are used for flags, input files, usage errors, etc where unit testing is difficult or impossible. 

### Sample Build
```
<pforman@pforman: .../fasta-example > make
go build -o bin/fasta-example ./cmd
go test
PASS
ok  	github.com/pforman/fasta-example	0.007s
go test -coverprofile=c.out
PASS
coverage: 100.0% of statements
ok  	github.com/pforman/fasta-example	0.007s
go tool cover -func=c.out
github.com/pforman/fasta-example/assemble.go:10:	assemble		100.0%
github.com/pforman/fasta-example/debug.go:12:		SetDebug		100.0%
github.com/pforman/fasta-example/file.go:10:		AssembleFile		100.0%
github.com/pforman/fasta-example/file.go:23:		readFile		100.0%
github.com/pforman/fasta-example/file.go:46:		fragFromChunk		100.0%
github.com/pforman/fasta-example/file.go:69:		sanityCheckSequence	100.0%
github.com/pforman/fasta-example/match.go:9:		match			100.0%
github.com/pforman/fasta-example/match.go:29:		recurseMatch		100.0%
github.com/pforman/fasta-example/struct.go:14:		Error			100.0%
total:							(statements)		100.0%
./scripts/test-fasta.sh
Live test 1 complete: OK
Live test 2 complete: OK
Live test 3 complete: OK
Live test 4 complete: OK
Live test 5 complete: OK
Live test 6 complete: OK
Live test 7 complete: OK
Live test 8 complete: OK
```

## Run

The binary `fasta-example` produced by the Build step above is the only artifact of the build.  It can be packaged, or run directly from the `bin/` directory.  `make install` will install the binary to the `bin/` directory of **$GOPATH** if desired.

`fasta-example` requires exactly one argument, the file containing FASTA sequences to assemble.  Some flags are also included, primarily to format the output to FASTA format as necessary.

> The default output of `fasta-example` is not FASTA, but simply the complete sequence with no line breaks.  For FASTA format, use the **-t** and **-w** flags in combination.

### Usage
```
<pforman@pforman: .../fasta-example > ./fasta-example -h
Usage: ./fasta-example filename
  -d	Enable debug info
  -t string
    	Set FASTA title/description
  -w int
    	Wrap lines at specified length
```

**-d** prints debug info to **stderr**, which is primarily used to show the flow of the program, which is fully detailed below in **Algorithms**.

## Algorithms

`fasta-example` makes heavy use of the Go [strings](https://golang.org/pkg/strings/) package for sequence matching.  The primary function that determines if two fragments can be merged is assemble().

#### assemble()

assemble() uses the strings.Index() and strings.LastIndex() to search for a substring of a minimum length (threshold) to indicate a match.  Since the match may be longer than the minimum threshold, strings.HasPrefix() and strings.HasSuffix() are used to check if the two strings can be safely merged.

> Within this example, the threshold is determined as "at least half the fragment length", ie `(len(sequence)/2)+1`.   This is calculated outside of assemble(), not a limitation of the method.

#### match() and recurseMatch()

match() is a setup function that invokes recurseMatch to do the search and assembly.  It serves as a wrapper around any particular implementation, which in this case is provided by recurseMatch()

recurseMatch() is a simple recursive function that walks a list of fragments, attempting to place each onto a **base** sequence.  When a match is found, recurseMatch() is invoked with the placed fragment adding to **base** and removed from the list.

Placing all the fragments correctly results in an empty list, which is  the successful terminating condition for recurseMatch().  If the end of any list is reached without a successful placement of any remaining fragment, a gap in the sequence is evident, and an error is raised.  This serves as an unsuccessful terminating condition for recurseMatch()

#### Other functions

The remainder of the functions in the package deal primarily with parsing FASTA sequences from files, and checking input for correct formatting and setting up data structures that match(), recurseMatch(), and assemble() work with.

## Future Work

The implementation of sequence matching is likely naive, as a literature review was not performed before the package was written.  Timing shows that assembly of 50 fragments of approximately 1000 base pairs each is performed in approximately 30ms.  If this is insufficient, a domain-specific algorithm for matching sequences is likely to give better results, and should be cosnidered the first path to optimization.  The Go strings library was chosen for implementation due to its presence in the standard library, and the assumption that it performs basic string matching with close to optimal performance for the language runtime.
