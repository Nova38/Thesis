FROM golang:1.21.4-bookworm as builder


WORKDIR /

COPY . .

RUN pwd && ls -la
# RUN go get -d -v ./

# RUN go get
ENV GOPRIVATE=github.com/nova38
RUN go env -w GOPRIVATE=github.com/nova38/*



RUN go build -o saacs -v ./pkg/chaincode/bin/main.go



FROM golang:1.21.4-bookworm

WORKDIR /

COPY --from=builder /saacs .

EXPOSE 9999
# ENV AUTH_MODE=roles

# CMD [ "/bin/sh" ]
CMD ["./saacs"]
