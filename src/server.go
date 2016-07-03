package main
import (	
	"io"
	"net/http"
	"log"
	"fmt"
	. "pp"
)


func PopCornTimeRequestHandler(res http.ResponseWriter, req *http.Request) {
	log.Print(req.RequestURI)
	io.WriteString(res, "hello, world!\n")
}



func main() {
	http.HandleFunc("/api/v2/", PopCornTimeRequestHandler)
	StartTracker("rutracker")
	var port int = GetConfiguration().Port
	log.Print(fmt.Sprintf("Starting server on %d", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}