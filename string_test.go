package osf

import (
	"strings"
	"testing"
)

// TestStringNilReceivers: every String method tolerates nil.
func TestStringNilReceivers(t *testing.T) {
	var (
		text  *Text
		para  *Para
		paras *Paragraphs
		tp    *TitlePage
		doc   *OpenScreenplay
	)
	for name, got := range map[string]string{
		"Text": text.String(), "Para": para.String(), "Paragraphs": paras.String(),
		"TitlePage": tp.String(), "OpenScreenplay": doc.String(),
	} {
		if got != "" {
			t.Errorf("nil %s: expected empty string, got %q", name, got)
		}
	}
}

// TestTextFormatting: each format attribute and their order of nesting.
func TestTextFormatting(t *testing.T) {
	for _, c := range []struct {
		name string
		text Text
		want string
	}{
		{"plain", Text{InnerText: "hi"}, "hi"},
		{"underline", Text{InnerText: "hi", Underline: UnderlineStyle}, "_hi_"},
		{"italic", Text{InnerText: "hi", Italic: ItalicStyle}, "*hi*"},
		{"bold", Text{InnerText: "hi", Bold: BoldStyle}, "**hi**"},
		{"bold italic", Text{InnerText: "hi", Bold: BoldStyle, Italic: ItalicStyle}, "***hi***"},
		{"all caps", Text{InnerText: "hi", AllCaps: AllCapsStyle}, "HI"},
		{"strikethrough", Text{InnerText: "hi", Strikethrough: StrikethroughStyle}, "~~hi~~"},
		{"blank text is not marked up", Text{InnerText: "  ", Bold: BoldStyle}, "  "},
		{"empty", Text{}, ""},
	} {
		if got := c.text.String(); got != c.want {
			t.Errorf("%s: expected %q, got %q", c.name, c.want, got)
		}
	}
}

// TestParaStyles: how each base style is rendered as Fountain text.
func TestParaStyles(t *testing.T) {
	for _, c := range []struct{ style, in, want string }{
		{GeneralType, "plain", "plain\n"},
		{SceneHeadingType, "int. house - day", "INT. HOUSE - DAY\n\n"},
		{ActionType, "She runs.", "She runs.\n\n"},
		{CharacterType, "Jane", "JANE\n"},
		{DialogueType, "Hello.", "Hello.\n\n"},
		{ParentheticalType, "softly", "(softly)\n"},
		{ParentheticalType, "(softly)", "(softly)\n"},
		{TransitionType, "cut to:", "CUT TO:\n\n"},
		{SingingType, "la la", "~la la\n"},
		{ShotType, "close on hand", "close on hand\n"},
	} {
		para := &Para{Style: &Style{BaseStyleName: c.style}, Text: []*Text{{InnerText: c.in}}}
		if got := para.String(); got != c.want {
			t.Errorf("%s %q: expected %q, got %q", c.style, c.in, c.want, got)
		}
	}
	noStyle := &Para{Text: []*Text{{InnerText: "a"}, {InnerText: "b"}}}
	if got := noStyle.String(); got != "ab\n" {
		t.Errorf("para without a style: got %q", got)
	}
}

// TestEmptyDocument: a document with nothing in it renders and round trips.
func TestEmptyDocument(t *testing.T) {
	doc := NewOpenScreenplay20()
	if doc.Version != "20" || doc.Type != "Open Screenplay Format document" {
		t.Errorf("NewOpenScreenplay20: %q %q", doc.Version, doc.Type)
	}
	if got := doc.String(); got != "" {
		t.Errorf("empty document String: %q", got)
	}
	src, err := doc.ToXML()
	if err != nil {
		t.Fatal(err)
	}
	back, err := Parse(src)
	if err != nil {
		t.Fatalf("%v\n%s", err, src)
	}
	if back.Version != "20" || back.Paragraphs != nil {
		t.Errorf("empty document changed in the round trip: %+v", back)
	}
}

// TestCleanupSelfClosingElements: empty elements become self-closing and
// non-empty ones are left alone.
func TestCleanupSelfClosingElements(t *testing.T) {
	in := `<a><text></text><titlepage></titlepage><text>x</text><para bookmark="b"></para></a>`
	want := `<a><text/><titlepage/><text>x</text><para bookmark="b"/></a>`
	if got := string(CleanupSelfClosingElements([]byte(in))); got != want {
		t.Errorf("expected %s, got %s", want, got)
	}
}

// TestFountainToTextRoundTrip: Fountain -> OSF -> text keeps the words,
// the order and the Fountain markers that String writes back.
func TestFountainToTextRoundTrip(t *testing.T) {
	doc, _ := convert(t, tinyScreenplay)
	got := doc.String()
	for _, want := range []string{
		"INT. OFFICE - NIGHT", "A PROGRAMMER stares at the screen.",
		"PROGRAMMER\n(drowsy)\nWhat algorithm is this?", "FADE TO BLACK.", "The sun rises.",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("text output lacks %q:\n%s", want, got)
		}
	}
	if strings.Index(got, "INT. OFFICE") > strings.Index(got, "EXT. COURTYARD") {
		t.Error("scenes are out of order")
	}
}
