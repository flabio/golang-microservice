package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"test/service"

	"github.com/gorilla/mux"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customers, err := ch.service.GetCustomer(id)

	if err != nil {
		w.WriteHeader(err.Code)
		writeResponse(w, err.Code, err.AsMessage())
		return
	} else {
		writeResponse(w, http.StatusOK, customers)
	}

}
func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}

}
