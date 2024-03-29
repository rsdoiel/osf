#
# Makefile for running pandoc on all Markdown docs ending in .md
#
PROJECT = osf

MD_PAGES = $(shell ls -1 *.md)

HTML_PAGES = $(shell ls -1 *.md | sed -E 's/.md/.html/g')

MAN_PAGES = osf2txt txt2osf fadein2osf

MAN_HTML = osf2txt.html txt2osf.html fadein2osf.html

build: $(HTML_PAGES) $(MD_PAGES) $(MAN_HTML) LICENSE.html

$(HTML_PAGES): $(MD_PAGES) .FORCE
	pandoc --metadata title=$(basename $@) -s --to html5 $(basename $@).md -o $(basename $@).html \
	    --template=page.tmpl
	@if [ $@ = "README.html" ]; then mv README.html index.html; fi

$(MAN_HTML): .FORCE
	pandoc docs/$(basename $@).md -s --to html5 -o $@ \
		--template=page.tmpl

LICENSE.html: LICENSE
	pandoc --metadata title="$(PROJECT) License" -s --from Markdown --to html5 LICENSE -o license.html \
	    --template=page.tmpl

clean:
	@if [ -f index.html ]; then rm *.html; fi
	#@if [ -f docs/index.html ]; then rm docs/*.html; fi

.FORCE:
