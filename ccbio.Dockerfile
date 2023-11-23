FROM golang:1.21.4-bookworm


COPY . .
RUN ls

WORKDIR /app/chaincode/auth/noauth





RUN ls

RUN go get -d -v ./
RUN go install -v ./...


EXPOSE 9999
# CMD [ "/bin/sh" ]
CMD ["bin"]
