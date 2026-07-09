package main

import (
	"log"
	"net/rpc"
)

type Move struct {
	Color int
	Col   int
}

type Board struct {
	BoardString string
}

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// First send one move to the server
	move := Move{
		Color: 0, // 0 means White
		Col:   3, // column 3
	}

	var replyMove int

	err = client.Call("ConnectGame.Move", &move, &replyMove)
	if err != nil {
		log.Fatal("move error:", err)
	}

	log.Printf("Move reply: %v", replyMove)

	// Then ask the server for the board
	var reply Board
	var args int

	err = client.Call("ConnectGame.Get", &args, &reply)
	if err != nil {
		log.Fatal("game error:", err)
	}

	log.Printf("Game:\n%v", reply.BoardString)
}
