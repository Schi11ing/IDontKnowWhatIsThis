package hasher

import "errors"

var urlsMap = make(map[string]string)

func registerUrl(url string) (string, error){
	if val, ok := urlsMap[url];ok{
		return "", errors.New("URL already registered")
	}
	else {
		 urlMap[val] = urls
	}
}

func returnOrigin(url string) (string, error){
	if val, ok := urlsMap[url];ok{
		return val
	}
	else {
		return "",errors.New("No such URL")
	}
}

func createShortUrl(url string)(string, error){

}