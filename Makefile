.PHONY: build test
all: build

DIR     := $(PWD)

clean:
	rm -rf $(PWD)/bin
	mkdir $(PWD)/bin

test:
	pushd test; $(PWD)/test/run.sh; popd

build:
	go build -o bin/reptile