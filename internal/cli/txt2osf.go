package cli

import (
	"io"
	"os"

	"github.com/rsdoiel/fountain"
	"github.com/rsdoiel/osf"
)

/** Txt2osf is txt2osf: it reads a Fountain or plain text screenplay and
 * writes an Open Screenplay Format 2.0 document.
 *
 * Parameters:
 *   appName (string) — the program name, for help and error messages
 *   args ([]string) — the command line without the program name
 *   stdin (io.Reader) — standard input
 *   stdout (io.Writer) — standard output
 *   stderr (io.Writer) — standard error
 *
 * Returns:
 *   int — the exit code
 *
 * Example:
 *   os.Exit(cli.Txt2osf("txt2osf", os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
 */
func Txt2osf(appName string, args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	var o options
	fset := newFlagSet(appName, stderr, &o, true)
	if code, done := start(appName, Txt2osfHelpText, fset, &o, args, stdout, stderr); done {
		return code
	}
	var src []byte
	var err error
	if namedInput(&o) {
		if src, err = os.ReadFile(o.input); err != nil {
			return fail(appName, &o, stderr, classify(err), "%v", err)
		}
	} else if src, err = io.ReadAll(stdin); err != nil {
		return fail(appName, &o, stderr, ExitIO, "%v", err)
	}
	screenplay, err := fountain.Parse(src)
	if err != nil {
		return fail(appName, &o, stderr, ExitData, "%v", err)
	}
	document := osf.NewOpenScreenplay20()
	document.FromFountain(screenplay)
	xmlSrc, err := document.ToXML()
	if err != nil {
		return fail(appName, &o, stderr, ExitInternal, "%v", err)
	}
	return write(appName, &o, xmlSrc, stdout, stderr)
}
