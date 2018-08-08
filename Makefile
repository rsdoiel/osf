#
# Simple Makefile
#
PROJECT = osf

VERSION = $(shell grep -m1 'Version = ' $(PROJECT).go | cut -d\"  -f 2)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

OS = $(shell uname)

EXT = 
ifeq ($(OS), Windows)
	EXT = .exe
endif

build: bin/osf2txt$(EXT) bin/fadein2osf$(EXT) bin/txt2osf$(EXT)

bin/osf2txt$(EXT): osf.go cmd/osf2txt/osf2txt.go 
	go build -o bin/osf2txt$(EXT) cmd/osf2txt/osf2txt.go

bin/fadein2osf$(EXT): osf.go cmd/fadein2osf/fadein2osf.go
	go build -o bin/fadein2osf$(EXT) cmd/fadein2osf/fadein2osf.go

bin/txt2osf$(EXT): osf.go cmd/txt2osf/txt2osf.go
	go build -o bin/txt2osf$(EXT) cmd/txt2osf/txt2osf.go

test:
	go test

man: build
	mkdir -p man/man1
	bin/osf2txt -generate-manpage | nroff -Tutf8 -man > man/man1/osf2txt.1
	bin/fadein2osf -generate-manpage | nroff -Tutf8 -man > man/man1/fadein2osf.1
	bin/txt2osf -generate-manpage | nroff -Tutf8 -man > man/man1/txt2osf.1

clean: 
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi
	if [ -d man ]; then rm -fR man; fi

website:
	./mk-website.bash

status:
	git status

save:
	git commit -am "Quick Save"
	git push origin $(BRANCH)

publish:
	./mk-website.bash
	./publish.bash

