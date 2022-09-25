PROJECT := ignoregit
BUILDROOT := build
EXE := ignore
EXEPATH := $(BUILDROOT)/$(EXE)

$(EXEPATH): $(wildcard %.go) $(BUILDROOT)
	go build -o $(EXEPATH)

$(BUILDROOT):
	mkdir -p build

.PHONY: test
test:

.PHONY: clean
clean:
	rm -rf $(BUILDROOT)

.PHONY: fmt
fmt: $(wildcard *.go)
	gofmt -s -w -e $<

.PHONY: install
install: $(EXEPATH)
	strip $(EXEPATH)
	install -D -m 644 presets.txt $(DESTDIR)/etc/ignoregit/presets.txt
	install -D -m 644 LICENSE $(DESTDIR)/etc/ignoregit/LICENSE
	install -D -m 511 $(EXEPATH) $(DESTDIR)/usr/bin/$(EXE)
