
# Action Items

## Bugs

## Next

+ [ ] Write osf.go, osf_test.go based on [Open Screenplay Format 2.0](https://sourceforge.net/projects/openscrfmt/) and in the mode of [fdx](https://github.com/rsdoiel/fdx) package
    + [ ] Convert testdata/sample-0?.fdx to testdata/sample-0?.osf
+ [ ] Write osf2txt
+ [ ] Write fadein2txt
+ [ ] Add ParseFile() to osf.go, if file extension is ".fadein" then it should handle the unzipping and and parsing of document.xml as OSF

## Someday, Maybe

+ [ ] Support parsing .fadein files (i.e. unzip the Fade In file, then parse document.xml)

## Completed

### Reference links

+ [Fountain](https://fountain.io)
+ [Open Screenplay Format 2.0](https://sourceforge.net/projects/openscrfmt/) (the one targetted by osf.go)
+ [Open Screenplay Format 2.1](https://github.com/severdia/Open-Screenplay-Format)
+ [Fade In](https://www.fadeinpro.com)
+ [Open Screenplay Format by Kent Tessman](http://www.kenttessman.com/2012/02/open-screenplay-format/)

