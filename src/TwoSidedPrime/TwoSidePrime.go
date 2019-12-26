package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math"
	"net/http"
	"strconv"
)

/*
Web request handler
*/
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/twoSidePrime/{num}", isTwoSidePrime)
	log.Fatal(http.ListenAndServe(":8082", myRouter))
}

/*
Accept request and return response.
*/
func isTwoSidePrime(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	value := vars["num"]
	// To convert from String to Int
	n, _ := strconv.Atoi(value)
	json.NewEncoder(w).Encode(checkTwoSidePrime(n))
}

func main() {
	handleRequests()
}

func checkTwoSidePrime(value int) bool {
	if value == 0 || value == 1 {
		return false
	}
	var tmp int = value
	var div int = 1
	for tmp > 0 {
		div = div * 10
		// "remainder" will be assigned with value of left side truncation of "value".
		var remainder int = value % div

		// If either of left side or right side truncated value is not prime, returning false.
		if !(tmp != 0 && isPrime(tmp)) || (remainder != 0 && !(isPrime(remainder))) {
			return false
		}
		// "tmp" will be assigned with value of right side truncation of "value".
		tmp = tmp / 10
	}

	// Else value is prime, returning true.
	return true
}

/*
Check whether value is prime or not.
*/
func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
