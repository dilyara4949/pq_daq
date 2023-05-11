FROM golang:1.20.2-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main .
# RUN apt-get update; apt-get install -y nginx; 

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main /app/main
# COPY wait-for.sh .
EXPOSE 8080

CMD ["/app/main"]
