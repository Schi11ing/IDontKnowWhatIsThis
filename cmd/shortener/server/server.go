package server

import(
	"IDontKnowWhatIsThis/cmd/shortener/router"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"time"
)
func Server(){
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/", router.WriteURL).Methods("POST")
	myRouter.HandleFunc("/{id}",router.ReturnOriginUrl).Methods("GET")
	listener, err := net.Listen("tcp","127.0.0.1:8080")
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