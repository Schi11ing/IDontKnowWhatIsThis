package server

import(
	"IDontKnowWhatIsThis/internal/app/router"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"time"
)
func Server(protocol, port string){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", router.WriteURL).Methods("POST")
	myRouter.HandleFunc("/{id}",router.ReturnOriginUrl).Methods("GET")
	listener, err := net.Listen(protocol,port)
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Handler: myRouter,
		WriteTimeout: 15* time.Second,
		ReadTimeout: 15*time.Second,
	}
	err = server.Serve(listener)
	if err != nil{
		panic(err)
	}
}