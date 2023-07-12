package server

import "time"

// credit where credit's due:
// https://www.youtube.com/watch?v=MDy7JQN5MN4&list=WL&index=4&ab_channel=AnthonyGG
// https://golang.cafe/blog/golang-functional-options-pattern.html
type Server struct {
	host    string
	port    int
	timeout time.Duration
	maxConn int
}

func New(options ...func(*Server)) *Server {
	srv := withDefaults()
	for _, fn := range options {
		fn(srv)
	}
	return srv
}

func WithHost(host string) func(*Server) {
	return func(s *Server) {
		s.host = host
	}
}

func WithPort(port int) func(*Server) {
	return func(s *Server) {
		s.port = port
	}
}

func WithTimeout(timeout time.Duration) func(*Server) {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithMaxConn(maxConn int) func(*Server) {
	return func(s *Server) {
		s.maxConn = maxConn
	}
}

func withDefaults() *Server {
	return &Server{
		host:    "http://localhost",
		port:    80,
		timeout: time.Second * 30,
		maxConn: 10,
	}
}
