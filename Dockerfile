FROM golang:alpine
EXPOSE 443

RUN apk update && apk add --no-cache git ca-certificates tzdata stunnel supervisor && update-ca-certificates

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
COPY proto go.mod go.sum main.go LICENSE README.md $GOPATH/src/mypackage/myapp/

RUN go mod download
RUN go mod verify

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/grpc-echo-server .

COPY stunnel.crt /etc/stunnel/stunnel.crt
COPY stunnel.key /etc/stunnel/stunnel.key
COPY stunnel.conf /etc/stunnel/

RUN mkdir /var/run/stunnel && \
    chown stunnel:root /var/run/stunnel && \
    chmod 0770 /var/run/stunnel

RUN mkdir /etc/supervisor.d
COPY server.ini /etc/supervisor.d/
CMD [ "/usr/bin/supervisord", "-n", "-c", "/etc/supervisord.conf"]
