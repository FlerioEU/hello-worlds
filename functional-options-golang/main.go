package main

import (
	"fmt"
	"time"

	"github.com/FlerioEU/hello-world/functional-options-golang/server"
)

func main() {
	fmt.Println("#### Server with Defaults ####")
	srv := server.New()
	fmt.Printf("%+v\n", srv)

	fmt.Println("#### Server with Options ####")
	srv = server.New(
		server.WithHost("127.0.0.1"),
		server.WithPort(443),
		server.WithTimeout(time.Hour),
		// server.WithMaxConn(5), // <- will use a default value instead
	)
	fmt.Printf("%+v\n", srv)
}
