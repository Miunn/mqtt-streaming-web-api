FROM golang:latest as BUILD

WORKDIR /var/www/api

COPY . .

RUN go mod tidy

RUN go build -o api .

FROM debian:latest as PROD

WORKDIR /var/www/api

COPY --from=BUILD --chown=www-data:www-data /var/www/api/api .

EXPOSE 8080

CMD ["/var/www/api/api"]
