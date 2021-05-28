# Hello World

Hi! This repository features some very simple day or weekend projects :) These are not meant for a productive environment. Its basically a sketchbook to refine my programming skills and keep them up to date.

If you want to give Feedback feel free to open an issue

## Projects

### [mq-ampq-golang](./mq-ampq-golang)

A project written in go with a sender and receiver service which connects to RabbitMQ using the AMQP protocol. You can run it locally using docker-compose that comes with this project

### [mq-stomp-golang](./mq-stomp-golang)

A project written in go with a sender/receiver service which connects to ActiveMQ using the STOMP protocol. I also tinkered with a webclient you can use to publish messages. You can run it locally using docker-compose that comes with this project. 

### [unit-test-golang](./unit-test-golang)

A project written in go to get familiar with unit testing and benchmarking in go. Implements two fibonacci algorithms and benchmarks them. 

### [mongo-golang](./mongo-golang)

A project written in go. It is a small microservice with a simple "books" api which allows to perfrom CRUD operations on the mongo database. This project also uses the routing library [gin](https://pkg.go.dev/github.com/gin-gonic/gin) and implements some simple integration testing for those CRUD operations. 