//
// osf is a package for working with Open Screenplay Format 2.x XML documents.
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
//
// BSD 2-Clause License
//
// Copyright (c) 2017, R. S. Doiel
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
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
//
package osf

import (
	"encoding/xml"
	"fmt"
	"strings"
)

const (
	Version   = "v0.0.0-dev"
	DocString = `<?xml version="1.0" encoding="UTF-8" standalone="no" ?>`
)

var (
	// MaxLineWidth is the number of characters wide a line can be
	// based on a monospace font.
	MaxLineWidth = 80
)

type OpenScreenplay struct {
	XMLName    xml.Name `xml:"document"`
	Type       string   `xml:"type,attr"`
	Version    string   `xml:"version,attr"`
	Info       *Info
	Settings   *Settings
	Styles     *Styles
	Paragraphs *Paragraphs
	Spelling   *Spelling
	Lists      *Lists
	TitlePage  *TitlePage
}

type Info struct {
	XMLName     xml.Name `xml:"info"`
	UUID        string   `xml:"uuid,attr,omitempty"`
	Title       string   `xml:"title,attr,omitempty"`
	TitleFormat string   `xml:"title_format,attr,omitempty"`
	WrittenBy   string   `xml:"written_by,attr,omitempty"`
	Copyright   string   `xml:"copyright,attr,omitempty"`
	Contact     string   `xml:"contact,attr,omitempty"`
	Drafts      string   `xml:"drafts,attr,omitempty"`
	PageCount   string   `xml:"pagecount,attr,omitempty"`
}

type Settings struct {
	XMLName            xml.Name `xml:"settings"`
	PageWidth          string   `xml:"page_width,attr,omitempty"`
	PageHeight         string   `xml:"page_height,attr,omitempty"`
	MarginTop          string   `xml:"margin_top,attr,omitempty"`
	MarginBottom       string   `xml:"margin_bottom,attr,omitempty"`
	MarginLeft         string   `xml:"margin_left,attr,omitempty"`
	MarginRight        string   `xml:"margin_right,attr,omitempty"`
	NormalLinesPerInch string   `xml:"normal_linesperinch,attr,omitempty"`
	DialogueContinues  string   `xml:"dialogue_continues,attr,omitempty"`
	ContText           string   `xml:"cont_text,attr,omitempty"`
	MoreText           string   `xml:"more_text,attr,omitempty"`
	ContinuedText      string   `xml:"continued_text,attr,omitempty"`
	OmittedText        string   `xml:"omitted_text,attr,omitempty"`
	PageNumberFormat   string   `xml:"pagenumber_format,attr,omitempty"`
	PageNumberStart    string   `xml:"pagenumber_start,attr,omitempty"`
	PageNumberFirst    string   `xml:"pagenumber_first,attr,omitempty"`
	Revision           string   `xml:"revision,attr,omitempty"`
	ShowRevisions      string   `xml:"show_revisions,attr,omitempty"`
	SceneNumbering     string   `xml:"scene_numbering,attr,omitempty"`
	ScenesLocked       string   `xml:"scenes_locked,attr,omitempty"`
	PageNumbering      string   `xml:"page_numbering,attr,omitempty"`
	PagesLockeed       string   `xml:"pages_locked,attr,omitempty"`
}

type Styles struct {
	XMLName xml.Name `xml:"styles"`
	Style   []*Style
}

type Style struct {
	XMLName       xml.Name `xml:"style"`
	Name          string   `xml:"name,attr,omitempty"`
	Builtin       string   `xml:"builtin,attr,omitempty"`
	BuiltinIndex  string   `xml:"builtin_index,attr,omitempty"`
	Label         string   `xml:"label,attr,omitempty"`
	BaseStyleName string   `xml:"basestylename,attr,omitempty"`
	StyleEnter    string   `xml:"style_enter,attr,omitempty"`
	Font          string   `xml:"font,attr,omitempty"`
	Size          string   `xml:"size,attr,omitempty"`
	SpaceBefore   string   `xml:"spacebefore,attr,omitempty"`
	StyleTab      string   `xml:"style_tab,attr,omitempty"`
	KeepWithNext  string   `xml:"keepwithnext,attr,omitempty"`
	Effects       string   `xml:"effects,attr,omitempty"`
	LeftIdent     string   `xml:"leftindent,attr,omitempty"`
	RightIdent    string   `xml:"rightindent,attr,omitempty"`
	Align         string   `xml:"align,attr,omitempty"`
}

