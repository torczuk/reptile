.PHONY: build test
all: build

DIR     := $(PWD)

test:
	pushd server; make test; popd
	pushd system_test; $(PWD)/system_test/run.sh;

build:
	pushd server; make build; popd