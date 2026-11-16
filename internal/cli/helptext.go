package cli

/** Osf2txtHelpText is the osf2txt(1) page, shown by `osf2txt -help`. The {app_name},
 * {version}, {release_hash} and {release_date} tokens are replaced by osf.FmtHelp.
 * The Makefile writes the page to osf2txt.1.md.
 *
 * Example:
 *   fmt.Println(osf.FmtHelp(cli.Osf2txtHelpText, "osf2txt", osf.Version, osf.ReleaseDate, osf.ReleaseHash))
 */
const Osf2txtHelpText = `%{app_name}(1) user manual | version {version} {release_hash}
% R. S. Doiel
% {release_date}

# NAME

{app_name} — convert an Open Screenplay Format document to plain text

# SYNOPSIS

{app_name} [OPTIONS] [INPUT_FILENAME [OUTPUT_FILENAME]]

# DESCRIPTION

{app_name} reads an Open Screenplay Format (OSF) document and writes the screenplay as plain text.

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

    {app_name} -i screenplay.osf -o screenplay.txt

Or alternatively

    cat screenplay.osf | {app_name} > screenplay.txt
`

/** Txt2osfHelpText is the txt2osf(1) page, shown by `txt2osf -help`. The {app_name},
 * {version}, {release_hash} and {release_date} tokens are replaced by osf.FmtHelp.
 * The Makefile writes the page to txt2osf.1.md.
 *
 * Example:
 *   fmt.Println(osf.FmtHelp(cli.Txt2osfHelpText, "txt2osf", osf.Version, osf.ReleaseDate, osf.ReleaseHash))
 */
const Txt2osfHelpText = `%{app_name}(1) user manual | version {version} {release_hash}
% R. S. Doiel
% {release_date}

# NAME

{app_name} — convert a Fountain or plain text screenplay to Open Screenplay Format

# SYNOPSIS

{app_name} [OPTIONS] [INPUT_FILENAME [OUTPUT_FILENAME]]

# DESCRIPTION

{app_name} reads a plain text screenplay in Fountain markup and writes an Open Screenplay Format (OSF) 2.0 document.

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

    {app_name} -i screenplay.txt -o screenplay.osf

Or alternatively

    cat screenplay.txt | {app_name} > screenplay.osf
`

/** Fadein2osfHelpText is the fadein2osf(1) page, shown by `fadein2osf -help`. The {app_name},
 * {version}, {release_hash} and {release_date} tokens are replaced by osf.FmtHelp.
 * The Makefile writes the page to fadein2osf.1.md.
 *
 * Example:
 *   fmt.Println(osf.FmtHelp(cli.Fadein2osfHelpText, "fadein2osf", osf.Version, osf.ReleaseDate, osf.ReleaseHash))
 */
const Fadein2osfHelpText = `%{app_name}(1) user manual | version {version} {release_hash}
% R. S. Doiel
% {release_date}

# NAME

{app_name} — convert a Fade In file to Open Screenplay Format

# SYNOPSIS

{app_name} [OPTIONS] INPUT_FILENAME [OUTPUT_FILENAME]

# DESCRIPTION

{app_name} reads a Fade In (.fadein) file and writes an Open Screenplay Format (OSF) 2.0 document. A Fade In file is a zip archive, so the input must be a named file and cannot be read from standard input.

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

    {app_name} -i screenplay.fadein -o screenplay.osf

Display the converted OSF 2.0 XML on the console

    {app_name} -i screenplay.fadein
`
