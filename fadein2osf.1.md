%fadein2osf(1) user manual | version 0.0.9 9812fdf
% R. S. Doiel
% 2026-10-20

# NAME

fadein2osf — convert a Fade In file to Open Screenplay Format

# SYNOPSIS

fadein2osf [OPTIONS] INPUT_FILENAME [OUTPUT_FILENAME]

# DESCRIPTION

fadein2osf reads a Fade In (.fadein) file and writes an Open Screenplay Format (OSF) 2.0 document. A Fade In file is a zip archive, so the input must be a named file and cannot be read from standard input.

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
: bad command line: an unknown option, a surplus argument, a missing input file name

65
: the input is not valid Fade In archive

66
: the input file does not exist

73
: the output file cannot be created

74
: reading or writing failed part way

77
: the operating system refused access to a file

# EXAMPLES

    fadein2osf -i screenplay.fadein -o screenplay.osf

Display the converted OSF 2.0 XML on the console

    fadein2osf -i screenplay.fadein

