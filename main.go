package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	log.Fatal(http.ListenAndServe(":3000", router))
	router.GET("/", Logging(IndexHundler, "index"))
	router.GET("/fib/:num", Logging(CalcFibHundler, "calc-fib"))
}
