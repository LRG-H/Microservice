package main

import (
	"log"
	"net/http"
	"os"

	"github.com/LRG-H/Microservice/tree/main/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello()
	gh := handlers.NewGoodbye()

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	http.ListenAndServe(":9090", sm)
}
