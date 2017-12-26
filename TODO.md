
# Action Items

## Bugs

+ [ ] String (Fountain style plain text) needs to be formatted correctly...

## Next

+ [ ] Write osf.go, osf_test.go based on [Open Screenplay Format 2.0](https://sourceforge.net/projects/openscrfmt/) and in the mode of [fdx](https://github.com/rsdoiel/fdx) package
    + [ ] Convert testdata/sample-0?.fdx to testdata/sample-0?.osf
+ [ ] Write osf2txt

## Someday, Maybe

## Completed

+ [x] Write fadein2txt
+ [x] self closing tags should be self closing
+ [x] Support parsing .fadein files (i.e. unzip the Fade In file, then parse document.xml)
+ [x] Add ParseFile() to osf.go, if file extension is ".fadein" then it should handle the unzipping and and parsing of document.xml as OSF

### Reference links

+ [Fountain](https://fountain.io)
+ [Open Screenplay Format 2.0](https://sourceforge.net/projects/openscrfmt/) (the one targetted by osf.go)
+ [Open Screenplay Format 2.1](https://github.com/severdia/Open-Screenplay-Format)
+ [Fade In](https://www.fadeinpro.com)
+ [Open Screenplay Format by Kent Tessman](http://www.kenttessman.com/2012/02/open-screenplay-format/)
+ [screenplay-parser](https://github.com/azcoppen/screenplay-parser) - a PHP repo with a really nice README.md discussing format and convertion issues and challenges

