GOLANG_TAG?=wpenv_proto_golang
PYTHON_TAG?=wpenv_proto_python

all: build_golang build_python make_golang make_python

build_golang:
	docker build --tag=${GOLANG_TAG} \
		-f golang.Dockerfile .

build_python:
	docker build --tag=${PYTHON_TAG} \
		-f python.Dockerfile .

make_golang:
	docker run --rm --volume=${PWD}:/app/models \
		${GOLANG_TAG} \
		-I=. --go_out=. --go-grpc_out=. ./models/user.proto

make_python:
	docker run --rm --volume=${PWD}:/app/models \
		${PYTHON_TAG} \
		-I=. --python_out=. --grpc_python_out=. ./models/user.proto

.PHONY: golang python
