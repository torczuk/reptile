.PHONY: build test
all: build

DIR     := $(PWD)

clean:
	rm -rf $(PWD)/bin
	mkdir $(PWD)/bin

test:
	go test -v github.com/torczuk/reptile/server/request/primary
	#pushd system_test; $(PWD)/system_test/run.sh; popd

build:
	go build -o bin/reptile