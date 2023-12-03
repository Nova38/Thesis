FROM golang:1.21.4-bookworm as builder


WORKDIR /app

COPY . .

RUN pwd && ls -la
# RUN go get -d -v ./

# RUN go get
ENV GOPRIVATE=github.com/nova38
RUN go env -w GOPRIVATE=github.com/nova38/*




RUN go build \
    -ldflags="-X 'github.com/nova38/thesis/lib/go/fabric/auth/common.EnableHiddenTx=0 '" \
    -ldflags="-X 'github.com/nova38/thesis/lib/go/fabric/auth/common.EnabledSuggestions=0 '" \
    -o bin/noauth-no-sub-items  -v /app/lib/go/fabric/contracts/noauth/bin/main.go

FROM golang:1.21.4-bookworm

WORKDIR /app

COPY --from=builder /app/bin/ .

EXPOSE 9999
# CMD [ "/bin/sh" ]
CMD ["/app/noauth-no-sub-items"]
