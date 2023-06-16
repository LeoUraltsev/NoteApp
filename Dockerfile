FROM golang:1.20

WORKDIR /go/NotesApp

COPY ./ ./
RUN go mod download

RUN go .build -o ./..build/main ./cmd/main/main.go
ENTRYPOINT ["./.build/main"]