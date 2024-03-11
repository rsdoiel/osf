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
	"archive/zip"
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"path"
	"strings"
)

const (
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
	RightAlignment  = "Right"

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
	XMLName    xml.Name    `xml:"document" json:"-" yaml:"-"`
	Type       string      `xml:"type,attr" json:"document_type" yaml:"document_type"`
	Version    string      `xml:"version,attr" json:"version" yaml:"version"`
	Info       *Info       `xml:"info" json:"info,omitempty" yaml:"info,omitempty"`
	Settings   *Settings   `xml:"settings" json:"settings" yaml:"settings"`
	Styles     *Styles     `xml:"styles,omitempty" json:"styles,omitempty" yaml:"styles,omitempty"`
	Paragraphs *Paragraphs `xml:"paragraphs" json:"paragraphs" yaml:"paragraphs"`
	Spelling   *Spelling   `xml:"spelling,omitempty" json:"spelling,omitempty" yaml:"spelling,omitempty"`
	Lists      *Lists      `xml:"lists" json:"lists,omitempty" yaml:"lists,omitempty"`
	TitlePage  *TitlePage  `xml:"titlepage,omitempty" json:"title_page,omitempty" yaml:"title_page,omitempty"`
}

type Info struct {
	XMLName     xml.Name `xml:"info" json:"-" yaml:"-"`
	UUID        string   `xml:"uuid,attr,omitempty" json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Title       string   `xml:"title,attr,omitempty" json:"title,omitempty" yaml:"title,omitempty"`
	TitleFormat string   `xml:"title_format,attr,omitempty" json:"title_format,omitempty" yaml:"title_format,omitempty"`
	WrittenBy   string   `xml:"written_by,attr,omitempty" json:"written_by,omitempty" yaml:"written_by,omitempty"`
	Copyright   string   `xml:"copyright,attr,omitempty" json:"copyright,omitempty" yaml:"copyright,omitempty"`
	Contact     string   `xml:"contact,attr,omitempty" json:"contact,omitempty" yaml:"contact,omitempty"`
	Drafts      string   `xml:"drafts,attr,omitempty" json:"drafts,omitempty" yaml:"drafts,omitempty"`
	PageCount   string   `xml:"pagecount,attr,omitempty" json:"page_count,omitempty" yaml:"page_count,omitempty"`
}

type Settings struct {
	XMLName            xml.Name `xml:"settings" json:"-" yaml:"-"`
	PageWidth          string   `xml:"page_width,attr,omitempty" json:"page_width,omitempty" yaml:"page_width,omitempty"`
	PageHeight         string   `xml:"page_height,attr,omitempty" json:"page_height,omitempty" yaml:"page_height,omitempty"`
	MarginTop          string   `xml:"margin_top,attr,omitempty" json:"margin_top,omitempty" yaml:"margin_top,omitempty"`
	MarginBottom       string   `xml:"margin_bottom,attr,omitempty" json:"margin_bottom,omitempty" yaml:"margin_bottom,omitempty"`
	MarginLeft         string   `xml:"margin_left,attr,omitempty" json:"margin_left,omitempty" yaml:"margin_left,omitempty"`
	MarginRight        string   `xml:"margin_right,attr,omitempty" json:"margin_right,omitempty" yaml:"margin_right,omitempty"`
	NormalLinesPerInch string   `xml:"normal_linesperinch,attr,omitempty" json:"normal_lines_per_inch,omitempty" yaml:"normal_lines_per_inch,omitempty"`
	DialogueContinues  string   `xml:"dialogue_continues,attr,omitempty" json:"dialog_continues,omitempty" yaml:"dialog_continues,omitempty"`
	ContText           string   `xml:"cont_text,attr,omitempty" json:"cont_text,omitempty" yaml:"cont_text,omitempty"`
	MoreText           string   `xml:"more_text,attr,omitempty" json:"more_text,omitempty" yaml:"more_text,omitempty"`
	ContinuedText      string   `xml:"continued_text,attr,omitempty" json:"continued_text,omitempty" yaml:"continued_text,omitempty"`
	OmittedText        string   `xml:"omitted_text,attr,omitempty" json:"omitted_text,omitempty" yaml:"omitted_text,omitempty"`
	PageNumberFormat   string   `xml:"pagenumber_format,attr,omitempty" json:"page_number_format,omitempty" yaml:"page_number_format,omitempty"`
	PageNumberStart    string   `xml:"pagenumber_start,attr,omitempty" json:"page_number_start,omitempty" yaml:"page_number_start,omitempty"`
	PageNumberFirst    string   `xml:"pagenumber_first,attr,omitempty" json:"page_number_first,omitempty" yaml:"page_number_first,omitempty"`
	Revision           string   `xml:"revision,attr,omitempty" json:"revision,omitempty" yaml:"revision,omitempty"`
	ShowRevisions      string   `xml:"show_revisions,attr,omitempty" json:"show_revisions,omitempty" yaml:"show_reviserions,omitempty"`
	SceneNumbering     string   `xml:"scene_numbering,attr,omitempty" json:"scene_numbering,omitempty" yaml:"scene_numbering,omitempty"`
	ScenesLocked       string   `xml:"scenes_locked,attr,omitempty" json:"scenes_locked,omitempty" yaml:"scenes_locked,omitempty"`
	PageNumbering      string   `xml:"page_numbering,attr,omitempty" json:"page_numbering,omitempty" yaml:"page_numbering,omitempty"`
	PagesLocked        string   `xml:"pages_locked,attr,omitempty" json:"pages_locked,omitempty" yaml:"pages_locked,omitempty"`
}

type Styles struct {
	XMLName xml.Name `xml:"styles" json:"styles" yaml:"styles"`
	Style   []*Style
}

type Style struct {
	XMLName       xml.Name `xml:"style" json:"-" yaml:"-"`
	Name          string   `xml:"name,attr,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Builtin       string   `xml:"builtin,attr,omitempty" json:"builtin,omitempty" yaml:"builtin,omitempty"`
	BuiltinIndex  string   `xml:"builtin_index,attr,omitempty" json:"builtin_index,omitempty" yaml:"builtin_index,omitempty"`
	Label         string   `xml:"label,attr,omitempty" json:"label,omitempty" yaml:"label,omitempty"`
	BaseStyleName string   `xml:"basestylename,attr,omitempty" json:"basestylename,omitempty" yaml:"basestylename,omitempty"`
	StyleEnter    string   `xml:"style_enter,attr,omitempty" json:"style_enter,omitempty" yaml:"style_enter,omitempty"`
	Font          string   `xml:"font,attr,omitempty" json:"font,omitempty" yaml:"font,omitempty"`
	Size          string   `xml:"size,attr,omitempty" json:"size,omitempty" yaml:"size,omitempty"`
	SpaceBefore   string   `xml:"spacebefore,attr,omitempty" json:"spacebefore,omitempty" yaml:"spacebefore,omitempty"`
	StyleTab      string   `xml:"style_tab,attr,omitempty" json:"style_tab,omitempty" yaml:"style_tab,omitempty"`
	KeepWithNext  string   `xml:"keepwithnext,attr,omitempty" json:"keepwithnext,omitempty" yaml:"keepwithnext,omitempty"`
	Effects       string   `xml:"effects,attr,omitempty" json:"effects,omitempty" yaml:"effects,omitempty"`
	LeftIdent     string   `xml:"leftindent,attr,omitempty" json:"leftindent,omitempty" yaml:"leftindent,omitempty"`
	RightIdent    string   `xml:"rightindent,attr,omitempty" json:"rightindent,omitempty" yaml:"rightindent,omitempty"`
	Align         string   `xml:"align,attr,omitempty" json:"align,omitempty" yaml:"align,omitempty"`
}

