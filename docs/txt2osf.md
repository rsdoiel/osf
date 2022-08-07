% txt2osf(1) txt2osf user manual
% R. S. Doiel
% August 7, 2022

# NAME

txt2osf

# SYNOPSIS

txt2osf [OPTIONS]

# DESCRIPTION

txt2osf is a command line program that reads an plain text file
and returns an OSF 2.0 text


# OPTIONS

-h, -help
: display help

-i, -input
: set the input filename

-l, -license
: display license

-nl, -newline
: add a trailing newline

-o, -output
: set the output filename

-quiet
: suppress error messages

-v, -version
: display version


# EXAMPLES

Convert *screenplay.txt* into *screenplay.osf*.

~~~shell
    txt2osf -i screenplay.txt -o screenplay.osf
~~~

Or alternatively

~~~shell
    cat screenplay.txt | txt2osf > screenplay.osf
~~~

