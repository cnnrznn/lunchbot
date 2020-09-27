package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/event", EventHandler)
	http.HandleFunc("/authorize", AuthorizeHandler)
	log.Fatal(http.ListenAndServeTLS(":443", os.Getenv("PEMFILE"), os.Getenv("KEYFILE"), nil))
}