type Paragraphs struct {
	XMLName xml.Name `xml:"paragraphs" json:"paragraphs" yaml:"paragraphs"`
	Para    []*Para  `xml:"para,omitempty" json:"para,omitempty" yaml:"para,omitempty"`
}

type Para struct {
	XMLName    xml.Name `xml:"para" json:"-" yaml:"-"`
	PageNumber string   `xml:"page_number,attr,omitempty" json:"page_number,omitempty" yaml:"page_number,omitempty"`
	Bookmark   string   `xml:"bookmark,attr,omitempty" json:"bookmark,omitempty" yaml:"bookmark,omitempty"`
	Style      *Style   `xml:"style,omitempty" json:"style,omitempty" yaml:"style,omitempty"`
	Text       []*Text  `xml:"text,omitempty" json:"text,omitempty" yaml:"text,omitempty"`
	Marks      *Marks   `xml:"marks,omitempty" json:"marks,omitempty" yaml:"marks,omitempty"`
}

type Text struct {
	XMLName       xml.Name `xml:"text" json:"-" yaml:"-"`
	Underline     string   `xml:"underline,attr,omitempty" json:"underline,omitempty" yaml:"underline,omitempty"`
	Italic        string   `xml:"italic,attr,omitempty" json:"italic,omitempty" yaml:"italic,omitempty"`
	Bold          string   `xml:"bold,attr,omitempty" json:"bold,omitempty" yaml:"bold,omitempty"`
	Strikethrough string   `xml:"strikethrough,attr,omitempty" json:"strikethrough,omitempty" yaml:"strikethrough,omitempty"`
	AllCaps       string   `xml:"allcaps,attr,omitempty" json:"allcaps,omitempty" yaml:"allcaps,omitempty"`
	InnerText     string   `xml:",chardata" json:"inner_text" yaml:"inner_text"`
}

