package main

import "IDontKnowWhatIsThis/internal/app/server"

func main() {
	var protocol = "http"
	var port = "800"
	srv := server.Server(protocol,port)
	go srv()

}
