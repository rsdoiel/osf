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
	"archive/zip"
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"path"
	"strings"
)

const (
	Version   = "v0.0.0-dev"
	DocString = `<?xml version="1.0" encoding="UTF-8" standalone="no" ?>`

	// Style
	UnderlineStyle     = "1"
	ItalicStyle        = "1"
	BoldStyle          = "1"
	AllCapsStyle       = "1"
	StrikethroughStyle = "1"

	// Alignments
	CenterAlignment = "Center"
	LeftAlignment   = "Left"
	RightAlignment  = "Left"

	// Types used in ElementSettings and Paragraph elements
	GeneralType       = "Normal Text"
	SceneHeadingType  = "Scene Heading"
	ActionType        = "Action"
	CharacterType     = "Character"
	DialogueType      = "Dialogue"
	ParentheticalType = "Parenthetical"
	TransitionType    = "Transition"
	CastListType      = "Cast List"
	ShotType          = "Shot"
	SingingType       = "Singing"

	// DynamicLabel types
	PageNoType      = "Page #"
	LastRevisedType = "Last Revised"

	// Tabstop types
	RightType = "Right"
	LeftType  = "Left"
)

var (
	// MaxLineWidth is the number of characters wide a line can be
	// based on a monospace font.
	MaxLineWidth = 80
)

// OpenScreenplay holds the root structure for Unmarshaling OSF 1.2 and 2.0
type OpenScreenplay struct {
	XMLName    xml.Name    `xml:"document"`
	Type       string      `xml:"type,attr"`
	Version    string      `xml:"version,attr"`
	Info       *Info       `xml:"info"`
	Settings   *Settings   `xml:"settings"`
	Styles     *Styles     `xml:"styles,omitempty"`
	Paragraphs *Paragraphs `xml:"paragraphs"`
	Spelling   *Spelling   `xml:"spelling,omitempty"`
	Lists      *Lists      `xml:"lists"`
	TitlePage  *TitlePage  `xml:"titlepage,omitempty"`
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
	Para    []*Para  `xml:"para,omitempty"`
}

type Para struct {
	XMLName    xml.Name `xml:"para"`
	PageNumber string   `xml:"page_number,attr,omitempty"`
	Bookmark   string   `xml:"bookmark,attr,omitempty"`
	Style      *Style   `xml:"style,omitempty"`
	Text       []*Text  `xml:"text,omitempty"`
	Marks      *Marks   `xml:"marks,omitempty"`
}

type Text struct {
	XMLName       xml.Name `xml:"text"`
	Underline     string   `xml:"underline,attr,omitempty"`
	Italic        string   `xml:"italic,attr,omitempty"`
	Bold          string   `xml:"bold,attr,omitempty"`
	Strikethrough string   `xml:"strikethrough,attr,omitempty"`
	AllCaps       string   `xml:"allcaps,attr,omitempty"`
	InnerText     string   `xml:",chardata"`
}

type Marks struct {
	XMLName xml.Name `xml:"marks"`
	Mark    []*Mark  `xml:"mark,omitempty"`
}

type Mark struct {
	XMLName  xml.Name `xml:"mark"`
	At       string   `xml:"at,attr,omitempty"`
	Revision string   `xml:"revision,attr,omitempty"`
}

type Spelling struct {
	XMLName        xml.Name        `xml:"spelling"`
	Language       string          `xml:"language,attr,omitempty"`
	UserDictionary *UserDictionary `xml:"user_dictionary,omitempty"`
}

type UserDictionary struct {
	XMLName xml.Name `xml:"user_dictionary"`
	Entry   []*Entry `xml:"entry,omitempty"`
}

type Entry struct {
	XMLName xml.Name `xml:"entry"`
	Word    string   `xml:"work,attr,omitempty"`
}

type Lists struct {
	XMLName        xml.Name        `xml:"lists"`
	Characters     *Characters     `xml:"characters,omitempty"`
	Locations      *Locations      `xml:"locations,omitempty"`
	SceneIntros    *SceneIntros    `xml:"scene_intros,omitempty"`
	SceneTimes     *SceneTimes     `xml:"scene_times,omitempty"`
	Extensions     *Extensions     `xml:"extensions,omitempty"`
	Transitions    *Transitions    `xml:"transitions,omitempty"`
	RevisionColors *RevisionColors `xml:"revision_colors,omitempty"`
	TagCategories  *TagCategories  `xml:"tag_categories,omitempty"`
}

type Characters struct {
	XMLName   xml.Name     `xml:"characters"`
	Character []*Character `xml:"character,omitempty"`
}

type Character struct {
	XMLName xml.Name `xml:"character"`
	Name    string   `xml:"name,attr,omitempty"`
}

type Locations struct {
	XMLName  xml.Name    `xml:"locations"`
	Location []*Location `xml:"location,omitempty"`
}

type Location struct {
	XMLName xml.Name `xml:"location"`
	Name    string   `xml:"name,attr,omitempty"`
}

type SceneIntros struct {
	XMLName    xml.Name      `xml:"scene_intros"`
	SceneIntro []*SceneIntro `xml:"scene_intro,omitempty"`
}

type SceneIntro struct {
	XMLName xml.Name `xml:"scene_intro"`
	Name    string   `xml:"name,attr,omitempty"`
}

type SceneTimes struct {
	XMLName   xml.Name     `xml:"scene_times"`
	SceneTime []*SceneTime `xml:"scene_time,omitempty"`
}

type SceneTime struct {
	XMLName xml.Name `xml:"scene_time"`
	Name    string   `xml:"name,attr,omitempty"`
}

type Extensions struct {
	XMLName   xml.Name     `xml:"extensions"`
	Extension []*Extension `xml:"extension,omitempty"`
}

