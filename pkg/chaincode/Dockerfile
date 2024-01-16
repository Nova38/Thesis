FROM golang:1.21.4-bookworm as builder


WORKDIR /

COPY . .

RUN pwd && ls -la
# RUN go get -d -v ./

# RUN go get
ENV GOPRIVATE=github.com/nova38
RUN go env -w GOPRIVATE=github.com/nova38/*



RUN go build -gcflags="all=-N -l" -o saacs -v ./pkg/chaincode/bin/main.go
# RUN go install github.com/go-delve/delve/cmd/dlv@latest



FROM golang:1.21.4-bookworm

WORKDIR /

COPY --from=builder /saacs .
# COPY --from=builder /go/bin/dlv .

EXPOSE 9999
ENV AUTH_MODE=roles

# CMD [ "/bin/sh" ]
CMD ["/saacs"]
# CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/saacs"]
