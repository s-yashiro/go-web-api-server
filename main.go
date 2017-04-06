package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/newrelic/go-agent"
	"github.com/newrelic/go-agent/_integrations/nrlogrus"
	"log"
	"net/http"
)

func main() {
	config := newrelic.NewConfig("go-web-api-server", "2aa5ce866eb47508f2d67070e63b18fe19012b17")
	app, err := newrelic.NewApplication(config)
	if err != nil {
		config.Logger = nrlogrus.StandardLogger()
	}

	router := httprouter.New()
	router.GET("/", Logging(IndexHundler, "index"))
	router.GET("/fib/:num", Logging(CalcFibHundler, "calc-fib"))
	log.Fatal(http.ListenAndServe(":80", router))

	txn := app.StartTransaction("main", nil, nil)
	defer txn.End()
}
