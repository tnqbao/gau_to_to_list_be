FROM golang:1.23-alpine AS builder
WORKDIR /gau_to_do_list
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -tags '!dev' -o main .

FROM alpine:latest
WORKDIR /gau_to_do_list
COPY --from=builder /gau_to_do_list/main .
EXPOSE 8088
CMD ["./main"]