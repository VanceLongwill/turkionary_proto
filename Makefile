.PHONY: build

build:
	protoc -I=./search --go_out=paths=source_relative:./search ./search/search.proto
