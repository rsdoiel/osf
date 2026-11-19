// osf is a package for working with Open Screenplay Format 1.2 and 2.0 XML documents.
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
//
// # BSD 2-Clause License
//
// Copyright (c) 2021, R. S. Doiel
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
//   - Redistributions of source code must retain the above copyright notice, this
//     list of conditions and the following disclaimer.
//
//   - Redistributions in binary form must reproduce the above copyright notice,
//     this list of conditions and the following disclaimer in the documentation
//     and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
package osf

import (
	"strings"

	// My Packages
	"github.com/rsdoiel/fountain"
)

/** baseStyleFor maps a Fountain element type name to the OSF base style
 * that Fade In and other OSF readers expect in <style basestylename="..."/>.
 *
 * Parameters:
 *   typeName (string) — the result of fountain.Element.TypeName()
 *
 * Returns:
 *   string — the OSF base style name
 *   bool — false for elements that are not paragraphs (blank lines and
 *          page feeds), which the caller leaves out
 *
 * Example:
 *   style, ok := baseStyleFor("Scene Heading") // "Scene Heading", true
 */
func baseStyleFor(typeName string) (string, bool) {
	switch typeName {
	case "Empty", "Page Feed":
		return "", false
	case "Scene Heading":
		return SceneHeadingType, true
	case "Action":
		return ActionType, true
	case "Character":
		return CharacterType, true
	case "Dialogue":
		return DialogueType, true
	case "Parenthetical":
		return ParentheticalType, true
	case "Transition":
		return TransitionType, true
	case "Lyric":
		return SingingType, true
	}
	// General Text, Note, Boneyard and anything else stay visible as
	// Normal Text; OSF has no better home for them.
	return GeneralType, true
}

/** stripForcedMarker removes the Fountain forcing marker that fountain keeps
 * in an element's Content, so OSF carries only the text.
 *
 * Parameters:
 *   typeName (string) — the result of fountain.Element.TypeName()
 *   content (string) — the element's Content as fountain.Parse gave it
 *
 * Returns:
 *   string — content without its marker: a leading "!" on Action, "@" on
 *            Character, "~" on Lyric, "." on Scene Heading (not an
 *            ellipsis), and a leading ">" and trailing "<" on Transition
 *
 * Example:
 *   stripForcedMarker("Transition", "> THE END <") // "THE END"
 */
func stripForcedMarker(typeName, content string) string {
	switch typeName {
	case "Action":
		return strings.TrimPrefix(content, "!")
	case "Character":
		return strings.TrimPrefix(content, "@")
	case "Lyric":
		return strings.TrimPrefix(content, "~")
	case "Scene Heading":
		if strings.HasPrefix(content, ".") && !strings.HasPrefix(content, "..") {
			return content[1:]
		}
	case "Transition":
		if strings.HasPrefix(content, ">") {
			content = strings.TrimSpace(strings.TrimPrefix(content, ">"))
			content = strings.TrimSpace(strings.TrimSuffix(content, "<"))
		}
	}
	return content
}

/** titleLines splits a title page value into its trimmed, non-empty lines.
 * fountain keeps the space after "Key:" and the indentation of a
 * multi-line value.
 *
 * Parameters:
 *   content (string) — an Element.Content from fountain's TitlePage
 *
 * Returns:
 *   []string — the lines, each trimmed, blank ones dropped
 *
 * Example:
 *   titleLines("\n\tContact\n\tinformation\n") // ["Contact", "information"]
 */
func titleLines(content string) []string {
	var lines []string
	for _, line := range strings.Split(content, "\n") {
		if line = strings.TrimSpace(line); line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}

func StringToTextArray(s string) []*Text {
	var a []*Text

	//FIXME: if string has embedded formatting split into parts and set format attributes
	// (e.g. text.Bold, text.Italic) as needed.
	text := new(Text)
	text.InnerText = s
	a = append(a, text)
	return a
}

// AddInfo parses a map[string]string and updates/adds Info struct to document as needed.
func (document *OpenScreenplay) FromFountain(screenplay *fountain.Fountain) {
	if screenplay.TitlePage != nil {
		// Build the Info section
		if document.Info == nil {
			document.Info = new(Info)
		}
		// Populate the TitlePage
		if document.TitlePage == nil {
			document.TitlePage = new(TitlePage)
		}

		// NOTE: build a map of elements that will belong to Info.
		for _, elem := range screenplay.TitlePage {
			lines := titleLines(elem.Content)
			one, block := strings.Join(lines, " "), strings.Join(lines, "\n")
			switch strings.ToLower(elem.Name) {
			case "title":
				document.Info.Title = one
				para := new(Para)
				para.Bookmark = "Title"
				para.Text = StringToTextArray(block)
				document.TitlePage.Para = append(document.TitlePage.Para, para)
			case "credit":
				document.Info.WrittenBy = one
				para := new(Para)
				para.Bookmark = "Credits"
				para.Text = StringToTextArray(block)
				document.TitlePage.Para = append(document.TitlePage.Para, para)
			case "author":
				document.Info.WrittenBy = one
				para := new(Para)
				para.Bookmark = "Author"
				para.Text = StringToTextArray(block)
				document.TitlePage.Para = append(document.TitlePage.Para, para)
			case "copyright":
				document.Info.Copyright = one
				para := new(Para)
				para.Bookmark = "Copyright"
				para.Text = StringToTextArray(block)
				document.TitlePage.Para = append(document.TitlePage.Para, para)
			case "draft date":
				document.Info.Drafts = one
				para := new(Para)
				para.Bookmark = "Drafts"
				para.Text = StringToTextArray(block)
				document.TitlePage.Para = append(document.TitlePage.Para, para)
			case "contact":
				document.Info.Contact = one
				para := new(Para)
				para.Bookmark = "Contact"
				para.Text = StringToTextArray(block)
				document.TitlePage.Para = append(document.TitlePage.Para, para)
			case "source":
				para := new(Para)
				para.Bookmark = "Source"
				para.Text = StringToTextArray(block)
				document.TitlePage.Para = append(document.TitlePage.Para, para)
			case "story by":
				para := new(Para)
				para.Bookmark = "Story By"
				para.Text = StringToTextArray(block)
				document.TitlePage.Para = append(document.TitlePage.Para, para)
			case "uuid":
				document.Info.UUID = one
			case "page_count":
				document.Info.PageCount = one
			case "title_format":
				document.Info.TitleFormat = one
			}
		}
	}

	if screenplay.Elements != nil {
		// Populate the Paragraphs array
		if document.Paragraphs == nil {
			document.Paragraphs = new(Paragraphs)
		}
		for _, elem := range screenplay.Elements {
			style, ok := baseStyleFor(elem.TypeName())
			if !ok {
				continue
			}
			para := new(Para)
			para.Style = &Style{BaseStyleName: style}
			para.Text = StringToTextArray(stripForcedMarker(elem.TypeName(), elem.Content))
			document.Paragraphs.Para = append(document.Paragraphs.Para, para)
		}
	}
}