type Extension struct {
	XMLName xml.Name `xml:"extension"`
	Name    string   `xml:"name,attr,omitempty"`
}

type Transitions struct {
	XMLName    xml.Name      `xml:"transitions"`
	Transition []*Transition `xml:"transition,omitempty"`
}

type Transition struct {
	XMLName xml.Name `xml:"transition"`
	Name    string   `xml:"name,attr,omitempty"`
}

type RevisionColors struct {
	XMLName       xml.Name         `xml:"revision_colors"`
	RevisionColor []*RevisionColor `xml:"revision_color,omitempty"`
}

type RevisionColor struct {
	XMLName    xml.Name `xml:"revision_color"`
	Name       string   `xml:"name,attr,omitempty"`
	Index      string   `xml:"index,attr,omitempty"`
	ColorName  string   `xml:"color_name,attr,omitempty"`
	ColorIndex string   `xml:"color_index,attr,omitempty"`
}

type TagCategories struct {
	XMLName     xml.Name       `xml:"tag_categories"`
	TagCategory []*TagCategory `xml:"tag_category,omitempty"`
}

type TagCategory struct {
	XMLName xml.Name `xml:"tag_category"`
	Name    string   `xml:"name,attr,omitempty"`
}

type TitlePage struct {
	XMLName xml.Name `xml:"titlepage"`
	Para    []*Para  `xml:"para,omitempty"`
}

func (text *Text) String() string {
	if text != nil {
		s := text.InnerText
		if strings.TrimSpace(s) != "" {
			if text.Underline == UnderlineStyle {
				s = "_" + s + "_"
			}
			if text.Italic == ItalicStyle {
				s = "*" + s + "*"
			}
			if text.Bold == BoldStyle {
				s = "**" + s + "**"
			}
			if text.AllCaps == AllCapsStyle {
				s = strings.ToUpper(s)
			}
			if text.Strikethrough == StrikethroughStyle {
				s = "~~" + s + "~~"
			}
		}
		return s
	}
	return ""
}

func (para *Para) String() string {
	if para != nil {
		src := []string{}
		for _, text := range para.Text {
			s := text.String()
			if para.Style != nil {
				switch para.Style.BaseStyleName {
				case GeneralType:
					//s = s + "\n"
				case SceneHeadingType:
					s = strings.ToUpper(s) + "\n"
				case ActionType:
					s = s + "\n"
				case CharacterType:
					s = strings.ToUpper(s)
				case DialogueType:
					s = s + "\n"
				case ParentheticalType:
					if strings.HasPrefix(s, "(") == false && strings.HasSuffix(s, ")") == false {
						s = "(" + s + ")"
					}
				case TransitionType:
					s = strings.ToUpper(s) + "\n"
				case SingingType:
					s = "~" + s + "~"
				}
			}
			//FIXME: Apply formatting, FF and LF as needed
			src = append(src, s)
		}
		return strings.Join(src, "") + "\n"
	}
	return ""
}

func (paragraphs *Paragraphs) String() string {
	if paragraphs != nil {
		src := []string{}
		for _, para := range paragraphs.Para {
			s := para.String()
			// FIXME: Apply formatting, FF, LF as needed
			src = append(src, s)
		}
		return strings.Join(src, "")
	}
	return ""
}

func (tp *TitlePage) String() string {
	if tp != nil {
		src := []string{}
		for _, para := range tp.Para {
			s := para.String()
			//FIXME: Applyformatting, FF, LF as needed
			src = append(src, s)
		}
		return strings.Join(src, "")
	}
	return ""
}

func (doc *OpenScreenplay) String() string {
	if doc != nil {
		src := []string{}
		if doc.TitlePage != nil {
			src = append(src, doc.TitlePage.String())
		}
		if doc.Paragraphs != nil {
			src = append(src, doc.Paragraphs.String())
		}
		return strings.Join(src, "")
	}
	return ""
}

// Parse takes a byte array and returns a OpenScreenplay object and error
func Parse(src []byte) (*OpenScreenplay, error) {
	doc := new(OpenScreenplay)
	err := xml.Unmarshal(src, &doc)
	return doc, err
}

// ParseFile reads in *.osf and *.fadin file and and returns
// a OpenScreenplay object and error
func ParseFile(fname string) (*OpenScreenplay, error) {
	var (
		src []byte
		ext string
		err error
	)
	src = []byte{}
	ext = path.Ext(fname)
	if strings.ToLower(ext) == ".fadein" {
		// Open a zip archive for reading.
		r, err := zip.OpenReader(fname)
		if err != nil {
			return nil, err
		}
		defer r.Close()

		// Iterate through the files in the archive,
		// printing some of their contents.
		for _, f := range r.File {
			if f.Name == "document.xml" {
				rc, err := f.Open()
				if err != nil {
					return nil, err
				}
				src, err = ioutil.ReadAll(rc)
				if err != nil {
					return nil, err
				}
				rc.Close()
				break
			}
		}
	} else {
		src, err = ioutil.ReadFile(fname)
		if err != nil {
			return nil, err
		}
	}
	//FIXME: Need to sniff version, 1.2 and 2.0 probably can use the same structs but
	// 2.1 uses camel case for element names
	return Parse(src)
}

// CleanupSelfClosingElements changes something like <styles></styles> to <styles/>
func CleanupSelfClosingElements(src []byte) []byte {
	for _, elem := range []string{"info", "settings", "styles", "style", "mark", "text", "entry", "character", "location", "scene_time", "extension", "revision_color", "tag_category", "transition", "spelling", "user_dictionary", "paragraphs", "para", "locations"} {
		src = bytes.Replace(src, []byte("></"+elem+">"), []byte("/>"), -1)
	}
	return src
}
