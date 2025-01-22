package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		fmt.Fprint(w, "Error : ", i)
	}

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Oops! The page you are looking for was not found")
	})

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello HttpRouter")
	})

	router.GET("/panic", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		panic("internal server error")
	})

	server := http.Server{
		Handler: router,
		Addr:    "localhost:3000",
	}

	fmt.Println("Server berjalan di http://localhost:3000")
	server.ListenAndServe()
}
