# STOMP World
This is a quick hello world project which aims to have a [STOMP](https://stomp.github.io/index.html) sender/receiver golang service which communicates with ActiveMQ. Also added a small Webpage to publish new messages

## Available Protocols 

to interact with message brokers

* AMQP
* MQTT
* STOMP (this project uses this one)

##  Prerequisites

* golang
* docker

## Run

Run the command
`docker-compose up`

If you made changes to code and want to see a difference when running via compose first you have to run
`docker-compose build`

When both services are running go to localhost:8080/mq to enter some values into the webpage. Values you enter there will be shown in the console. Only the console it part of the subscription. The webpage only shows the values you entered but they are not received via activemq

## Known Problems

* docker-compose: mq-talker starts up to soon and will restart at least once until activemq is ready
    * possible solutions: 
    https://stackoverflow.com/a/55504335 https://docs.docker.com/compose/startup-order/

