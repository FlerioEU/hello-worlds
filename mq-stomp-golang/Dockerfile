FROM golang:alpine3.13

RUN mkdir /app
ADD . /app

WORKDIR /app
RUN go build -o  main .

CMD ["/app/main"]