package main

import "IDontKnowWhatIsThis/internal/app/server"

func main() {
	var protocol = "http"
	var port = "8080"
	server.Server(protocol,port)
}
