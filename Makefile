GOEXE = go
GOFMTEXE = gofmt
GOFMTARGS = -s -w -l
GOLINTEXE = staticcheck
GOLINTARGS =

TESTDIR = .save

all: build

build:
	$(GOEXE) mod download
	$(GOEXE) build

install:
	$(GOEXE) install .

fmt:
	$(GOFMTEXE) $(GOFMTARGS) *.go

check: lint
vet: lint
lint:
	$(GOLINTEXE) $(GOLINTARGS) *.go

clean: clean-bin clean-test

clean-bin:
	rm -f save save.exe

clean-test:
	rm -rf $(TESTDIR)
