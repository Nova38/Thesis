FROM golang:1.22-bookworm as builder


WORKDIR /

COPY go.work /go.work
COPY pkg /pkg
COPY cmd/protoc-gen-saacs-go/ cmd/protoc-gen-saacs-go/

# RUN pwd && ls -la
# RUN go get -d -v ./

# RUN go get
ENV GOPRIVATE=github.com/nova38
RUN go env -w GOPRIVATE=github.com/nova38/*



RUN go build -gcflags="all=-N -l" -o saacs -v ./pkg/saacs-cc/bin/main.go
# RUN go install github.com/go-delve/delve/cmd/dlv@latest



FROM golang:1.21.4-bookworm

WORKDIR /

COPY --from=builder /saacs .
# COPY --from=builder /go/bin/dlv .

EXPOSE 9999


# CMD [ "/bin/sh" ]
CMD ["/saacs"]
# CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/saacs"]
