package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/rpc/v2"
	"ufirst.com/bitcoin/bitcoin"
	"ufirst.com/bitcoin/jsonrpc"
)

var port = os.Args[1]

const maxDaysDifference = 100

func main() {
	router := http.NewServeMux()
	rpcServer := rpc.NewServer()
	rpcServer.RegisterCodec(jsonrpc.NewCodec(), "application/json")
	rpcServer.RegisterService(bitcoin.NewService(uint(maxDaysDifference)), "bitcoin")
	router.Handle("/rpc", rpcServer)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), router))
}
