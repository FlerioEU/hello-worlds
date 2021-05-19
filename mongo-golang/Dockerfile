FROM golang:alpine3.13

RUN mkdir /app
ADD . /app

WORKDIR /app
RUN go build -o  main .

ADD https://raw.githubusercontent.com/eficode/wait-for/v2.1.0/wait-for /wait-for
RUN chmod +x /wait-for

CMD ["/app/main"]