FROM golang:alpine AS builder

WORKDIR /app
COPY . .
RUN go get ./...
RUN go build -o app .

FROM golang:alpine
COPY --from=builder /app/app /bin
ENTRYPOINT ["app"]
