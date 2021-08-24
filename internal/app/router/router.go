package router
import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"IDontKnowWhatIsThis/internal/app/hasher"
)
func HandleRequests(w http.ResponseWriter, r *http.Request) {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", writeURL).Methods("POST")
	myRouter.HandleFunc("/{id}",returnOriginUrl).Methods("GET")
}

func returnOriginUrl(w http.ResponseWriter, r *http.Request){
	if r == nil {
		vars := mux.Vars(r)
		key := vars["id"]
	}
	hasher.
}

func writeURL(w http.ResponseWriter, r *http.Request){
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("err")
	}
}
