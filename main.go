package main

import (
	"log"
	"net/http"

	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json"
	"ufirst.com/bitcoin/bitcoin"
)

func main() {
	router := http.NewServeMux()
	rpcServer := rpc.NewServer()
	router.Handle("/rpc", rpcServer)
	log.Fatal(http.ListenAndServe(":8080", router))
}
