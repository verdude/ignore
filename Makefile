PROJECT := ignoregit
BUILDROOT := build
EXE := $(BUILDROOT)/ignore

$(EXE): $(wildcard %.go) $(BUILDROOT)
	go build -o $(EXE)

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
install: $(EXE)
	strip $(EXE)
	install -D -m 644 presets.txt $(DESTDIR)/etc/ignoregit/presets.txt
	install -D -m 644 LICENSE $(DESTDIR)/etc/ignoregit/LICENSE
	install -D -m 511 $(EXE) $(DESTDIR)/usr/bin/$(EXE)
