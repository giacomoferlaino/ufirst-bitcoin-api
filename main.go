package main

import (
	"log"
	"net/http"

	"github.com/gorilla/rpc/v2"
	"ufirst.com/bitcoin/bitcoin"
	"ufirst.com/bitcoin/jsonrpc"
)

const maxDaysDifference = 100

func main() {
	router := http.NewServeMux()
	rpcServer := rpc.NewServer()
	rpcServer.RegisterCodec(jsonrpc.NewCodec(), "application/json")
	rpcServer.RegisterService(bitcoin.NewService(uint(maxDaysDifference)), "bitcoin")
	router.Handle("/rpc", rpcServer)
	log.Fatal(http.ListenAndServe(":8080", router))
}
