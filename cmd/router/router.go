package router
import (
	hasher "IDontKnowWhatIsThis/cmd/hasher"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func ReturnOriginUrl(w http.ResponseWriter, r *http.Request){
	if r != nil {
		vars := mux.Vars(r)
		key := vars["id"]
		origin, err := hasher.ReturnOrigin(key)
		defer r.Body.Close()
		if err != nil{
			log.Println("Error in finding string")
		} else {
			w.Header().Set("Location", origin)
			w.WriteHeader(307)
		}
	} else {
		w.WriteHeader(400)
	}

}

func WriteURL(w http.ResponseWriter, r *http.Request){
	reqBody, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println("err")
	} else {
		shorted, err := hasher.RegisterUrl(string(reqBody))
		if err != nil {
			log.Println("err")
		}else {
			output := ("http://127.0.0.1:8080/" +shorted)
			fmt.Fprintf(w, output)
		}
	}
}
