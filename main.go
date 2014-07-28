package main

import (
	"flag"
	git "github.com/libgit2/git2go"
	"os"
)

// refname="$1"
// oldrev="$2"
// newrev="$3"

// get files for newrev
// find yaml files
// parse yaml files

// exit code = 0 if ok
// exit code > 0 if failed
// ... print messages to stdout

func main() {
	refname := flag.Arg(1)
	oldrev := flag.Arg(2)
	newrev := flag.Arg(3)

	_ = refname
	_ = oldrev
	_ = newrev

	git.OpenRepository("/Users/mattes/Developer/developermail/gitd/tmprepos/peter/123")

	os.Exit(0)
}
