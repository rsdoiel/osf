[![Project Status: Inactive – The project has reached a stable, usable state but is no longer being actively developed; support/maintenance will be provided as time allows.](https://www.repostatus.org/badges/latest/inactive.svg)](https://www.repostatus.org/#inactive)


OSF
===

A Go package support Open Screenplay Format
-------------------------------------------

Experimental Golang package for working with Open Screenplay Format 2.0.
Open Screenplay Format is an open XML format for screenplays and the
native format (when zipped) for [Fade In](https://www.fadeinpro.com).
Two package will include several demonstration command line programs 
[osf2txt](docs/osf2txt.html) which will read a osf file and render plain 
text in a [Fountain](https://fountain.io) like format, [txt2osf](docs/txt2osf.html) 
which takes a plain text file and attempts to render an OSF 2.0 document 
and finally [fadein2osf](docs/fadein2osf) which will read in a Fade In file 
and write out Open Screenplay Format.

