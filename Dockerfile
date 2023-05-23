FROM golang:alpine as builder
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

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

WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .

RUN go mod download
RUN go mod verify

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/grpc-echo-server .

FROM scratch
EXPOSE 5050

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /go/bin/grpc-echo-server /go/bin/grpc-echo-server
# RUN chmod +x /bin/grpc-echo-server
ENTRYPOINT [ "/go/bin/grpc-echo-server" ]
