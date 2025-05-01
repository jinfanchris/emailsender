#!/bin/bash

set -e

ROOT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)
cd "$ROOT_DIR"

build_proto() {
    service=$1
    protoc \
        --proto_path=pkg/grpc/"${service}" \
        --go_out=. \
        --go-grpc_out=. \
        pkg/grpc/"${service}"/*.proto
}

build_proto mailer
