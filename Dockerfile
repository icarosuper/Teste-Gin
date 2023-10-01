FROM golang:1.21.0-bullseye AS build

WORKDIR /app

COPY . .

RUN go build -o /build

EXPOSE 8080

CMD ["/build"]

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /build /build

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/build"]