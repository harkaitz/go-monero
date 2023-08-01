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
