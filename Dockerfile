FROM golang:1.26 AS build
WORKDIR /app
COPY . .
RUN go build -o portfolio .

FROM gcr.io/distroless/static
COPY --from=build /app/portfolio /portfolio
COPY --from=build /app/static /static
EXPOSE 3000
CMD ["/portfolio"]

