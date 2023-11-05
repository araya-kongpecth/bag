<<<<<<< HEAD
# Use the official Golang image
FROM golang:1.21.1-alpine
=======
FROM golang:1.21.1
>>>>>>> f6808698cf1fc5c0913c6e085840c8c7ad8fdda1

WORKDIR /app

# Copy the entire project into the container
COPY . .

RUN go build -o bag ./cmd

EXPOSE 9000

<<<<<<< HEAD
CMD ["./bag"]
=======
CMD [ "./api" ]
>>>>>>> f6808698cf1fc5c0913c6e085840c8c7ad8fdda1
