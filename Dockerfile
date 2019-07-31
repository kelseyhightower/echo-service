FROM golang:1.12.7-stretch as builder
ENV GO111MODULE=on
WORKDIR /module
COPY . /module/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -a -tags netgo \
      -ldflags '-w -extldflags "-static"' \
      -mod vendor \
      -o echo

FROM scratch
COPY --from=builder /module/echo .
ENTRYPOINT ["/echo"]
