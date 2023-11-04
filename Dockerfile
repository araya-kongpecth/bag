FROM golang:1.21.1

WORKDIR /cmd

COPY . .

RUN go build -o api

EXPOSE 9000

CMD [ "./api" ]
