package main

import (
	"fmt"
	"io"
	"log"

	gocomm "github.com/mv-kan/go-comm"
)

func main() {
	log.SetOutput(io.Discard)
	fmt.Println("Start reader...")
	p, err := gocomm.NewPort("/tmp/virtual_dev_output", 115200, 0)
	if err != nil {
		panic(err)
	}
	input := make(chan string)
	conn, msgChan, err := gocomm.NewConnection(p, input, 0, "\n", "\n")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	input <- "hello from reader"
	for {
		msg := <-msgChan
		fmt.Printf("accepted message = %s\n", msg.Data)
	}
}
