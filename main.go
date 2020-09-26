package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/event", EventHandler)
	http.HandleFunc("/authorize", AuthorizeHandler)
	log.Fatal(http.ListenAndServeTLS(":443", "server.crt", "server.key", nil))
}
