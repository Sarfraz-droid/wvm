package socket

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"wvm/pkg/config"
)

type Args struct{}

type RPCHandler int64

func (t *RPCHandler) Reload(args *Args, reply *string) error {
	// Set the value at the pointer got from the client
	config.LoadConfig()
	*reply = "Config reloaded"
	return nil
}

func Init() {
	config.LoadConfig()
	config := config.ConfigData
	rpchandler := new(RPCHandler)
	// Register the timeserver object upon which the GiveServerTime
	// function will be called from the RPC server (from the client)
	rpc.Register(rpchandler)
	// Registers an HTTP handler for RPC messages
	rpc.HandleHTTP()
	// Start listening for the requests on port 1234
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.ClientPort))
	if err != nil {
		log.Fatal("Listener error: ", err)
	}
	// Serve accepts incoming HTTP connections on the listener l, creating
	// a new service goroutine for each. The service goroutines read requests
	// and then call handler to reply to them
	log.Printf("Server is running on port %d", config.ClientPort)
	http.Serve(listener, nil)
}
