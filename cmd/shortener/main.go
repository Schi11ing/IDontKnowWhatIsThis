package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
		myRouter := mux.NewRouter()
		myRouter.HandleFunc("/", WriteURL).Methods("POST")
		myRouter.HandleFunc("/{id}",ReturnOriginUrl).Methods("GET")
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




func ReturnOriginUrl(w http.ResponseWriter, r *http.Request){
	if r != nil {
		vars := mux.Vars(r)
		key := vars["id"]
		origin, err := ReturnOrigin(key)
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
		shorted, err := RegisterUrl(string(reqBody))
		if err != nil {
			log.Println("err")
		}else {
			output := ("http://127.0.0.1:8080/" +shorted)
			w.WriteHeader(201)
			fmt.Fprintf(w, output)
		}
	}
}

var urlsMap = make(map[string]string)


func RegisterUrl(url string) (string, error){
	if _, found := urlsMap[url]; found {
		return "", errors.New("URL already registered")
	} else {
		val, err := createShortUrl(url)
		if err != nil {
			return "", errors.New("Returned empty string")
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
		return "",errors.New("No such URL")
	}
}

func createShortUrl(url string)(string, error){
	input := []byte(url)
	encoded := base64.StdEncoding.EncodeToString(input)
	if encoded != "" {
		return encoded, nil
	} else {
		return "", errors.New("Empty string")
	}
}



