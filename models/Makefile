GOLANG_TAG?=wpenv_proto_golang
PYTHON_TAG?=wpenv_proto_python
GRPC_TAG?=wpenv_proto_grpc

all: build_golang build_python build_grpc_gateway gen_golang gen_python gen_grpc_gateway
build: build_golang build_grpc_gateway build_python
gen: gen_golang gen_python gen_grpc_gateway

build_golang:
	docker build --tag=${GOLANG_TAG} \
		-f ./scripts/golang.Dockerfile .

build_python:
	docker build --tag=${PYTHON_TAG} \
		-f ./scripts/python.Dockerfile .

build_grpc_gateway: build_golang
	docker build --tag=${GRPC_TAG} \
		--build-arg GO_GRPC_TAG=${GOLANG_TAG} \
		-f ./scripts/grpc.Dockerfile .

gen_grpc_gateway:
	docker run --rm --volume=${PWD}:/app \
		${GRPC_TAG} \
	-I . --grpc-gateway_out ./go \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt generate_unbound_methods=true \
    ./user.proto

gen_grpc_openapi:
	docker run --rm --volume=${PWD}:/app \
		${GRPC_TAG} \
	-I . --openapiv2_out ./openapi/gen \
	--openapiv2_opt logtostderr=true \
	./user.proto

gen_golang:
	docker run --rm --volume=${PWD}:/app \
		${GOLANG_TAG} \
		-I=. --go_out=./go --go-grpc_out=./go user.proto

gen_python:
	docker run --rm --volume=${PWD}:/app \
		${PYTHON_TAG} \
		-I=. --python_out=./python/gen --grpc_python_out=./python/gen user.proto


.PHONY: all build gen
