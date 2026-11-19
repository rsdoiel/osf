package osf

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/rsdoiel/fountain"
)

const tinyScreenplay = `Title: Tiny

INT. OFFICE - NIGHT

A PROGRAMMER stares at the screen.

PROGRAMMER
(drowsy)
What algorithm is this?

> FADE TO BLACK.

EXT. COURTYARD - DAY

The sun rises.
`

/** convert runs Fountain source through FromFountain and ToXML.
 *
 * Parameters:
 *   t (*testing.T) — the running test
 *   src (string) — Fountain source
 *
 * Returns:
 *   *OpenScreenplay — the converted document
 *   string — its XML
 *
 * Example:
 *   doc, xml := convert(t, tinyScreenplay)
 */
func convert(t *testing.T, src string) (*OpenScreenplay, string) {
	t.Helper()
	screenplay, err := fountain.Parse([]byte(src))
	if err != nil {
		t.Fatalf("fountain.Parse: %v", err)
	}
	doc := NewOpenScreenplay20()
	doc.FromFountain(screenplay)
	xmlSrc, err := doc.ToXML()
	if err != nil {
		t.Fatalf("ToXML: %v", err)
	}
	return doc, string(xmlSrc)
}

// TestFromFountainStyles: Fade In and OSF carry a paragraph's kind in
// <style basestylename="..."/>, as every testdata/*.osf shows. A bookmark
// attribute is only a label and Fade In does not read it as a style.
func TestFromFountainStyles(t *testing.T) {
	doc, xmlSrc := convert(t, tinyScreenplay)
	want := []string{SceneHeadingType, ActionType, CharacterType, ParentheticalType, DialogueType, TransitionType, SceneHeadingType, ActionType}
	var got []string
	for _, para := range doc.Paragraphs.Para {
		if para.Style == nil {
			t.Errorf("paragraph %q has no <style>", para.String())
			continue
		}
		got = append(got, para.Style.BaseStyleName)
	}
	if strings.Join(got, ",") != strings.Join(want, ",") {
		t.Errorf("styles: expected %v, got %v", want, got)
	}
	for _, s := range want {
		if !strings.Contains(xmlSrc, `<style basestylename="`+s+`"/>`) {
			t.Errorf("XML lacks <style basestylename=%q/>", s)
		}
	}
	if strings.Contains(xmlSrc, `bookmark="Scene Heading"`) {
		t.Error("paragraph type is still written as a bookmark")
	}
}

// TestFromFountainNoEmptyParagraphs: Fountain blank lines are separators,
// not paragraphs; the Fade In samples have none.
func TestFromFountainNoEmptyParagraphs(t *testing.T) {
	doc, xmlSrc := convert(t, tinyScreenplay)
	for _, para := range doc.Paragraphs.Para {
		if strings.TrimSpace(para.String()) == "" {
			t.Errorf("empty paragraph in output: %+v", para)
		}
	}
	if strings.Contains(xmlSrc, "Empty") {
		t.Error("XML still carries Empty paragraphs")
	}
}

// TestFromFountainSamples: every sample converts, round trips through
// XML, and keeps only styles that appear in the real Fade In samples.
func TestFromFountainSamples(t *testing.T) {
	known := map[string]bool{GeneralType: true, SceneHeadingType: true, ActionType: true,
		CharacterType: true, DialogueType: true, ParentheticalType: true, TransitionType: true}
	files, _ := filepath.Glob(filepath.Join("testdata", "sample-0*.fountain"))
	if len(files) == 0 {
		t.Fatal("no Fountain samples")
	}
	for _, fname := range files {
		t.Run(filepath.Base(fname), func(t *testing.T) {
			src, err := readFile(fname)
			if err != nil {
				t.Fatal(err)
			}
			doc, xmlSrc := convert(t, src)
			if len(doc.Paragraphs.Para) == 0 {
				t.Fatal("no paragraphs")
			}
			for _, para := range doc.Paragraphs.Para {
				if para.Style == nil || !known[para.Style.BaseStyleName] {
					t.Errorf("paragraph %q has style %+v", para.String(), para.Style)
				}
			}
			if _, err := Parse([]byte(xmlSrc)); err != nil {
				t.Errorf("output does not parse: %v", err)
			}
		})
	}
}

func readFile(fname string) (string, error) {
	b, err := os.ReadFile(fname)
	return string(b), err
}

