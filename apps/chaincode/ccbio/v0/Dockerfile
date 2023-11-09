
ARG GO_VER=1.21.1
ARG ALPINE_VER=3.17
FROM golang:${GO_VER}-alpine${ALPINE_VER}

ENV CORE_CHAINCODE_ID_NAME=biochain

WORKDIR /app/chaincode
COPY . .

WORKDIR /app/chaincode/bin/ccas

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 9999
# CMD [ "/bin/sh" ]
CMD ["bin"]
