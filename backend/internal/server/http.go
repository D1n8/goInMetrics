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
	http.HandleFunc("/ws", wsHandler)
	http.HandleFunc("/", homePage)
}

func HttpLaunch() {
	cfg, err := config.ConfigLoader()
	if err != nil {
		log.Fatal(err)
	}

	Port :=  cfg.Server.Port

	log.Println("HTTP up and running...")
	log.Printf("%s:%d", "http://localhost", Port)
	setupRoutes()
	http.ListenAndServe(fmt.Sprintf(":%d", Port), nil)
}
