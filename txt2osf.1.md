%txt2osf(1) user manual | version 0.0.10 722e57f
% R. S. Doiel
% 2026-11-18

# NAME

txt2osf — convert a Fountain or plain text screenplay to Open Screenplay Format

# SYNOPSIS

txt2osf [OPTIONS] [INPUT_FILENAME [OUTPUT_FILENAME]]

# DESCRIPTION

txt2osf reads a plain text screenplay in Fountain markup and writes an Open Screenplay Format (OSF) 2.0 document.

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
: add a trailing newline (the default is to add one; -nl=false leaves it out)

-quiet
: suppress error messages; the exit status still says what happened

The input and output may also be given as the first and second argument.

# EXIT STATUS

0
: success

2
: bad command line: an unknown option, a surplus argument

65
: the input is not valid Fountain

66
: the input file does not exist

73
: the output file cannot be created

74
: reading or writing failed part way

77
: the operating system refused access to a file

# EXAMPLES

    txt2osf -i screenplay.txt -o screenplay.osf

Or alternatively

    cat screenplay.txt | txt2osf > screenplay.osf

