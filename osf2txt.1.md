
USAGE: osf2txt [OPTIONS]

DESCRIPTION

osf2txt is a command line program that reads an osf file
and returns plain text

OPTIONS

    -generate-manpage   generate man page
    -generate-markdown  generate Markdown documentation
    -h, -help           display help
    -i, -input          set the input filename
    -l, -license        display license
    -nl, -newline       add a trailing newline
    -o, -output         set the output filename
    -quiet              suppress error messages
    -v, -version        display version


EXAMPLES

Convert *screenplay.osf* into *screenplay.txt*.

    osf2txt -i screenplay.osf -o screenplay.txt

Or alternatively

    cat screenplay.osf | osf2txt > screenplay.txt

osf2txt 0.0.4
