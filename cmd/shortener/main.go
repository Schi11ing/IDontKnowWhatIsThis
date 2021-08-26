package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)
// IN main function exists router with two handlers for get and post methods
// also here is created listner with parameters and server started
func main() {
		myRouter := mux.NewRouter()
		myRouter.HandleFunc("/", WriteURL).Methods("POST")
		myRouter.HandleFunc("/{id}",ReturnOriginURL).Methods("GET")
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



//IN ReturnOriginURL we search wia id in map and if we found
func ReturnOriginURL(w http.ResponseWriter, r *http.Request){
	if r != nil {
		vars := mux.Vars(r)
		key := vars["id"]
		origin, err := ReturnOrigin(key)
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Println("return originULR body close error")
			}
		}(r.Body)
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
	if err != nil {
		log.Println("err")
	} else {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Println("writeURL body close error")
			}
		}(r.Body)
		shorted, err := RegisterURL(string(reqBody))
		if err != nil {
			log.Println("err")
		}else {
			output := "http://127.0.0.1:8080/" +shorted
			w.WriteHeader(201)
			_, err2 := fmt.Fprint(w, output)
			if err2 != nil {
				return 
			}
		}
	}
}

var urlsMap = make(map[string]string)


func RegisterURL(url string) (string, error){
	if _, found := urlsMap[url]; found {
		return "", errors.New("url already registered")
	} else {
		val, err := createShortURL(url)
		if err != nil {
			return "", errors.New("returned empty string")
		} else {
			urlsMap[val] = url
			return val, nil
		}
	}
}

func ReturnOrigin(url string) (string, error){
	if val, found := urlsMap[url]; found{
		return val, nil
	} else{
		return "",errors.New("no such URL")
	}
}

func createShortURL(url string)(string, error){
	input := []byte(url)
	encoded := base64.StdEncoding.EncodeToString(input)
	if encoded != "" {
		return encoded, nil
	} else {
		return "", errors.New("empty string")
	}
}



