package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	app "loan-calculator/application"

	"github.com/gorilla/mux"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	io.WriteString(w, "Hello "+vars["name"])
}

/*
{
	"amount": 500,
	"rate": 500,
	"requestTime": 1597640186380,
	"repayDay": 17,
	"termNum": 6,
	"interestType": "EQUAL_LOAN"
}
*/
func TrialPayment(w http.ResponseWriter, req *http.Request) {
	reqBody, _ := ioutil.ReadAll(req.Body)
	var request *app.PaymentTrailRequest = new(app.PaymentTrailRequest)
	fmt.Println(reqBody)
	json.Unmarshal(reqBody, request)
	response := app.TrailPayment(request)
	json.NewEncoder(w).Encode(response)
}
func corsHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			//handle preflight in here
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			h.ServeHTTP(w, r)
		}
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello/{name}", Hello)
	r.Handle("/loan/trail", corsHandler(http.HandlerFunc(TrialPayment)))
	// http.Handle("/", r)
	http.ListenAndServe(":8000", r)
}