type Marks struct {
	XMLName xml.Name `xml:"marks" json:"marks" yaml:"marks"`
	Mark    []*Mark  `xml:"mark,omitempty" json:"mark,omitempty" yaml:"mark,omitempty"`
}

type Mark struct {
	XMLName  xml.Name `xml:"mark" json:"-" yaml:"-"`
	At       string   `xml:"at,attr,omitempty" json:"at,omitempty" yaml:"at,omitempty"`
	Revision string   `xml:"revision,attr,omitempty" json:"revision,omitempty" yaml:"revision,omitempty"`
}

type Spelling struct {
	XMLName        xml.Name        `xml:"spelling" json:"-" yaml:"-"`
	Language       string          `xml:"language,attr,omitempty" json:"language,omitempty" yaml:"language,omitempty"`
	UserDictionary *UserDictionary `xml:"user_dictionary,omitempty" json:"user_dictionary,omitempty" yaml:"user_dictionary,omitempty"`
}

type UserDictionary struct {
	XMLName xml.Name `xml:"user_dictionary" json:"-" yaml:"-"`
	Entry   []*Entry `xml:"entry,omitempty" json:"entry,omitempty" yaml:"entry,omitempty"`
}

type Entry struct {
	XMLName xml.Name `xml:"entry" json:"-" yaml:"-"`
	Word    string   `xml:"work,attr,omitempty" json:"word,omitempty" yaml:"word,omitempty"`
}

type Lists struct {
	XMLName        xml.Name        `xml:"lists" json:"lists" yaml:"lists"`
	Characters     *Characters     `xml:"characters,omitempty" json:"characters,omitempty" yaml:"characters,omitempty"`
	Locations      *Locations      `xml:"locations,omitempty" json:"locations,omitempty" yaml:"locations,omitempty"`
	SceneIntros    *SceneIntros    `xml:"scene_intros,omitempty" json:"scene_intros,omitempty" yaml:"scene_intros,omitempty"`
	SceneTimes     *SceneTimes     `xml:"scene_times,omitempty" json:"scene_times,omitempty" yaml:"scene_times,omitempty"`
	Extensions     *Extensions     `xml:"extensions,omitempty" json:"extensions,omitempty" yaml:"extensions,omitempty"`
	Transitions    *Transitions    `xml:"transitions,omitempty" json:"transitions,omitempty" yaml:"transitions,omitempty"`
	RevisionColors *RevisionColors `xml:"revision_colors,omitempty" json:"revision_colors,omitempty" yaml:"revision_colors,omitempty"`
	TagCategories  *TagCategories  `xml:"tag_categories,omitempty" json:"tag_categories,omitempty" yaml:"tag_categories,omitempty"`
}

type Characters struct {
	XMLName   xml.Name     `xml:"characters" json:"characters" yaml:"characters"`
	Character []*Character `xml:"character,omitempty" json:"character,omitempty" yaml:"character,omitempty"`
}

