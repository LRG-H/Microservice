package main

import (
	"log"
	"net/http"
	"os"

	"github.com/LRG-H/Microservice/tree/main/Handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := Handlers.NewHello()
	gh := Handlers.NewGoodbye()

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	http.ListenAndServe(":9090", sm)
}
