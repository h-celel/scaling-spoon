#!/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

docker build -f $SCRIPT_DIR/protoc.Dockerfile -t protoc .
docker run --rm  -v $SCRIPT_DIR/../proto-spoon/proto:/input -v $SCRIPT_DIR/../proto:/output protoc