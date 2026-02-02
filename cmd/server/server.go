package server

import (
	"annafenzl/leftoverfoodloop/internal/models"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Store struct {
    Users    map[string]*models.User
    Offers   map[string]*models.Offer
    Requests map[string]*models.Request
    mu       sync.Mutex
}

var store = &Store{
    Users:    make(map[string]*models.User),
    Offers:   make(map[string]*models.Offer),
    Requests: make(map[string]*models.Request),
}

// if it over the phone a loved one can register once, and then the phone number is tied to an address
func OfferFood(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	food_type := r.FormValue("food")
	portions := r.FormValue("portions")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
	fmt.Fprintf(w, " %s portions of %s\n", portions, food_type)
}

// same here, if the user is an elder you can put the address equivalent to the phone number
func RequestFood(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/request" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Println("In the Get Handler")
	fmt.Fprintf(w, "Hello, you might get something!")


	// if there is a match, notify the person
}


func Start() {

	fileServer := http.FileServer(http.Dir("./static")) 

    http.Handle("/", fileServer) 
	
	// Func to Post Food
	http.HandleFunc("/request", RequestFood)

	// Func to offer
	http.HandleFunc("/offer", OfferFood)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}