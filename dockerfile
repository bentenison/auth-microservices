FROM golang:1.22-alpine3.20 as builder

WORKDIR /app
COPY go.mod .
RUN go mod download
COPY .env .
COPY . .
RUN go build -ldflags="-w -s" -o auth-service ./auth-server/cmd/server/main.go

FROM scratch
COPY --from=builder /app/auth-service /app/.env ./
EXPOSE 35580 50051
CMD [ "./auth-service" ]