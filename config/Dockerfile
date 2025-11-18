FROM golang:1.24 as build
WORKDIR /app
COPY . .
RUN CGO_ENABLE=0 GOOD=linux GOARCH=amd64 go build -o cloudrun

FROM scratch
WORKDIR /app
COPY --from=build /app/cloudrun .
ENTRYPOINT ["./cloudrun"]