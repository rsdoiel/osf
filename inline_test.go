package osf

import (
	"fmt"
	"strings"
	"testing"
)

// runs renders text runs for comparison: plain text as is, formatted
// text as {flags:text}, e.g. {b:bold} {i:it} {u:under}.
func runs(texts []*Text) string {
	var out []string
	for _, tx := range texts {
		flags := ""
		if tx.Bold == BoldStyle {
			flags += "b"
		}
		if tx.Italic == ItalicStyle {
			flags += "i"
		}
		if tx.Underline == UnderlineStyle {
			flags += "u"
		}
		if flags == "" {
			out = append(out, tx.InnerText)
		} else {
			out = append(out, fmt.Sprintf("{%s:%s}", flags, tx.InnerText))
		}
	}
	return strings.Join(out, "|")
}

// TestStringToTextArray: inline Fountain emphasis becomes text runs.
func TestStringToTextArray(t *testing.T) {
	for _, c := range []struct{ in, want string }{
		{"plain", "plain"},
		{"", ""},
		{"a *it* b", "a |{i:it}| b"},
		{"a **bold** b", "a |{b:bold}| b"},
		{"a ***both*** b", "a |{bi:both}| b"},
		{"a _under_ b", "a |{u:under}| b"},
		{"**bold _and under_**", "{b:bold }|{bu:and under}"},
		{"*it **and bold** it*", "{i:it }|{bi:and bold}|{i: it}"},
		{"*start* middle *end*", "{i:start}| middle |{i:end}"},
		// not emphasis: stays literal
		{"5 * 3 = 15", "5 * 3 = 15"},
		{"an *unclosed marker", "an *unclosed marker"},
		{"snake_case_name", "snake_case_name"},
		{"a ** b **", "a ** b **"},
		{"** **", "** **"},
		// escapes
		{`literal \*star\* here`, "literal *star* here"},
		{`back\\slash`, `back\slash`},
		{`a \x b`, `a \x b`},
		// newlines are kept, in and out of runs
		{"line one\nline *two*", "line one\nline |{i:two}"},
	} {
		if got := runs(StringToTextArray(c.in)); got != c.want {
			t.Errorf("%q: expected %q, got %q", c.in, c.want, got)
		}
	}
}

// TestStringToTextArrayEmpty: empty input still gives one empty run, so a
// <para> always has a <text>.
func TestStringToTextArrayEmpty(t *testing.T) {
	if got := StringToTextArray(""); len(got) != 1 || got[0].InnerText != "" {
		t.Errorf("expected one empty run, got %v", runs(got))
	}
}

// TestInlineRoundTripThroughString: emphasis written by FromFountain is
// written back as Fountain by String.
func TestInlineRoundTripThroughString(t *testing.T) {
	doc, _ := convert(t, "INT. A - DAY\n\nHe *runs*, **shouts** and _hides_.\n")
	got := doc.String()
	if !strings.Contains(got, "He *runs*, **shouts** and _hides_.") {
		t.Errorf("emphasis lost: %q", got)
	}
}

// TestInlineSkipsNotesAndBoneyards: their asterisks are delimiters or text.
func TestInlineSkipsNotesAndBoneyards(t *testing.T) {
	doc, xmlSrc := convert(t, "INT. A - DAY\n\n/* a *boneyard* */\n\nText.\n")
	for _, para := range doc.Paragraphs.Para {
		for _, tx := range para.Text {
			if tx.Italic != "" || tx.Bold != "" {
				t.Errorf("formatting applied inside %q", tx.InnerText)
			}
		}
	}
	_ = xmlSrc
}

// TestInlineXMLRoundTrip: runs survive ToXML and Parse as attributes.
func TestInlineXMLRoundTrip(t *testing.T) {
	_, xmlSrc := convert(t, "INT. A - DAY\n\nHe *runs* **far**.\n")
	if !strings.Contains(xmlSrc, `<text italic="1">runs</text>`) || !strings.Contains(xmlSrc, `<text bold="1">far</text>`) {
		t.Errorf("XML lacks the attributes:\n%s", xmlSrc)
	}
	back, err := Parse([]byte(xmlSrc))
	if err != nil {
		t.Fatal(err)
	}
	if got := back.Paragraphs.Para[1].String(); got != "He *runs* **far**.\n\n" {
		t.Errorf("after XML round trip: %q", got)
	}
}

// FuzzStringToTextArray: the scanner always terminates, always returns a
// run, and never loses or invents a character other than the markers
// (*, _ and \) themselves.
func FuzzStringToTextArray(f *testing.F) {
	for _, seed := range []string{"", "*a*", "**a**", "***a***", "_a_", `\*a`, "a * b", "*a **b** c*", "_*a*_", "****", "__", "**_a_*"} {
		f.Add(seed)
	}
	strip := strings.NewReplacer("*", "", "_", "", `\`, "")
	f.Fuzz(func(t *testing.T, in string) {
		texts := StringToTextArray(in)
		if len(texts) == 0 {
			t.Fatal("no runs")
		}
		var joined strings.Builder
		for _, tx := range texts {
			joined.WriteString(tx.InnerText)
		}
		if strip.Replace(joined.String()) != strip.Replace(in) {
			t.Fatalf("characters changed: %q -> %q", in, joined.String())
		}
	})
}
