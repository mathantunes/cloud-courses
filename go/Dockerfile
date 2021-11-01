FROM golang:latest

WORKDIR /server

COPY . .

RUN go mod tidy

RUN go build -o ./bin

CMD ["/bin/go"]
