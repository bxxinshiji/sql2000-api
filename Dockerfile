FROM bigrocs/golang-gcc:1.13 as builder

WORKDIR /go/src/github.com/bxxinshiji/sql2000-api
COPY . .

ENV GO111MODULE=on CGO_ENABLED=1 GOOS=linux GOARCH=amd64
RUN go build -a -installsuffix cgo -o bin/sql2000Api

FROM bigrocs/alpine:ca-data

RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

COPY --from=builder /go/src/github.com/bxxinshiji/sql2000-api/bin/sql2000Api /usr/local/bin/
CMD ["sql2000Api"]
EXPOSE 8080