FROM golang:1.12

ENV GO111MODULE=on

WORKDIR /bot
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN go build ./bot/main.go

CMD ["./main"]