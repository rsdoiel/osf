// osf2txt is part of osf, a Go package for Open Screenplay Format. Run osf2txt -help.
package main

import (
	"os"
	"path"

	"github.com/rsdoiel/osf/internal/cli"
)

func main() {
	os.Exit(cli.Osf2txt(path.Base(os.Args[0]), os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
}
