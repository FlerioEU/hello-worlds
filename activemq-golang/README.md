# STOMP World
This is a small hello world project which aims to have a [STOMP](https://stomp.github.io/index.html) sender/receiver golang service which communicates with ActiveMQ. Also added a small Webpage to publish new messages

## TODO

* build image and run with docker-compose
* rename folder to mq-stomp-golang

## Available Protocols 

to interact with message queues

* AMQP
* MQTT
* STOMP (this project uses this one)

##  Prerequisites

* golang
* docker

## Run

run the command
`docker-compose up`