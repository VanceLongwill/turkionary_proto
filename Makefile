.PHONY: build

build:
	protoc -I=./search --go_out=./search ./search/search.proto
