
USAGE: txt2osf [OPTIONS]

DESCRIPTION

txt2osf is a command line program that reads an plain text file
and returns an OSF 2.0 text

OPTIONS

    -generate-manpage    generate man page
    -generate-markdown   generate Markdown documentation
    -h, -help            display help
    -i, -input           set the input filename
    -l, -license         display license
    -nl, -newline        add a trailing newline
    -o, -output          set the output filename
    -quiet               suppress error messages
    -v, -version         display version


EXAMPLES

Convert *screenplay.txt* into *screenplay.osf*.

    txt2osf -i screenplay.txt -o screenplay.osf

Or alternatively

    cat screenplay.txt | txt2osf > screenplay.osf

txt2osf 0.0.8
