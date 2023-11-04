FROM golang:1.21.10-alpine3.18

WORKDIR /cmd

COPY . .

RUN go build -o api

EXPOSE 9000

CMD [ "./api" ]