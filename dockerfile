# Build stage
FROM golang:1.17-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o title .
ENTRYPOINT ["./title"]
# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/title .
EXPOSE 3456
CMD ["./title"]