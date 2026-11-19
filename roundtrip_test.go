package osf

import (
	"path/filepath"
	"reflect"
	"testing"
)

/** sampleFiles returns the OSF and Fade In sample documents in testdata.
 *
 * Parameters:
 *   t (*testing.T) — the running test
 *
 * Returns:
 *   []string — paths of testdata/*.osf and testdata/*.fadein, never empty
 *
 * Example:
 *   for _, fname := range sampleFiles(t) { ... }
 */
func sampleFiles(t *testing.T) []string {
	t.Helper()
	var files []string
	for _, pattern := range []string{"*.osf", "*.fadein"} {
		matches, err := filepath.Glob(filepath.Join("testdata", pattern))
		if err != nil {
			t.Fatal(err)
		}
		files = append(files, matches...)
	}
	if len(files) == 0 {
		t.Fatal("no sample files in testdata")
	}
	return files
}

// TestRoundTrip: ParseFile -> ToXML -> Parse must give back an equal
// document for every sample, through both the plain XML and the zip paths.
func TestRoundTrip(t *testing.T) {
	for _, fname := range sampleFiles(t) {
		t.Run(filepath.Base(fname), func(t *testing.T) {
			first, err := ParseFile(fname)
			if err != nil {
				t.Fatalf("ParseFile(%q): %v", fname, err)
			}
			if first.Paragraphs == nil || len(first.Paragraphs.Para) == 0 {
				t.Fatalf("%s parsed with no paragraphs", fname)
			}
			src, err := first.ToXML()
			if err != nil {
				t.Fatalf("ToXML: %v", err)
			}
			second, err := Parse(src)
			if err != nil {
				t.Fatalf("Parse of ToXML output: %v\n%s", err, src)
			}
			if !reflect.DeepEqual(first, second) {
				t.Errorf("%s: document changed in the round trip", fname)
			}
			// A second render must be byte-identical (stable output).
			again, err := second.ToXML()
			if err != nil {
				t.Fatal(err)
			}
			if string(src) != string(again) {
				t.Errorf("%s: ToXML is not stable across a second pass", fname)
			}
		})
	}
}

// TestParseFileErrors covers the failure paths of ParseFile and Parse.
func TestParseFileErrors(t *testing.T) {
	if _, err := ParseFile(filepath.Join("testdata", "no-such.osf")); err == nil {
		t.Error("missing .osf: expected an error")
	}
	if _, err := ParseFile(filepath.Join("testdata", "no-such.fadein")); err == nil {
		t.Error("missing .fadein: expected an error")
	}
	if _, err := ParseFile(filepath.Join("testdata", "sample-01.fountain")); err == nil {
		t.Error("a Fountain file read as OSF: expected an error")
	}
	if _, err := Parse([]byte("<document><info")); err == nil {
		t.Error("truncated XML: expected an error")
	}
}
