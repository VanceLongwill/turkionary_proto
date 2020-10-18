.PHONY: build

build:
	# protoc -I=./search --go_out=paths=source_relative:./search ./search/search.proto
	protoc -I=./search --go_out=./search --go_opt=paths=source_relative \
	    --go-grpc_out=./search --go-grpc_opt=paths=source_relative \
	    ./search/search.proto
