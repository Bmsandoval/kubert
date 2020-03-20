package grpc

import (
	"fmt"
	"github.com/bmsandoval/kubert/configs"
	"github.com/bmsandoval/kubert/grpc/protos"
	"google.golang.org/grpc"
)

type Connection struct {
	Server *grpc.ClientConn
	GreeterClient protos.GreeterClient
}

var server Connection

func Start(config configs.Configuration) (*Connection, error){
	address := fmt.Sprintf("%s:%s", config.EkubeHost, config.EkubePort)
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	server = Connection{
		Server:        conn,
	}

	registerClients()

	return &server, nil
}

func registerClients() {
	c := protos.NewGreeterClient(server.Server)
	server.GreeterClient = c
}

func Stop() error {
	return server.Server.Close()
}
