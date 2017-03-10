package fastaexample

import (
	"log"
	"os"
)

var debug = false
var logger = log.New(os.Stderr, "debug: ", 0)

// SetDebug registers an internal variable to print debug information
func SetDebug(d bool) {
	debug = d
}