type Character struct {
	XMLName xml.Name `xml:"character" json:"-" yaml:"-"`
	Name    string   `xml:"name,attr,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

type Locations struct {
	XMLName  xml.Name    `xml:"locations" json:"locations" yaml:"locations"`
	Location []*Location `xml:"location,omitempty" json:"location,omitempty" yaml:"location,omitempty"`
}

type Location struct {
	XMLName xml.Name `xml:"location" json:"-" yaml:"-"`
	Name    string   `xml:"name,attr,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

type SceneIntros struct {
	XMLName    xml.Name      `xml:"scene_intros" json:"scene_intros" yaml:"scene_intros,omitempty"`
	SceneIntro []*SceneIntro `xml:"scene_intro,omitempty" json:"scene_intro,omitempty" yaml:"screne_intro"`
}

type SceneIntro struct {
	XMLName xml.Name `xml:"scene_intro" json:"-" yaml:"-"`
	Name    string   `xml:"name,attr,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

type SceneTimes struct {
	XMLName   xml.Name     `xml:"scene_times" json:"scene_times" yaml:"scene_times"`
	SceneTime []*SceneTime `xml:"scene_time,omitempty" json:"scene_time,omitempty" yaml:"scene_time,omitempty"`
}

type SceneTime struct {
	XMLName xml.Name `xml:"scene_time" json:"-" yaml:"-"`
	Name    string   `xml:"name,attr,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

type Extensions struct {
	XMLName   xml.Name     `xml:"extensions" json:"extentions" yaml:"extentions"`
	Extension []*Extension `xml:"extension,omitempty" json:"extension,omitempty" yaml:"extension,omitempty"`
}

type Extension struct {
	XMLName xml.Name `xml:"extension" json:"-" yaml:"-"`
	Name    string   `xml:"name,attr,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

type Transitions struct {
	XMLName    xml.Name      `xml:"transitions" json:"transitions" yaml:"transitions"`
	Transition []*Transition `xml:"transition,omitempty" json:"transition,omitempty" yaml:"transition,omitempty"`
}

type Transition struct {
	XMLName xml.Name `xml:"transition" json:"-" yaml:"-"`
	Name    string   `xml:"name,attr,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

type RevisionColors struct {
	XMLName       xml.Name         `xml:"revision_colors" json:"revision_colors" yaml:"revision_colors,omitempty"`
	RevisionColor []*RevisionColor `xml:"revision_color,omitempty" json:"revision_color,omitempty" yaml:"revision_color,omitempty"`
}

type RevisionColor struct {
	XMLName    xml.Name `xml:"revision_color" json:"-" yaml:"-"`
	Name       string   `xml:"name,attr,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Index      string   `xml:"index,attr,omitempty" json:"index,omitempty" yaml:"index,omitempty"`
	ColorName  string   `xml:"color_name,attr,omitempty" json:"color_name,omitempty" yaml:"color_name,omitempty"`
	ColorIndex string   `xml:"color_index,attr,omitempty" json:"color_index,omitempty" yaml:"color_index,omitempty"`
}

type TagCategories struct {
	XMLName     xml.Name       `xml:"tag_categories" json:"tag_categories" yaml:"tag_categories"`
	TagCategory []*TagCategory `xml:"tag_category,omitempty" json:"tag_category,omitempty" yaml:"tag_category,omitempty"`
}

type TagCategory struct {
	XMLName xml.Name `xml:"tag_category" json:"-" yaml:"-"`
	Name    string   `xml:"name,attr,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

type TitlePage struct {
	XMLName xml.Name `xml:"titlepage" json:"titlepage" yaml:"titlepage"`
	Para    []*Para  `xml:"para,omitempty" json:"para,omitempty" yaml:"para,omitempty"`
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
					s = "~" + s
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

// NewOpenScreenplay20 creates a new OpenScreenplay document set to version 2.0
func NewOpenScreenplay20() *OpenScreenplay {
	doc := new(OpenScreenplay)
	doc.Version = "20"
	doc.Type = "Open Screenplay Format document"
	return doc
}

// CleanupSelfClosingElements changes something like <styles></styles> to <styles/>
func CleanupSelfClosingElements(src []byte) []byte {
	for _, elem := range []string{"info", "settings", "styles", "style", "mark", "text", "entry", "character", "location", "scene_time", "extension", "revision_color", "tag_category", "transition", "spelling", "user_dictionary", "paragraphs", "para", "locations"} {
		src = bytes.Replace(src, []byte("></"+elem+">"), []byte("/>"), -1)
	}
	for _, elem := range []string{"titlepage"} {
		src = bytes.Replace(src, []byte("<"+elem+"></"+elem+">"), []byte("<"+elem+"/>"), -1)
	}

	return src
}

// ToXML takes a OpenScreenplay struct and renders XML
func (document *OpenScreenplay) ToXML() ([]byte, error) {
	src, err := xml.MarshalIndent(document, "", "    ")
	if err != nil {
		return nil, err
	}
	return CleanupSelfClosingElements(src), nil
}
