FROM python:3.9.2-slim

RUN apt-get update \
    && python -m pip install --upgrade pip \
    && python -m pip install grpcio\
    && python -m pip install grpcio-tools

WORKDIR /app
ENTRYPOINT [ "python", "-m", "grpc_tools.protoc" ]
