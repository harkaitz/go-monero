PROJECT=go-monero
VERSION=1.0.2b
PREFIX=/usr/local

## -- BLOCK:go --
all: all-go
install: install-go
clean: clean-go
deps: deps-go

build/monero-cli$(EXE): deps
	go build -o $@ $(GO_CONF) ./cmd/monero-cli

all-go:  build/monero-cli$(EXE)
deps-go:
	mkdir -p build
install-go:
	install -d $(DESTDIR)$(PREFIX)/bin
	cp  build/monero-cli$(EXE) $(DESTDIR)$(PREFIX)/bin
clean-go:
	rm -f  build/monero-cli$(EXE)
## -- BLOCK:go --
## -- BLOCK:license --
install: install-license
install-license: 
	mkdir -p $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
	cp  README.md $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
update: update-license
update-license:
	ssnip README.md
## -- BLOCK:license --
