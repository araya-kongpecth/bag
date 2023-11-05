# Use the official Golang image
FROM golang:1.21.1-alpine

WORKDIR /app

# Copy the entire project into the container
COPY . .

RUN go build -o bag ./cmd

EXPOSE 9000

CMD ["./bag"]
