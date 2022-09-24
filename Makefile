PROJECT := ignoregit
BUILDROOT := build
EXE := ignore

$(EXE): $(wildcard %.go) $(BUILDROOT)
	go build -o $(BUILDROOT)/$(EXE)

$(BUILDROOT):
	mkdir -p build

.PHONY: test
test:

.PHONY: clean
clean:
	rm -rf $(BUILDROOT)

.PHONY: install
install: $(EXE)
	strip $(BUILDROOT)/$(EXE)
	install -D -m 444 presets.txt $(DESTDIR)/etc/ignoregit/presets.txt
	install -D -m 444 LICENSE $(DESTDIR)/etc/ignoregit/LICENSE
	install -D -m 111 $(BUILDROOT)/$(EXE) $(DESTDIR)/usr/bin/$(EXE)
