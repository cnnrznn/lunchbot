package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/event", EventHandler)
	http.HandleFunc("/imback", ImbackHandler)
	http.HandleFunc("/authorize", AuthorizeHandler)
	http.HandleFunc("/test", TestHandler)
	log.Fatal(http.ListenAndServeTLS(":443", os.Getenv("PEMFILE"), os.Getenv("KEYFILE"), nil))
}

func TestHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Response from /test"))
}
