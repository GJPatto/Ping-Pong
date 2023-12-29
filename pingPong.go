package main

import (
	"fmt"
	"time"
)

func main() {

	for {

		fmt.Println("PING")

		time.Sleep(1 * time.Second)

		fmt.Println("PONG")

		time.Sleep(1 * time.Second)
	}

}
