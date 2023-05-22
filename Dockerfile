FROM alpine:latest
EXPOSE 5050

COPY grpc-echo-server .
RUN chmod +x grpc-echo-server
ENTRYPOINT [ "./grpc-echo-server" ]
