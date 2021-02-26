FROM python:3.9.2-slim

RUN apt-get update \
    && python -m pip install --upgrade pip \
    && python -m pip install grpcio\
    && python -m pip install grpcio-tools

WORKDIR /app-install
RUN apt-get install -y git \
    && git clone https://github.com/googleapis/googleapis \
    && mkdir -p /usr/local/include/google/api \
    && cp -r ./googleapis/google/api/* /usr/local/include/google/api

WORKDIR /app
ENTRYPOINT [ "python", "-m", "grpc_tools.protoc", "-I/usr/local/include" ]
