package main

import (
	"context"
	"fmt"
	"go-grpc-stream-api/proto"
	"google.golang.org/grpc"
	"log"
)

const host = "localhost:12345"

var nums  = []int32{1, 3, 4, 6, 9, 34}

func main() {
    conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

    client := proto.NewCalculatorClient(conn)
    stream, err := client.Add(context.Background())
    if err != nil {
    	log.Fatal(err)
	}


    for _, num := range nums {
    	err := stream.Send(&proto.Request{Value: num})
    	if err != nil {
    		log.Fatal(err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.GetResult())

}
