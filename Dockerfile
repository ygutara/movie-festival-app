FROM golang:1.18-alpine as builder

WORKDIR /app/
COPY . ./
RUN go mod download
RUN go build -o ./app/migration/migration ./app/migration/migration.go
RUN go build -o ./app/main ./app/main.go

FROM alpine:3
WORKDIR /app/
EXPOSE 8080
COPY --from=builder /app/app/migration/migration /app
COPY --from=builder /app/app/main /app
CMD /app/migration && /app/main