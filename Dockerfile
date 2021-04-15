FROM golang:alpine as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    APP_NAME=bank-transfer \
    MY_APP_PORT=5000

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -a --installsuffix cgo --ldflags="-s" -o main

FROM scratch

COPY --from=builder /build .

ENTRYPOINT ["./main"]

EXPOSE 5000