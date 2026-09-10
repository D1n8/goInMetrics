package server

import (
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
func setupRoutes() {
	http.HandleFunc("/", homePage)
}

func httpLaunch() {
	log.Println("HTTP up and running...")
	log.Fatal()
}
