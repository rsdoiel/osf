[![Project Status: Active – The project has reached a stable, usable state and is being actively developed.](https://www.repostatus.org/badges/latest/active.svg)](https://www.repostatus.org/#active)

# osf

_osf_ is a Go package and three small command line programs for working with Open Screenplay Format (OSF) 2.0. OSF is an open XML format for screenplays and, zipped, the native format of [Fade In](https://www.fadeinpro.com). The package parses an OSF document, writes it back as XML, renders it as Fountain-like plain text, and builds one from a [Fountain](https://fountain.io) document parsed by the [fountain](https://github.com/rsdoiel/fountain) package.

OSF keeps no number for an individual scene (numbering is a document setting, `scene_numbering`), so a Fountain scene number is not carried into OSF. A Fountain boneyard (`/* ... */`) becomes an ordinary paragraph, so it shows in Fade In.

## The programs

[osf2txt](osf2txt.1.md)
: reads an OSF file and writes plain text in a Fountain-like format.

[txt2osf](txt2osf.1.md)
: reads a plain text file and renders an OSF 2.0 document.

[fadein2osf](fadein2osf.1.md)
: reads a Fade In (`.fadein`) file and writes Open Screenplay Format.

## Install

Get a release from [GitHub](https://github.com/rsdoiel/osf/releases), or see [INSTALL.md](INSTALL.md), which also covers installing from source. To use the package in a Go program:

```shell
go get github.com/rsdoiel/osf
```

## Build and test

You need Go 1.26.4 or later. Pandoc and GNU Make are only needed for the man pages and the website, and `cmt` (CMTools) only to regenerate `version.go`, `about.md` and `CITATION.cff` from `codemeta.json`.

```shell
make          # build the programs into bin/ and the man pages into man/
make test     # go test ./...
make check    # gofmt -l and go vet; rewrites nothing
```
