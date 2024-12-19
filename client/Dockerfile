FROM golang:alpine as builder
WORKDIR $GOPATH/src/mypackage/myapp/
COPY main.go go.sum go.mod $GOPATH/src/mypackage/myapp/

ENV USER=appuser
ENV UID=10001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"


RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/grpc-echo-client .

FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /go/bin/grpc-echo-client /go/bin/grpc-echo-client
USER appuser
ENTRYPOINT ["/go/bin/grpc-echo-client"]