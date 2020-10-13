package main

import (
	"go-grpc-stream-api/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)
const port = ":12345"

func startServer() error{
    conn, err := net.Listen("tcp", port)
    if err  != nil {
    	return err
	}

	server := grpc.NewServer()

	proto.RegisterCalculatorService(server, &proto.CalculatorService{
		Add: Add,
	})

	log.Println("<- Server starting at port 12345 ->")
	if err = server.Serve(conn); err != nil {
		return err
	}

	return nil
}

var result = int32(0)

func Add(server proto.Calculator_AddServer)  error{
	for {
		in, err := server.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		result += in.GetValue()

	}

	return server.SendAndClose(&proto.Response{Result: result})
}