package main

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"os"
	"strconv"
)

var port = os.Getenv("PORT")

type server struct {
	router *chi.Mux
}

func main() {
	if port == "" {
		port = "5000"
	}

	if err := checkPort(port); err != nil {
		log.Fatalln(err)
	}

	r := chi.NewRouter()
	s := server{router: r}

	address := fmt.Sprintf(":%s", port)
	log.Printf("listening on %s\n", address)
	if err := http.ListenAndServe(address, s.router); err != nil {
		log.Fatalln(err)
	}
}

func checkPort(port string) error {
	p, err := strconv.Atoi(port)
	if err != nil {
		return err
	}

	if p < 0 || p > 65535 {
		return errors.New("PORT is out of acceptable range [0..65535]")
	}

	return nil
}
