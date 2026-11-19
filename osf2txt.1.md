%osf2txt(1) user manual | version 0.0.10 722e57f
% R. S. Doiel
% 2026-11-18

# NAME

osf2txt — convert an Open Screenplay Format document to plain text

# SYNOPSIS

osf2txt [OPTIONS] [INPUT_FILENAME [OUTPUT_FILENAME]]

# DESCRIPTION

osf2txt reads an Open Screenplay Format (OSF) document and writes the screenplay as plain text.

# OPTIONS

-h, -help
: display this help page

-l, -license
: display the license

-v, -version
: display the version

-i, -input FILENAME
: read from FILENAME; the default, and -, is standard input

-o, -output FILENAME
: write to FILENAME; the default, and -, is standard output

-nl, -newline
: add a trailing newline (the default is none)

-quiet
: suppress error messages; the exit status still says what happened

The input and output may also be given as the first and second argument.

# EXIT STATUS

0
: success

2
: bad command line: an unknown option, a surplus argument

65
: the input is not valid OSF XML

66
: the input file does not exist

73
: the output file cannot be created

74
: reading or writing failed part way

77
: the operating system refused access to a file

# EXAMPLES

    osf2txt -i screenplay.osf -o screenplay.txt

Or alternatively

    cat screenplay.osf | osf2txt > screenplay.txt

