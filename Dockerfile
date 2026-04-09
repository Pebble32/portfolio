FROM golang:1.26-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o portfolio .

FROM alpine:latest
COPY --from=build /app/portfolio /portfolio
COPY --from=build /app/static /static
EXPOSE 3000
CMD ["/portfolio"]
