package main

import (
	"fmt"
	"io"
	"log"
	"time"

	gocomm "github.com/mv-kan/go-comm"
)

func main() {
	log.SetOutput(io.Discard)
	fmt.Println("Start writer...")
	p, err := gocomm.NewPort("/tmp/virtual_dev_input", 115200, 0)
	if err != nil {
		panic(err)
	}
	input := make(chan string)
	conn, msgChan, err := gocomm.NewConnection(p, input, 0, "\n", "\n")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	go func() {
		for {
			msg := <-msgChan
			fmt.Printf("ACCEPT MESSAGE FROM READER \"%v\"\n", msg.Data)
		}
	}()
	i := 0
	for {
		input <- fmt.Sprintf("I am writer msg count = [%d]", i)
		time.Sleep(time.Second)
		i++
	}

}
