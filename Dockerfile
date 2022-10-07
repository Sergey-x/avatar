ARG GO_VERSION=1.19.2
FROM golang:${GO_VERSION}-alpine as avatar-deps
WORKDIR /avatar
COPY . .
RUN go mod tidy
RUN go build -buildvcs=false -o avatar .


FROM avatar-deps as avatar-runner
WORKDIR /avatar
COPY --from=avatar-deps /avatar/avatar .
RUN chmod 755 ./avatar
ENTRYPOINT ./avatar
