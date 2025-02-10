package main

import (
	api "localhost/pkg/API"
	"log"
	"net/http"
)

func main() {
    api := api.New("localhost:8090", http.NewServeMux())
    api.FillEndPoint()
    log.Fatal(api.ListenAndServe())

}



