package osf

import "strings"

/** plainText returns s as a single unformatted text run, with no
 * emphasis parsing.
 *
 * Parameters:
 *   s (string) — the text
 *
 * Returns:
 *   []*Text — one run holding s unchanged
 *
 * Example:
 *   runs := plainText("a *boneyard* note")
 */
func plainText(s string) []*Text {
	return []*Text{{InnerText: s}}
}

/** StringToTextArray turns a Fountain string into OSF text runs. Inline
 * emphasis becomes attributes on a run: ***x*** is bold italic, **x** is
 * bold, *x* is italic and _x_ is underline; they nest, and a backslash
 * escapes a marker (\*, \_ and \\). A marker that is unmatched, or that
 * sits next to a space (as in "5 * 3"), or sits inside a word (snake_case),
 * stays literal text.
 *
 * Parameters:
 *   s (string) — Fountain text for one element
 *
 * Returns:
 *   []*Text — the runs in order, never empty (an empty s gives one empty run)
 *
 * Example:
 *   runs := StringToTextArray("He *runs* far") // "He ", italic "runs", " far"
 */
func StringToTextArray(s string) []*Text {
	disabled := map[int]bool{}
	for {
		texts, unclosed := scanInline(s, disabled)
		if len(unclosed) == 0 {
			if len(texts) == 0 {
				return plainText("")
			}
			return texts
		}
		for _, pos := range unclosed {
			disabled[pos] = true
		}
	}
}

func isSpace(b byte) bool {
	return b == ' ' || b == '\t' || b == '\n' || b == '\r'
}

func isWordByte(b byte) bool {
	return b >= '0' && b <= '9' || b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' || b >= 0x80
}

// scanInline makes one pass over s. Markers at the positions in disabled
// are literal. It returns the runs and the positions of openers that were
// never closed, which the caller disables before scanning again.
func scanInline(s string, disabled map[int]bool) ([]*Text, []int) {
	var (
		texts            []*Text
		buf              strings.Builder
		bold, ital, undr bool
		openBold         = -1
		openItal         = -1
		openUndr         = -1
	)
	flush := func() {
		if buf.Len() == 0 {
			return
		}
		t := &Text{InnerText: buf.String()}
		if bold {
			t.Bold = BoldStyle
		}
		if ital {
			t.Italic = ItalicStyle
		}
		if undr {
			t.Underline = UnderlineStyle
		}
		texts = append(texts, t)
		buf.Reset()
	}
	// canOpen: the marker is followed by a non-space; canClose: it follows one.
	canOpen := func(i, n int) bool {
		return i+n < len(s) && !isSpace(s[i+n]) && s[i+n] != s[i]
	}
	canClose := func(i int) bool { return i > 0 && !isSpace(s[i-1]) }
	for i := 0; i < len(s); {
		c := s[i]
		if c == '\\' && i+1 < len(s) && (s[i+1] == '*' || s[i+1] == '_' || s[i+1] == '\\') {
			buf.WriteByte(s[i+1])
			i += 2
			continue
		}
		if (c != '*' && c != '_') || disabled[i] {
			buf.WriteByte(c)
			i++
			continue
		}
		n := 0
		switch {
		case c == '_':
			wordStart := i == 0 || !isWordByte(s[i-1])
			switch {
			case undr && canClose(i):
				flush()
				undr, openUndr = false, -1
				n = 1
			case !undr && wordStart && canOpen(i, 1):
				flush()
				undr, openUndr = true, i
				n = 1
			}
		default:
			run := 0
			for i+run < len(s) && s[i+run] == '*' {
				run++
			}
			if run > 3 {
				run = 3
			}
			for m := run; m >= 1 && n == 0; m-- {
				var wantBold, wantItal bool
				switch m {
				case 3:
					wantBold, wantItal = true, true
				case 2:
					wantBold = true
				default:
					wantItal = true
				}
				closing := (!wantBold || bold) && (!wantItal || ital)
				opening := (!wantBold || !bold) && (!wantItal || !ital)
				switch {
				case closing && canClose(i):
					flush()
					if wantBold {
						bold, openBold = false, -1
					}
					if wantItal {
						ital, openItal = false, -1
					}
					n = m
				case opening && canOpen(i, m):
					flush()
					if wantBold {
						bold, openBold = true, i
					}
					if wantItal {
						ital, openItal = true, i
					}
					n = m
				}
			}
		}
		if n == 0 {
			buf.WriteByte(c)
			i++
			continue
		}
		i += n
	}
	flush()
	var unclosed []int
	for _, pos := range []int{openBold, openItal, openUndr} {
		if pos >= 0 {
			unclosed = append(unclosed, pos)
		}
	}
	return texts, unclosed
}
