# First stage
FROM golang:alpine3.13 as builder

WORKDIR /app/
COPY . /app 
RUN  go build

# Final stage
FROM alpine:3.13.0
WORKDIR /app/
COPY --from=builder /app/signal-generator . 

CMD ["/app/signal-generator"]
