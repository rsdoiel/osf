---
title: osf
abstract: |-
  A library and command line programs for working with FadeIn and Open Screenplay Format (osf) files.
authors:
  - family_name: Doiel
    given_name: R. S.
    id: https://orcid.org/0000-0003-0900-6903


maintainer:
  - family_name: Doiel
    given_name: R. S.
    id: https://orcid.org/0000-0003-0900-6903

repository_code: https://github.com/rsdoiel/osf
version: 0.0.10
license_url: https://www.gnu.org/licenses/agpl-3.0.txt

programming_language:
  - Go >= 1.27.3

keywords:
  - GitHub
  - text markup
  - screen play

date_released: 2026-11-18
---

About this software
===================

## osf 0.0.10

Modernization to the workspace standard. Go 1.27.3 is the minimum version. The three programs (fadein2osf, osf2txt, txt2osf) use the standard library flag package instead of github.com/caltechlibrary/cli, with -h, -l and -v beside the long spellings. Requires fountain v1.0.4. Fixed: txt2osf (FromFountain) wrote each paragraph's kind as a bookmark attribute, so Fade In could not tell a scene heading from action; it now writes <style basestylename="..."/> as Fade In's own files do, and no longer writes an empty paragraph for each Fountain blank line. Also fixed in txt2osf: Fountain forcing markers (!, @, ~, ., >) are no longer left in the text, title page values are trimmed, and inline emphasis (*italic*, **bold**, ***both***, _underline_) becomes formatted text runs instead of literal asterisks. osf2txt applies a paragraph's style once, not once per text run. Test coverage of the library rose from 28% to 97%. The 32-bit Raspberry Pi (armv7l) target is dropped. The release scripts follow the workspace standard (a release is built from a clean, pushed commit, and an interrupted draft release resumes).

## Authors

- [R. S. Doiel](https://orcid.org/0000-0003-0900-6903)




## Maintainers

- [R. S. Doiel](https://orcid.org/0000-0003-0900-6903)


A library and command line programs for working with FadeIn and Open Screenplay Format (osf) files.

- [License](https://www.gnu.org/licenses/agpl-3.0.txt)
- [Code Repository](https://github.com/rsdoiel/osf)
  - [Issue Tracker](https://github.com/rsdoiel/osf/issues)

## Programming languages

- Go >= 1.27.3




## Software Requirements

- Go >= 1.27.3
- github.com/rsdoiel/fountain >= 1.0.4
- gopkg.in/yaml.v3 >= 3.0.1


## Software Suggestions

- CMTools >= 0.0.45b
- Pandoc >= 3.9
- GNU Make >= 3


