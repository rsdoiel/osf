{
    "title": "OSF, A Go package support Open Screenplay Format"
}

# About the test data

The [fountain](fountain.io) website has some good files for reviewing
formatting difference between fountain, fdx and PDF.  Some are referenced in the 
test programs but they are optional.  It is not clear to me the licensing 
arrangements for the text so I have not included them in this repository.

If you want to include them in the test sequence go to the fountain website and
download them and place them in the _testdata_ directory. When you
run `go test` they will be found and included in the basic test process.

## Optional test FDX files

+ [Big Fish](https://fountain.io/_downloads/Big-Fish.fountain)
    + [fdx](https://fountain.io/_downloads/Big-Fish.fdx)
    + [pdf](https://fountain.io/_downloads/Big-Fish.pdf) 
+ [Brick & Stell](https://fountain.io/_downloads/Brick-&-Steel.fountain)
    + [fdx](https://fountain.io/_downloads/Brick-&-Steel.fdx)
    + [pdf](https://fountain.io/_downloads/Brick-&-Steel.pdf)
+ [Birthday Card](https://fountain.io/_downloads/The-Last-Birthday-Card.fountain)
    + [fdx](https://fountain.io/_downloads/The-Last-Birthday-Card.fdx)
    + [pdf](https://fountain.io/_downloads/The-Last-Birthday-Card.pdf)

