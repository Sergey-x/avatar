ARG GO_VERSION=1.19.2

FROM golang:${GO_VERSION}-alpine as avatar-deps
WORKDIR /avatar

COPY . .
RUN go mod tidy
RUN go build -buildvcs=false -o avatar .
RUN chmod 755 ./avatar


FROM alpine:latest as built_avatar
WORKDIR /avatar
COPY --from=avatar-deps /avatar/avatar /avatar/
ENTRYPOINT ./avatar
