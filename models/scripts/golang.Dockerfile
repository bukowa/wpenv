FROM golang

RUN apt-get update \
    && apt-get install -y curl unzip git \
    && curl -L https://github.com/protocolbuffers/protobuf/releases/download/v3.15.3/protoc-3.15.3-linux-x86_64.zip --output proto.zip \
    && unzip -o proto.zip -d /usr/local bin/protoc \
    && unzip -o proto.zip -d /usr/local 'include/*' \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


WORKDIR /app-install
RUN git clone https://github.com/googleapis/googleapis
RUN cp -r ./googleapis/google/api/ /usr/local/include/google/

WORKDIR /app
RUN rm -frd /app-install

ENTRYPOINT [ "protoc" ]
