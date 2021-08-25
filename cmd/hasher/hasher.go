package hasher

import (
	"encoding/base64"
	"errors"
)

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

