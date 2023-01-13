package main

import (
	"fmt"
	"net/http"
)

const (
	apiPort = ":8888"
	test = iota
	ziurek = "krabas"
	kas = iota
	cia = iota
	bus = iota
)

type bandomLaime struct {
	sitasNumeris int
	nuCiaString  string
}

func main() {
	
	http.HandleFunc("/joke", func(w http.ResponseWriter, r *http.Request) {
		joke := getJoke()
		fmt.Fprint(w, joke)	
	})
		
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		grazinsimSita := bandomLaime{cia, ziurek}
		stringas := fmt.Sprintf("kodel iota visada buna kitoks? %d", bus)

		fmt.Println(stringas)
		fmt.Fprint(w, grazinsimSita)	
	})

	fmt.Println("Servakas up n runnin...")
	http.ListenAndServe(apiPort, nil)

}