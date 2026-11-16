package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// The tests call each program the way main does, through the same function,
// with real files from the module's testdata and the golden output the
// programs wrote before they left caltechlibrary/cli.

var programs = map[string]func(string, []string, *strings.Reader, *bytes.Buffer, *bytes.Buffer) int{
	"osf2txt": func(a string, args []string, in *strings.Reader, out, e *bytes.Buffer) int {
		return Osf2txt(a, args, in, out, e)
	},
	"txt2osf": func(a string, args []string, in *strings.Reader, out, e *bytes.Buffer) int {
		return Txt2osf(a, args, in, out, e)
	},
	"fadein2osf": func(a string, args []string, in *strings.Reader, out, e *bytes.Buffer) int {
		return Fadein2osf(a, args, in, out, e)
	},
}

func call(name string, stdin string, args ...string) (string, string, int) {
	var out, errOut bytes.Buffer
	code := programs[name](name, args, strings.NewReader(stdin), &out, &errOut)
	return out.String(), errOut.String(), code
}

func read(t *testing.T, name string) string {
	t.Helper()
	b, err := os.ReadFile(name)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func TestStandardOptionsInEverySpelling(t *testing.T) {
	for name := range programs {
		for _, opt := range []string{"-h", "-help", "--help"} {
			out, _, code := call(name, "", opt)
			if code != 0 || !strings.Contains(out, name+"(1) user manual") || !strings.Contains(out, "# EXIT STATUS") {
				t.Errorf("%s %s: exit %d\n%.300s", name, opt, code, out)
			}
			for _, tok := range []string{"{app_name}", "{version}", "{release_date}", "{release_hash}"} {
				if strings.Contains(out, tok) {
					t.Errorf("%s %s: token %s was not replaced", name, opt, tok)
				}
			}
		}
		for _, opt := range []string{"-v", "-version", "--version"} {
			out, _, code := call(name, "", opt)
			if code != 0 || !strings.HasPrefix(out, name+" 0.") {
				t.Errorf("%s %s: exit %d %q", name, opt, code, out)
			}
		}
		for _, opt := range []string{"-l", "-license", "--license"} {
			out, _, code := call(name, "", opt)
			if code != 0 || !strings.Contains(out, "GNU AFFERO GENERAL PUBLIC LICENSE") {
				t.Errorf("%s %s: exit %d %.100q", name, opt, code, out)
			}
		}
	}
}

func TestConversionsMatchWhatTheProgramsWroteBefore(t *testing.T) {
	cases := []struct{ prog, in, gold string }{
		{"osf2txt", "sample-01.osf", "osf2txt-01.golden"},
		{"osf2txt", "sample-02.osf", "osf2txt-02.golden"},
		{"txt2osf", "sample-01.fountain", "txt2osf-01.golden"},
		{"txt2osf", "sample-02.fountain", "txt2osf-02.golden"},
		{"fadein2osf", "sample-01.fadein", "fadein2osf-01.golden"},
		{"fadein2osf", "sample-02.fadein", "fadein2osf-02.golden"},
	}
	for _, c := range cases {
		in := filepath.Join("..", "..", "testdata", c.in)
		want := read(t, filepath.Join("testdata", c.gold))
		// -i, -input, and a positional name all read the file.
		for _, args := range [][]string{{"-i", in}, {"-input", in}, {in}} {
			out, errOut, code := call(c.prog, "", args...)
			if code != 0 || out != want || errOut != "" {
				t.Errorf("%s %v: exit %d, stderr %q, output differs from %s (got %d bytes, want %d)", c.prog, args, code, errOut, c.gold, len(out), len(want))
			}
		}
		if c.prog == "fadein2osf" {
			continue // a Fade In file is a zip archive and cannot come from standard input
		}
		out, _, code := call(c.prog, read(t, in), "-")
		if code != 0 || out != want {
			t.Errorf("%s from standard input: exit %d, output differs from %s", c.prog, code, c.gold)
		}
		out, _, code = call(c.prog, read(t, in))
		if code != 0 || out != want {
			t.Errorf("%s with no arguments: exit %d, output differs from %s", c.prog, code, c.gold)
		}
		// -o writes the same bytes to a file and nothing to standard output.
		dest := filepath.Join(t.TempDir(), "out")
		for _, flag := range []string{"-o", "-output"} {
			out, _, code = call(c.prog, "", "-i", in, flag, dest)
			if code != 0 || out != "" || read(t, dest) != want {
				t.Errorf("%s %s: exit %d, stdout %d bytes", c.prog, flag, code, len(out))
			}
		}
	}
}

func TestTrailingNewline(t *testing.T) {
	in := filepath.Join("..", "..", "testdata", "sample-01.osf")
	plain, _, _ := call("osf2txt", "", "-i", in)
	for _, flag := range []string{"-nl", "-newline"} {
		out, _, code := call("osf2txt", "", flag, "-i", in)
		if code != 0 || out != plain+"\n" {
			t.Errorf("osf2txt %s: exit %d, want the plain output and a newline", flag, code)
		}
	}
	// txt2osf adds the newline unless told not to; fadein2osf does not.
	fin := filepath.Join("..", "..", "testdata", "sample-01.fountain")
	with, _, _ := call("txt2osf", "", "-i", fin)
	without, _, code := call("txt2osf", "", "-nl=false", "-i", fin)
	if code != 0 || with != without+"\n" {
		t.Errorf("txt2osf -nl=false: exit %d, want the default output without its newline", code)
	}
}

func TestExitCodes(t *testing.T) {
	dir := t.TempDir()
	missing := filepath.Join(dir, "missing")
	noDir := filepath.Join(dir, "no", "such", "dir", "out")
	notZip := filepath.Join(dir, "garbage.fadein")
	if err := os.WriteFile(notZip, []byte("not a zip archive"), 0o644); err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		name  string
		prog  string
		stdin string
		args  []string
		want  int
	}{
		{"unknown option", "osf2txt", "", []string{"-bogus"}, 2},
		{"generate-markdown is gone", "osf2txt", "", []string{"-generate-markdown"}, 2},
		{"too many arguments", "txt2osf", "", []string{"a", "b", "c"}, 2},
		{"fadein2osf with no input", "fadein2osf", "", nil, 2},
		{"fadein2osf with - as input", "fadein2osf", "", []string{"-"}, 2},
		{"missing input", "osf2txt", "", []string{missing}, 66},
		{"missing input option", "txt2osf", "", []string{"-i", missing}, 66},
		{"missing fadein input", "fadein2osf", "", []string{missing + ".fadein"}, 66},
		{"malformed XML on standard input", "osf2txt", "<<<", nil, 65},
		{"not a zip archive", "fadein2osf", "", []string{notZip}, 65},
		{"uncreatable output", "osf2txt", "<document/>", []string{"-", noDir}, 73},
	}
	for _, c := range cases {
		out, errOut, code := call(c.prog, c.stdin, c.args...)
		if code != c.want {
			t.Errorf("%s: exit %d, want %d (stdout %.60q, stderr %q)", c.name, code, c.want, out, errOut)
			continue
		}
		if c.want != 0 {
			// One plain line on standard error, prefixed with the tool name (DR-0011).
			if !strings.HasPrefix(errOut, c.prog+": ") || strings.Count(strings.TrimRight(errOut, "\n"), "\n") != 0 {
				t.Errorf("%s: standard error is %q, want one line starting %q", c.name, errOut, c.prog+": ")
			}
			if out != "" {
				t.Errorf("%s: wrote to standard output: %.60q", c.name, out)
			}
		}
	}
}

func TestQuietSuppressesTheMessageNotTheExitCode(t *testing.T) {
	out, errOut, code := call("osf2txt", "", "-quiet", filepath.Join(t.TempDir(), "missing"))
	if code != 66 || errOut != "" || out != "" {
		t.Errorf("exit %d, stdout %q, stderr %q; want 66 and nothing printed", code, out, errOut)
	}
}

func TestBadInputDoesNotTruncateTheOutputFile(t *testing.T) {
	dest := filepath.Join(t.TempDir(), "keep.txt")
	if err := os.WriteFile(dest, []byte("keep me\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, _, code := call("osf2txt", "<<<", "-", dest); code != 65 {
		t.Fatalf("exit %d, want 65", code)
	}
	if got := read(t, dest); got != "keep me\n" {
		t.Errorf("a failed run replaced the output file with %q", got)
	}
}
