FROM golang:alpine as builder
RUN mkdir /app
COPY . /app
WORKDIR /app

RUN go build -o viralgame .

FROM alpine as builder2
COPY --from=builder /app/viralgame /app/
WORKDIR /app
COPY --from=builder /app/.env .
ENTRYPOINT ["./viralgame"]