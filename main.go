package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Vendor struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Address  string   `json:"address"`
	Products *Product `json:"product"`
}

type Product struct {
	Sku         string `json:"sku"`
	Description string `json:"description"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

var vendors []Vendor

func getVendors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vendors)
}

func deleteVendor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range vendors {
		if item.Id == params["id"] {
			vendors = append(vendors[:index], vendors[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(vendors)
}

func getVendor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range vendors {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createVendor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var vendor Vendor
	_ = json.NewDecoder(r.Body).Decode(&vendor)
	vendor.Id = strconv.Itoa(rand.Intn(1000000))
	vendors = append(vendors, vendor)
	json.NewEncoder(w).Encode(vendor)

}

func updateVendor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range vendors {
		if item.Id == params["id"] {
			vendors = append(vendors[:index], vendors[index+1:]...)
			var vendor Vendor
			_ = json.NewDecoder(r.Body).Decode(&vendor)
			vendor.Id = params["id"]
			vendors = append(vendors, vendor)
			json.NewEncoder(w).Encode(vendor)
			return
		}
	}

}

func main() {
	r := mux.NewRouter()

	vendors = append(vendors, Vendor{Id: "1", Name: "Supply Corp", Address: "100 Gulf st.", Products: &Product{Sku: "1072567", Description: "Firewood .75 Cord"}})
	vendors = append(vendors, Vendor{Id: "2", Name: "Manchester Systems", Address: "1146 Tyree st.", Products: &Product{Sku: "1072567", Description: "Firewood .75 Cord"}})
	vendors = append(vendors, Vendor{Id: "3", Name: "Orbits Inc", Address: "700 Springer dr.", Products: &Product{Sku: "1072567", Description: "Firewood .75 Cord"}})

	r.HandleFunc("/vendors", getVendors).Methods("GET")
	r.HandleFunc("/vendors/{id}", getVendor).Methods("GET")
	r.HandleFunc("/vendors", createVendor).Methods("POST")
	r.HandleFunc("/vendors/{id}", updateVendor).Methods("PUT")
	r.HandleFunc("/vendors/{id}", deleteVendor).Methods("DELETE")

	fmt.Printf("Starting Server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
