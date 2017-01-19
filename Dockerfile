FROM golang:latest

EXPOSE 8080

ENV GIN_MODE=release

WORKDIR /go/src/app
ADD ./main.go /go/src/app/main.go
RUN go get && go build main.go

CMD ./main
