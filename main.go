package main

import "log"

func main() {
	err := startServer()
	if err != nil {
		log.Fatal("Server failed to", err)
	}
}
