FROM golang:buster

RUN apt-get update
RUN apt-get install unzip

WORKDIR /wd

RUN curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v3.15.8/protoc-3.15.8-linux-x86_64.zip
RUN unzip protoc-3.15.8-linux-x86_64.zip -d $HOME/.local
RUN export PATH="$PATH:$HOME/.local/bin"

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

CMD $HOME/.local/bin/protoc --proto_path=/input --go_opt=paths=source_relative --go_out=/output --go-grpc_opt=paths=source_relative --go-grpc_out=/output /input/**/*.proto --experimental_allow_proto3_optional