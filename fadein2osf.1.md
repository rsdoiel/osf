
USAGE: fadein2osf [OPTIONS]

DESCRIPTION

fadein2osf is a command line program that reads an ".fadein" file
and write outs a OSF 2.0 XML.

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

Convert *screenplay.fadein* into *screenplay.osf*.

    fadein2osf -i screenplay.fadein -o screenplay.osf

Display converted OSF 2.0 XML to the console

	fadein2osf -i screenplay.fadein

fadein2osf 0.0.6
