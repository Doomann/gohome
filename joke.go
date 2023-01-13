package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Joke struct {
	Value string `json:"value"`
}


func getJoke() string {
	response, _ := http.Get("https://api.chucknorris.io/jokes/random")
	
	defer response.Body.Close()
	
	body, _ := ioutil.ReadAll(response.Body)
	
	joke := Joke{}
	json.Unmarshal(body, &joke)

	return joke.Value

}