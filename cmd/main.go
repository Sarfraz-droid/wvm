package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	config_service "wvm/pkg/config"
)

type Args struct{}

func main() {
	var reply string
	args := Args{}

	org_key := os.Args[1]

	config_service.LoadConfig()
	config := config_service.ConfigData
	client, err := rpc.DialHTTP("tcp", fmt.Sprintf("localhost:%d", config.ClientPort))

	if err != nil {
		log.Fatal("Client connection error: ", err)
	}

	switch org_key {
	case "reload":
		err = client.Call("RPCHandler.Reload", args, &reply)
	}

	if err != nil {
		log.Fatal("Client invocation error: ", err)
	}

	// Print the reply from the server
	fmt.Println(reply)
}
