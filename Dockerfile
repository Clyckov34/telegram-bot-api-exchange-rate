FROM golang:1.18-alpine as bilder
RUN mkdir /app
COPY . /app
WORKDIR /app
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/telegramBot/main.go


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root
ENV RU UTC
COPY --from=bilder /app .
CMD ["./app"]