# syntax=docker/dockerfile:1
FROM golang:latest
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY *.go ./
COPY /config/*.go ./config/
COPY /constants/*.go ./constants/
COPY /controllers/*.go ./controllers/
COPY /lang/*.go ./lang/
COPY /models/*.go ./models/
COPY /target/*.go ./target/
COPY *.go ./
COPY swagger.yaml ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app .
COPY --from=0 /usr/local/go/lib/time/zoneinfo.zip /
ENV ZONEINFO=/zoneinfo.zip
EXPOSE 5557

CMD ["./app"]