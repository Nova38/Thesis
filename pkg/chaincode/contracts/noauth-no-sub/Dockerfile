FROM golang:1.21.4-bookworm as builder


WORKDIR /app

COPY . .

RUN pwd && ls -la
# RUN go get -d -v ./

# RUN go get
ENV GOPRIVATE=github.com/nova38
RUN go env -w GOPRIVATE=github.com/nova38/*




RUN go build -o bin/noauth-no-sub-items -v ./pkg/saacs/bin/main.go


FROM golang:1.21.4-bookworm

WORKDIR /app

COPY --from=builder /app/bin/ .

EXPOSE 9999
ENV AUTH_MODE=noauth-no-sub
# CMD [ "/bin/sh" ]
CMD ["/app/noauth-no-sub-items"]
