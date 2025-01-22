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
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// fmt.Fprint(w, "Hello HttpRouter")
		panic("internal server error")
	})

	server := http.Server{
		Handler: router,
		Addr:    "localhost:3000",
	}

	fmt.Println("Server berjalan di http://localhost:3000")
	server.ListenAndServe()
}
