FROM golang:1.25.3 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /api-monitoring

FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=build /api-monitoring /api-monitoring
COPY .env /app/.env
EXPOSE 8080
ENTRYPOINT ["/api-monitoring"]