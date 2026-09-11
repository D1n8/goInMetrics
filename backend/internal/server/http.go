package server

import (
	"backend/internal/config"
	"fmt"
	"log"
	"net/http"
)


func homePage(w http.ResponseWriter, r *http.Request) {
	//http.ServeFile(w, r, "index.html")
	log.Println(r)
	fmt.Fprintf(w, "Hello World")
}
func setupRoutes() {
	http.HandleFunc("/", homePage)
}

func HttpLaunch() {
	cfg, err := config.ConfigLoader()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("HTTP up and running...")
	setupRoutes()
	http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), nil)
}
