package cli

import (
	"io"
	"os"

	"github.com/rsdoiel/osf"
)

/** Osf2txt is osf2txt: it reads an Open Screenplay Format document and writes
 * the screenplay as plain text.
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
 *   os.Exit(cli.Osf2txt("osf2txt", os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
 */
func Osf2txt(appName string, args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	var o options
	fset := newFlagSet(appName, stderr, &o, false)
	if code, done := start(appName, Osf2txtHelpText, fset, &o, args, stdout, stderr); done {
		return code
	}
	var screenplay *osf.OpenScreenplay
	if namedInput(&o) {
		if _, err := os.Stat(o.input); err != nil {
			return fail(appName, &o, stderr, classify(err), "%v", err)
		}
		var err error
		if screenplay, err = osf.ParseFile(o.input); err != nil {
			return fail(appName, &o, stderr, ExitData, "%s: %v", o.input, err)
		}
	} else {
		src, err := io.ReadAll(stdin)
		if err != nil {
			return fail(appName, &o, stderr, ExitIO, "%v", err)
		}
		if screenplay, err = osf.Parse(src); err != nil {
			return fail(appName, &o, stderr, ExitData, "%v", err)
		}
	}
	return write(appName, &o, []byte(screenplay.String()), stdout, stderr)
}
