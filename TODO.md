{
    "title": "OSF, A Go package support Open Screenplay Format"
}

# Action Items

## Bugs


## Next

- [x] Update code to Go 1.27.3
- [x] Remove dependency on `github.com/caltechlibrary/cli`
- [ ] implement a txt2osf demonstration
- [x] review Text element: inline emphasis (`***`, `**`, `*`, `_`, `\` escapes) becomes text runs; newlines kept; Para.String applies the style once per paragraph (2026-11-18)
- [ ] validate ToXML() after FromFountain() can be read by FadeIn

## Someday, Maybe

- [ ] write and osf2html using [scrippets](https://fountain.io/scrippets) approach
- [ ] add support for Ron Severdia's Open Screenplay Format 2.1 spec
    - rename divergent structs' xml defs with 20 and 21 suffix
    - make sure they are all 20/21 structs are treated as tag ",omitempty" 
    - duplicate String methods as needed
    - Parse, ParseFile should work without sniffing using a single struct tree for 1.2, 2.0 or 2.1 

## Completed

- [x] String (Fountain style plain text) needs to be formatted correctly...
- [x] Write osf.go, osf_test.go based on [Open Screenplay Format 2.0](https://sourceforge.net/projects/openscrfmt/) and in the mode of [fdx](https://github.com/rsdoiel/fdx) package
- [x] Write osf2txt
- [x] Write fadein2osf
- [x] Write fadein2txt
- [x] self closing tags should be self closing
- [x] Support parsing .fadein files (i.e. unzip the Fade In file, then parse document.xml)
- [x] Add ParseFile() to osf.go, if file extension is ".fadein" then it should handle the unzipping and and parsing of document.xml as OSF

### Reference links

- [Fountain](https://fountain.io)
- [Open Screenplay Format 2.0](https://sourceforge.net/projects/openscrfmt/) (the one targeted by osf.go)
- [Open Screenplay Format 2.1](https://github.com/severdia/Open-Screenplay-Format)
- [Fade In](https://www.fadeinpro.com)
- [Open Screenplay Format by Kent Tessman](http://www.kenttessman.com/2012/02/open-screenplay-format/)
- [screenplay-parser](https://github.com/azcoppen/screenplay-parser) - a PHP repo with a really nice README.md discussing format and conversion issues and challenges


## Modernization (2026-11-16)

To match the workspace standard (knowledge and opml are the models; Go 1.27.3 across repositories). Do one repository at a time; run the tests and `make -n release` after each step.

- [x] Go 1.27.3: change `go 1.26.4` in go.mod, the `Go >=` lines in codemeta.json, README.md and INSTALL.md; run `go mod tidy`, `go vet ./...` and the tests.
- [x] Replace `github.com/caltechlibrary/cli` with the standard library `flag`, with `-h/-l/-v` and long spellings (Go CLI conventions).
- [x] Makefile: add `check-go-version` (go.mod, codemeta.json and the Go in use agree); `check-release` calls it.
- [x] Drop the armv7l target and 32-bit Pi wording (dropped 2026-10-27).
- [x] release.ps1: take knowledge's (existing-draft check, checksums).
- [x] Requires fountain v1.0.4 (2026-11-18).
- [x] Tests: round trip of every sample, FromFountain styles, ParseFile errors (library 95%; String, ToXML, nil receivers, every title-page key).
- [ ] Open a generated .osf in Fade In by hand and record the result (needs a machine with Fade In).
- [x] Forced markers (`!`, `@`, `~`, `.`, `>`/`<`) are stripped in FromFountain (2026-11-18).
- [x] Title page values are trimmed; Info is one line, title-page text keeps one line per line (2026-11-18).
- [ ] fountain classifies `CUT TO:` as Action (isAction runs before isTransition); report upstream.
