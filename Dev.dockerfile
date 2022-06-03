ARG GO_VERSION=1.18.1
FROM golang:${GO_VERSION}-alpine as avatar-deps
WORKDIR /avatar
COPY . .
#FROM avatar-deps as avatar-builder
RUN go build -buildvcs=false -o avatar .


FROM avatar-deps as avatar-runner
WORKDIR /avatar
COPY --from=avatar-deps /avatar/avatar .
RUN chmod 755 ./avatar
ENTRYPOINT ./avatar
