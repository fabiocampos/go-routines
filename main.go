package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type FibonacciResponse struct {
	Request int `json:"number"`
	Value   int `json:"value"`
}

func main() {
	http.HandleFunc("/fibonacci/", FibonacciHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
	requestedValue, conversionError := strconv.Atoi(r.URL.Path[len("/fibonacci/"):])
	if conversionError != nil {
		requestedValue = 0
	}

	fibonacci := CalcFibonacci(requestedValue)
	fibonacciResponse := &FibonacciResponse{Request: requestedValue, Value: fibonacci}
	response, error := json.Marshal(fibonacciResponse)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func CalcFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return CalcFibonacci(n-1) + CalcFibonacci(n-2)
}
