package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func IndexHundler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Go Web API Server.")
}

func CalcFibHundler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	num, _ := strconv.Atoi(ps.ByName("num"))
	var fibs Fibs

	for i := 0; i < num; i++ {
		fib := fibonacci(i)
		fibs = append(fibs, Fib{
			Input:  i,
			Output: fib,
		})
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(fibs); err != nil {
		panic(err)
	}
}
