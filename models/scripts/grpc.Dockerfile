ARG GO_GRPC_TAG=

FROM $GO_GRPC_TAG

WORKDIR /scripts
RUN printf '#!/bin/bash\n \
    go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc \
    && protoc $@' > ./install.sh \
    && chmod +x ./install.sh

WORKDIR /app
ENTRYPOINT [ "/scripts/install.sh" ]
