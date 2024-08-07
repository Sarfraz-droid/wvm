package main

import (
	"fmt"
	"sync"
	service "wvm/pkg/service"
	socket "wvm/pkg/socket"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(2)
	go service.Init()
	go socket.Init()

	wg.Wait()
	fmt.Printf("Services have stopped")
}
