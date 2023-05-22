FROM alpine:latest
EXPOSE 5050

COPY grpc-echo-server /bin
RUN chmod +x /bin/grpc-echo-server
ENTRYPOINT [ "grpc-echo-server" ]
