# Build Stage
FROM golang:1.23-alpine AS builder
WORKDIR /gau_to_do_list
COPY go.mod go.sum ./
RUN go mod download
COPY . .
## build with release tag
RUN go build -tags release /cmd/ -o main

# Runtime Stage
FROM alpine:latest
WORKDIR /gau_to_do_list
COPY --from=builder /gau_to_do_list/main .
EXPOSE 8088
CMD ["./main"]