// TestFromFountainForcedMarkers: Fountain's forcing markers tell the parser
// what an element is; they are not part of the text OSF should carry.
func TestFromFountainForcedMarkers(t *testing.T) {
	const src = "!FADE IN:\n\n.FORCED HEADING\n\n@McCLANE\nYippee.\n\n~Happy birthday to you\n\n> FADE TO BLACK.\n\n> THE END <\n\nINT. REAL - DAY\n\n...and so on.\n"
	doc, _ := convert(t, src)
	var got []string
	for _, para := range doc.Paragraphs.Para {
		got = append(got, para.Style.BaseStyleName+"|"+strings.TrimSuffix(para.Text[0].InnerText, "\n"))
	}
	want := []string{
		"Action|FADE IN:",
		"Scene Heading|FORCED HEADING",
		"Character|McCLANE",
		"Dialogue|Yippee.",
		"Singing|Happy birthday to you",
		"Transition|FADE TO BLACK.",
		"Transition|THE END",
		"Scene Heading|INT. REAL - DAY",
		// fountain calls an ellipsis line a scene heading; the dots are text
		"Scene Heading|...and so on.",
	}
	if strings.Join(got, "\n") != strings.Join(want, "\n") {
		t.Errorf("expected:\n%s\ngot:\n%s", strings.Join(want, "\n"), strings.Join(got, "\n"))
	}
}

// TestFromFountainTitlePage: fountain keeps the whitespace after "Key:" and
// the indentation of multi-line values; OSF should get clean text.
func TestFromFountainTitlePage(t *testing.T) {
	const src = "Title: TITLE\nAuthor: Author's Name\nContact:\n\tContact\n\tinformation\n\nDraft date:\n\tDraft\n\n\tinformation\n\nINT. A - DAY\n\nText.\n"
	doc, xmlSrc := convert(t, src)
	info := doc.Info
	for _, c := range []struct{ name, got, want string }{
		{"Info.Title", info.Title, "TITLE"},
		{"Info.WrittenBy", info.WrittenBy, "Author's Name"},
		{"Info.Contact", info.Contact, "Contact information"},
		{"Info.Drafts", info.Drafts, "Draft information"},
	} {
		if c.got != c.want {
			t.Errorf("%s: expected %q, got %q", c.name, c.want, c.got)
		}
	}
	want := map[string]string{
		"Title":   "TITLE",
		"Author":  "Author's Name",
		"Contact": "Contact\ninformation",
		"Drafts":  "Draft\ninformation",
	}
	for _, para := range doc.TitlePage.Para {
		if w, ok := want[para.Bookmark]; ok && para.Text[0].InnerText != w {
			t.Errorf("title page %s: expected %q, got %q", para.Bookmark, w, para.Text[0].InnerText)
		}
	}
	if strings.Contains(xmlSrc, "&#x9;") || strings.Contains(xmlSrc, "\t\t<text>\n") {
		t.Error("tabs from the Fountain source leaked into the XML")
	}
	back, err := Parse([]byte(xmlSrc))
	if err != nil {
		t.Fatal(err)
	}
	if back.TitlePage.Para[2].Text[0].InnerText != "Contact\ninformation" {
		t.Errorf("multi-line value did not survive XML: %q", back.TitlePage.Para[2].Text[0].InnerText)
	}
}

// TestFromFountainTitlePageKeys: every title page key fountain hands over
// lands in the Info element or a titled paragraph.
func TestFromFountainTitlePageKeys(t *testing.T) {
	const src = "Title: T\nCredit: Written by\nAuthor: A\nSource: S\nStory by: B\nCopyright: (c)\nDraft date: D\nContact: C\nUUID: u-1\nPage_count: 9\nTitle_format: <u>\nUnknown key: ignored\n\nINT. A - DAY\n\nText.\n"
	doc, _ := convert(t, src)
	info := doc.Info
	if info.Title != "T" || info.Copyright != "(c)" || info.Drafts != "D" || info.Contact != "C" ||
		info.UUID != "u-1" || info.PageCount != "9" || info.TitleFormat != "<u>" {
		t.Errorf("Info: %+v", info)
	}
	got := map[string]string{}
	for _, para := range doc.TitlePage.Para {
		got[para.Bookmark] = para.Text[0].InnerText
	}
	for k, v := range map[string]string{"Title": "T", "Credits": "Written by", "Author": "A", "Source": "S",
		"Story By": "B", "Copyright": "(c)", "Drafts": "D", "Contact": "C"} {
		if got[k] != v {
			t.Errorf("title page %q: expected %q, got %q", k, v, got[k])
		}
	}
	if len(got) != 8 {
		t.Errorf("expected 8 title page paragraphs, got %v", got)
	}
}
