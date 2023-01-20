package main

import (
	"fmt"
	"net/http"

	"github.com/Doomann/gohome/controllers"
)

const (
	apiPort = 8888
	number  = iota
)

func main() {
	startWebServer(apiPort)
}

func startWebServer(port int) (int, error) {
	fmt.Println("Starting server...")
	controllers.RegisterControllers()

	fmt.Println("Server is up n runnin' on port", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	return port, nil
}
