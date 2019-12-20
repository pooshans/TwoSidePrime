package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math"
	"net/http"
	"strconv"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/twoSidePrime/{num}", isTwoSidePrime)
	log.Fatal(http.ListenAndServe(":8082", myRouter))
}

func isTwoSidePrime(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	value := vars["num"]
	n, _ := strconv.Atoi(value) // To convert from String to Int
	json.NewEncoder(w).Encode(checkTwoSidePrime(n))
}

func main() {
	handleRequests()
}

func checkTwoSidePrime(value int) bool {
	if value == 0 || value == 1 {
		return false
	}
	var a int = value
	var tmp int = a
	var out bool = true
	var div int = 1
	for tmp > 0 {
		var rem int = tmp / (div * 10)
		if !(isPrime(tmp)) || (rem != 0 && !(isPrime(rem))) {
			out = false
			break
		}
		tmp = tmp / 10
	}

	return out
}

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			//fmt.Println("NO for ",  value)
			return false
		}
	}
	return value > 1
}
