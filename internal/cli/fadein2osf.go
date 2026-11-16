package cli

import (
	"encoding/xml"
	"io"
	"os"

	"github.com/rsdoiel/osf"
)

/** Fadein2osf is fadein2osf: it reads a Fade In file and writes an Open
 * Screenplay Format 2.0 document. A Fade In file is a zip archive, so the
 * input must be a named file, not standard input.
 *
 * Parameters:
 *   appName (string) — the program name, for help and error messages
 *   args ([]string) — the command line without the program name
 *   stdin (io.Reader) — standard input, which this program does not read
 *   stdout (io.Writer) — standard output
 *   stderr (io.Writer) — standard error
 *
 * Returns:
 *   int — the exit code
 *
 * Example:
 *   os.Exit(cli.Fadein2osf("fadein2osf", os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
 */
func Fadein2osf(appName string, args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	var o options
	fset := newFlagSet(appName, stderr, &o, false)
	if code, done := start(appName, Fadein2osfHelpText, fset, &o, args, stdout, stderr); done {
		return code
	}
	if !namedInput(&o) {
		return fail(appName, &o, stderr, ExitUsage, "missing a Fade In filename, e.g. %s -i screenplay.fadein", appName)
	}
	if _, err := os.Stat(o.input); err != nil {
		return fail(appName, &o, stderr, classify(err), "%v", err)
	}
	screenplay, err := osf.ParseFile(o.input)
	if err != nil {
		return fail(appName, &o, stderr, ExitData, "%s: %v", o.input, err)
	}
	src, err := xml.MarshalIndent(screenplay, " ", "    ")
	if err != nil {
		return fail(appName, &o, stderr, ExitInternal, "%v", err)
	}
	return write(appName, &o, osf.CleanupSelfClosingElements(src), stdout, stderr)
}
