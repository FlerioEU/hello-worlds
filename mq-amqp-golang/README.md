# AMQP World
This is a quick hello world project which aims to have a [AMQP](https://www.amqp.org/) sender/receiver golang services which communicates with RabbitMQ. Sadly the [used module](https://github.com/streadway/amqp) does not support AMQP 1.0 which is required for ActiveMQ

The sender service will publish a message every three seconds to the configured broker. The receiver service subscribes to a channel that the sender service publishes too and logs the message.

After starting you can access the dashboard of RabbitMQ [here](http://localhost:15672/)

To wait for the broker to start up i am using the [wait-for](https://github.com/eficode/wait-for) script provides by [eficode](https://github.com/eficode)

## Available Protocols 

to interact with message brokers

* AMQP (This project uses this one)
* MQTT
* STOMP 

##  Prerequisites

* golang
* docker

## Setup

Run `go mod download` in both the sender and receiver folder to retrieve the needed module

## Run

Run the command
`docker-compose up`

If you made changes to code and want to see a difference when running via compose first you have to run
`docker-compose build`

