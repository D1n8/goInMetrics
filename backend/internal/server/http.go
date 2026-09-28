package server

import (
	"backend/internal/config"
	"backend/internal/pcstat"
	"fmt"
	"log"
	"net/http"
)


func homePage(w http.ResponseWriter, r *http.Request) {
	//http.ServeFile(w, r, "index.html")
	log.Println(r)
	// FIX: REMOVE JsonData
	JsonData, err := pcstat.SendDataToServer()
	if err != nil {
		return
	}
	fmt.Fprintf(w, string(JsonData))
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
	// FIX: Availability Port check
	http.ListenAndServe(fmt.Sprintf(":%d", Port), nil)
}
