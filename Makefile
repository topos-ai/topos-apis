PWD=`pwd`

all: pkg-go pkg-python

pkg-go: includes/googleapis genproto/go
	docker run --rm -v $(PWD)/includes:/includes:ro -v $(PWD)/topos:/protos/topos:ro -v $(PWD)/genproto/go:/genproto/go golang:1.12.0-alpine sh -c \
	'set -ex \
		&& apk add git protobuf-dev \
		&& go get github.com/golang/protobuf/protoc-gen-go \
		&& find /protos/topos -name '*.proto' | xargs protoc \
			-I/includes/googleapis \
			-I/protos \
			--go_out=plugins=grpc:/genproto/go \
		&& cd /genproto/go \
		&& go mod init github.com/topos-ai/topos-apis/genproto/go \
		&& go mod tidy'

pkg-python: includes/googleapis genproto/python
	docker run --rm -v $(PWD)/includes:/includes:ro -v $(PWD)/topos:/protos/topos:ro -v $(PWD)/genproto/python:/genproto/python python:3.7.2 sh -c \
	'set -ex \
		&& pip install grpcio-tools \
		&& find /protos/topos -name '*.proto' | xargs python \
			-m grpc_tools.protoc \
			-I/includes/googleapis \
			-I/protos \
			--python_out=/genproto/python \
			--grpc_python_out=/genproto/python'

includes/googleapis: includes
	git clone --depth 1 https://github.com/googleapis/googleapis.git includes/googleapis

includes:
	mkdir includes

genproto/go: genproto
	mkdir genproto/python

genproto/python: genproto
	mkdir genproto/python

genproto:
	mkdir genproto

clean:
	rm -rf genproto includes

.PHONY:	all clean pkg-go pkg-python
