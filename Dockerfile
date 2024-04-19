FROM golang:1.21.5-alpine3.19 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o /server ./cmd/web/main.go

FROM alpine:3.18.2 AS final
WORKDIR /app
COPY .env .
COPY schema /app/schema
COPY --from=builder /server .
COPY ./shell-scripts/start.sh .
COPY ./shell-scripts/wait-for.sh .

EXPOSE 3000
CMD ["./server"]
ENTRYPOINT ["./app/start.sh"]