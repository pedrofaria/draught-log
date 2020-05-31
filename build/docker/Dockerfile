FROM golang:1.14-alpine AS builder

RUN apk add --update nodejs npm build-base

WORKDIR /go/src/app
COPY ./ ./

RUN make deps build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=builder /go/src/app/draught-log .

EXPOSE 80

CMD ["./draught-log", "-config=draught-log.config.yml"]
