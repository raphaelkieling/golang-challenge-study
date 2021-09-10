#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d
RUN go build -o /go/bin/app
RUN go build -o /go/bin/populate ./pkg/populate/populate.go 

#final stage
FROM alpine:latest

# Usado para utilizar no wait-for-postgres.sh
RUN apk add postgresql
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
COPY --from=builder /go/bin/populate /populate
COPY --from=builder /go/src/app/resource /resource
RUN ["chmod", "+x", "/resource/wait-for-postgres.sh"]
LABEL Name=challengestudydeliverymuch Version=0.0.1
EXPOSE 3000
