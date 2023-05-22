FROM alpine:latest
EXPOSE 5050

ADD grpc-echo-server .
ENTRYPOINT [ "./grpc-echo-server" ]
