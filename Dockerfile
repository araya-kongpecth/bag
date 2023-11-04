FROM golang:1.20.4-alpine3.17

WORKDIR /cmd

COPY . .

RUN go build -o api

EXPOSE 9000

CMD [ "./api" ]