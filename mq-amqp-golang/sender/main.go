package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	log.SetOutput(os.Stdout)
	log.Println("Starting Sender-Service...")

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

	_, err = ch.QueueDeclare(q, false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Error on '%s' channel: %v", q, err)
	}

	// publish a message every 3 seconds
	go func() {
		for {
			msg := amqp.Publishing{
				Timestamp:   time.Now(),
				ContentType: "text/plain",
				Body:        []byte("Hello World"),
			}

			log.Println("Sending a new message...")
			err = ch.Publish("", q, false, false, msg)
			if err != nil {
				log.Fatalf("Error publishing message on '%s' channel: %v", q, err)
			}

			time.Sleep(3 * time.Second)
		}
	}()

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
