PWD=`pwd`

all: pkg-go pkg-python

endpoints: pkg-descriptors
	gcloud components update \
		&& find topos -name api_config.yaml -execdir gcloud --project topos-ai endpoints services deploy api_config.yaml api_descriptor.pb \;

endpoints-dev: api-config-dev pkg-descriptors
	gcloud components update \
		&& find topos -name api_config_dev.yaml -execdir gcloud --project topos-ai-dev endpoints services deploy api_config_dev.yaml api_descriptor.pb \;

api-config-dev:
	find topos -name api_config.yaml -execdir sh -c "sed -E 's/\.([a-z-]+)\.endpoints/-dev.\1.endpoints/' api_config.yaml > api_config_dev.yaml" \;

pkg-go: includes/googleapis genproto/go
	docker run --rm -v $(PWD)/includes:/includes:ro -v $(PWD)/topos:/protos/topos:ro -v $(PWD)/genproto/go:/genproto/go golang:1.13.7-alpine sh -c \
	'set -ex \
		&& apk add git protobuf-dev \
		&& cd /genproto/go \
		&& go mod init github.com/topos-ai/topos-apis/genproto/go \
		&& go get github.com/golang/protobuf/protoc-gen-go \
		&& find /protos/topos -type d | xargs -I{.} find {.} -maxdepth 1 -name \*.proto -exec \
			protoc -I/includes/googleapis \
				-I/protos \
				--go_out=paths=source_relative,plugins=grpc:/genproto/go \
				{} + \
		&& go mod tidy'

pkg-python: includes/googleapis genproto/python
	docker run --rm -v $(PWD)/includes:/includes:ro -v $(PWD)/topos:/protos/topos:ro -v $(PWD)/genproto/python:/genproto/python python:3.8.1 sh -c \
	'set -ex \
		&& pip install --upgrade pip \
		&& pip install grpcio-tools \
		&& find /protos/topos -type d | xargs -I{.} find {.} -maxdepth 1 -name \*.proto -exec \
			python \
				-m grpc_tools.protoc \
				-I/includes/googleapis \
				-I/protos \
				--python_out=/genproto/python \
				--grpc_python_out=/genproto/python \
				{} + '

pkg-descriptors: includes/googleapis
	docker run --rm -v $(PWD)/includes:/includes:ro -v $(PWD)/topos:/protos/topos alpine:3.11.3 sh -c \
	'set -ex \
		&& apk add protobuf-dev \
		&& find /protos/topos -mindepth 2 -maxdepth 2 -type d -exec sh -c "protoc \
			-I/includes/googleapis \
			-I/protos \
			--proto_path=. \
			--include_imports \
			--include_source_info \
			--descriptor_set_out='{}'/api_descriptor.pb \
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
	rm -rf genproto includes \
		&& find topos -name api_descriptor.pb -exec rm '{}' \; \
		&& find topos -name api_config_dev.yaml -exec rm '{}' \;

.PHONY:	all clean api-config-dev endpoints endpoints-dev pkg-go pkg-python pkg-descriptors