type Paragraphs struct {
	XMLName xml.Name `xml:"paragraphs"`
	Para    []*Para
}

type Para struct {
	XMLName    xml.Name `xml:"para"`
	PageNumber string   `xml:"page_number,attr,omitempty"`
	Bookmark   string   `xml:"bookmark,attr,omitempty"`
	Style      *Style
	Text       []*Text
	Marks      *Marks
}

type Text struct {
	XMLName   xml.Name `xml:"text"`
	InnerText string   `xml:",chardata"`
}

type Marks struct {
	XMLName xml.Name `xml:"marks"`
	Mark    []*Mark
}

type Mark struct {
	XMLName  xml.Name `xml:"mark"`
	At       string   `xml:"at,attr,omitempty"`
	Revision string   `xml:"revision,attr,omitempty"`
}

type Spelling struct {
	XMLName        xml.Name `xml:"spelling"`
	Language       string   `xml:"language,attr,omitempty"`
	UserDictionary *UserDictionary
}

type UserDictionary struct {
	XMLName xml.Name `xml:"user_dictionary"`
	Entry   []*Entry
}

type Entry struct {
	XMLName xml.Name `xml:"entry"`
	Word    string   `xml:"work,attr,omitempty"`
}

type Lists struct {
	XMLName        xml.Name         `xml:"lists"`
	Characters     []*Character     `xml:"characters"`
	Locations      []*Location      `xml:"locations"`
	SceneIntros    []*SceneIntro    `xml:"scene_intros"`
	SceneTimes     []*SceneTime     `xml:"scene_times"`
	Extensions     []*Extension     `xml:"extensions"`
	Transitions    []*Transition    `xml:"transitions"`
	RevisionColors []*RevisionColor `xml:"revision_colors"`
	TagCategories  []*TagCategory   `xml:"tag_categories"`
}

type Character struct {
	XMLName xml.Name `xml:"character"`
	Name    string   `xml:"name,attr,omitempty"`
}

type Location struct {
	XMLName xml.Name `xml:"location"`
	Name    string   `xml:"name,attr,omitempty"`
}

type SceneIntro struct {
	XMLName xml.Name `xml:"scene_intro"`
	Name    string   `xml:"name,attr,omitempty"`
}

type SceneTime struct {
	XMLName xml.Name `xml:"scene_time"`
	Name    string   `xml:"name,attr,omitempty"`
}

type Extension struct {
	XMLName xml.Name `xml:"extension"`
	Name    string   `xml:"name,attr,omitempty"`
}

type Transition struct {
	XMLName xml.Name `xml:"transition"`
	Name    string   `xml:"name,attr,omitempty"`
}

type RevisionColor struct {
	XMLName    xml.Name `xml:"revision_color"`
	Name       string   `xml:"name,attr,omitempty"`
	Index      string   `xml:"index,attr,omitempty"`
	ColorName  string   `xml:"color_name,attr,omitempty"`
	ColorIndex string   `xml:"color_index,attr,omitempty"`
}

type TagCategory struct {
	XMLName xml.Name `xml:"tag_category"`
	Name    string   `xml:"name,attr,omitempty"`
}

type TitlePage struct {
	XMLName xml.Name `xml:"titlepage"`
	Para    []*Para
}

func (doc *OpenScreenplay) String() string {
	src := []string{}
	if doc.Info != nil {
		if doc.TitlePage != nil {
			src = append(src, doc.TitlePage.String())
		}
		if doc.Paragraphs != nil {
			src = append(src, doc.Paragraphs.String())
		}
	}
	return strings.Join(src, "")
}

// Parse takes a byte array and returns a OpenScreenplay object and error
func Parse(src []byte) (*OpenScreenplay, error) {
	doc := new(OpenScreenplay)
	err := xml.Unmarshal(src, &doc)
	return doc, err
}

// ParseFile reads in *.osf and *.fadin file and and returns
// a OpenScreenplay object and error
func ParseFile(fname string) (*OpenScreenplay, err) {
	return nil, fmt.Errorf("ParseFile() not implemented")
}
