package server

import(
	"IDontKnowWhatIsThis/internal/app/router"
	"net"
	"net/http"
)
func Server(protocol, port string){
	Handler1 := router.route(r &http.Request)
	listener, err := net.Listen(protocol,port)
	if err != nil {
		panic(err)
	}
	server := &http.Server{

	}
	err = server.Serve(listener)
	if err != nil{
		panic(err)
	}
}