.PHONY: clean build test

build: clean test
	go build

test:
	go test -v -count=1 ./test/automated/...

clean:
	go clean -i github.com/thehungry-dev/line...
