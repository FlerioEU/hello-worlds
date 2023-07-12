package main

import (
	"context"
	"errors"
	"log"
	"time"
)

func main() {

	log.Println("#### First Request (Timeout) ####")
	start := time.Now()
	str, err := fetch("123", time.Millisecond*200)
	if err != nil {
		log.Printf("err: %v", err)
	} else {
		log.Printf("took %v", time.Since(start))
		log.Printf("val %s", str)
	}

	log.Println("#### Second Request (Success) ####")
	start = time.Now()
	str, err = fetch("123", time.Millisecond*700)
	if err != nil {
		log.Printf("err: %v", err)
	} else {
		log.Printf("val: %s", str)
		log.Printf("took %v", time.Since(start))
	}
}

type Response struct {
	value string
	err   error
}

func fetch(userID string, timeout time.Duration) (string, error) {
	context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel() // releases resources if slowOperation completes before timeout elapses

	ch := make(chan Response) // does not need to be closed: https://stackoverflow.com/questions/8593645/is-it-ok-to-leave-a-channel-open
	go func() {
		str, err := networkCall(userID)
		ch <- Response{
			value: str,
			err:   err,
		}
	}()

	// react to either timeout or successful response
	for {
		select {
		case <-ctx.Done():
			return "", errors.New("timeout exceeded")
		case resp := <-ch:
			return resp.value, resp.err
		}
	}
}

// simulate a network request that takes 600ms
func networkCall(userID string) (string, error) {
	time.Sleep(time.Millisecond * 600)

	return "123", nil
}
