FROM golang:1.18-alpine AS build_base

RUN apk add --no-cache git

WORKDIR /tmp/scaling-spoon

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/scaling-spoon ./cmd/scaling-spoon


FROM alpine

COPY --from=build_base /tmp/scaling-spoon/out/scaling-spoon /app/scaling-spoon

# COPY --from=build_base /tmp/scaling-spoon/sql /sql

ENV GODEBUG madvdontneed=1

EXPOSE 50051

CMD ["/app/scaling-spoon"]
