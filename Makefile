PREFIX ?= /usr/local
BINDIR ?= $(PREFIX)/bin

# GO ?= $(shell which go)
GO ?= go
GOFLAGS ?=

GOSRC := $(shell find . -name '*.go')
GOSRC += go.mod go.sum

IMGS := $(shell find . -name '*.png')

RM ?= rm -f

default: shape-of-code

shape-of-code: $(GOSRC)
	$(GO) build $(GOFLAGS) -o $@

install: default
	install -D -m755 shape-of-code $(DESTDIR)$(BINDIR)/shape-of-code

uninstall:
	$(RM) $(DESTDIR)$(BINDIR)/shape-of-code

clean:
	$(RM) shape-of-code

clean-imgs: $(IMGS)
	$(RM) $(IMGS)

.PHONY: default shape-of-you install uninstall clean clean-imgs
