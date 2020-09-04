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

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello/{name}", Hello)
	r.HandleFunc("/loan/trail", TrialPayment)
	// http.Handle("/", r)
	http.ListenAndServe(":8000", r)
}
