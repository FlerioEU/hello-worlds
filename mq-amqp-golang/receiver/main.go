package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/streadway/amqp"
)

func main() {
	log.SetOutput(os.Stdout)
	log.Println("Starting Receiver-Service...")

	url, q := os.Getenv("AMQP_URL"), os.Getenv("AMQP_QUEUE")

	log.Printf("Connecting to broker on '%s'", url)
	log.Printf("Connecting to queue '%s'", q)

	mqConn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("Error while establishing: %v", err)
	}
	defer mqConn.Close()

	ch, err := mqConn.Channel()
	if err != nil {
		log.Fatalf("Error on mqConn channel: %v", err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(q, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Error subscribing to channel '%s': %v", q, err)
	}

	go func() {
		for msg := range msgs {
			log.Printf("Received message '%s'\n", msg.Body)
		}
	}()

	log.Println("Ready to receive messages")

	// taken from https://gist.github.com/p4tin/cefafe03401bd28f21c12cba99a60c79
	shutdown := make(chan int)
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		log.Println("Shutting down...")
		shutdown <- 1
	}()

	<-shutdown
}
