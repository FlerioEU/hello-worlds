package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/go-stomp/stomp/v3"
)

const path string = "helloworld/test"

var tpl *template.Template
var mqConn *stomp.Conn
var values []string

type handleSub func(<-chan *stomp.Message)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	mqConn = initMQ(os.Getenv("STOMPHOST"))
	defer mqConn.Disconnect()

	subscribe(path, logSub)

	http.HandleFunc("/mq", serve)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func serve(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("message")

	publish(path, []byte(v))

	values = append(values, v)
	tpl.ExecuteTemplate(w, "main.gohtml", values)
}

func initMQ(host string) *stomp.Conn {
	if host != "" {
		host = "localhost"
	}

	// heartbeaterror reason when empty queue: https://stackoverflow.com/a/39951033
	mqConn, err := stomp.Dial("tcp", host+":61613", stomp.ConnOpt.HeartBeatError(360*time.Second))
	if err != nil {
		fmt.Print(err)
	}
	return mqConn
}

func publish(path string, payload []byte) {
	mqConn.Send(path, "text/plain", payload, nil)
}

func subscribe(path string, fn handleSub) {
	sub, err := mqConn.Subscribe(path, stomp.AckClient, stomp.SubscribeOpt.Id("1337"))
	if err != nil {
		fmt.Println("Error while subscribing to test")
	}
	go fn(sub.C)
}

func logSub(c <-chan *stomp.Message) {
	for m := range c {
		m.Conn.Ack(m)
		s := string(m.Body)
		fmt.Println("Value from subscription: " + s)
	}
}
