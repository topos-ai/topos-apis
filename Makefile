PWD=`pwd`

all: pkg-go pkg-python

pkg-go: includes/googleapis genproto/go
	docker run --rm -v $(PWD)/includes:/includes:ro -v $(PWD)/topos:/protos/topos:ro -v $(PWD)/genproto/go:/genproto/go golang:1.12.1-alpine sh -c \
	'set -ex \
		&& apk add git protobuf-dev \
		&& cd /genproto/go \
		&& go mod init github.com/topos-ai/topos-apis/genproto/go \
		&& go get github.com/golang/protobuf/protoc-gen-go \
		&& find /protos/topos -mindepth 2 -maxdepth 2 -type d -exec sh -c "protoc \
			-I/includes/googleapis \
			-I/protos \
			--go_out=plugins=grpc:/genproto/go \
			'{}'/*.proto" \; \
		&& go mod tidy'

pkg-python: includes/googleapis genproto/python
	docker run --rm -v $(PWD)/includes:/includes:ro -v $(PWD)/topos:/protos/topos:ro -v $(PWD)/genproto/python:/genproto/python python:3.7.3 sh -c \
	'set -ex \
		&& pip install grpcio-tools \
		&& find /protos/topos -mindepth 2 -maxdepth 2 -type d -exec sh -c "python \
			-m grpc_tools.protoc \
			-I/includes/googleapis \
			-I/protos \
			--python_out=/genproto/python \
			--grpc_python_out=/genproto/python \
			'{}'/*.proto" \;'

includes/googleapis: includes
	git clone --depth 1 https://github.com/googleapis/googleapis.git includes/googleapis

includes:
	mkdir includes

genproto/go: genproto
	mkdir genproto/go

genproto/python: genproto
	mkdir genproto/python

genproto:
	mkdir genproto

clean:
	rm -rf genproto includes

.PHONY:	all clean pkg-go pkg-python
