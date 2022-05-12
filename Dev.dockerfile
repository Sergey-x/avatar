ARG GO_VERSION=1.18.1
FROM golang:${GO_VERSION}-alpine as avatar-deps
WORKDIR /avatar
COPY . .


FROM avatar-deps as avatar-builder
RUN go build -buildvcs=false -o avatar .


FROM avatar-builder as avatar-runner
ENTRYPOINT go run ./avatar
