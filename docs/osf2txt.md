% osf2txt(1) osf2txt user manual
% R. S. Doiel
% August 7, 2022

# NAME

osf2txt

# SYNOPSIS

osf2txt [OPTIONS]

# DESCRIPTION

osf2txt is a command line program that reads an osf file
and returns plain text


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

Cervert *screenplay.osf* into *screenplay.txt*.

~~~shell
    osf2txt -i screenplay.osf -o screenplay.txt
~~~

Or alternatively

~~~shell
    cat screenplay.osf | osf2txt > screenplay.txt
~~~



