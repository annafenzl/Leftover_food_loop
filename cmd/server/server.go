package server

import (
	"annafenzl/leftoverfoodloop/internal/models"
	"annafenzl/leftoverfoodloop/internal/store"
	"fmt"
	"log"
	"net/http"
	"strconv"
)


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

	// put it into db
	

	// check up if somebody wants to have some food
	// match them, notify them, forward info,

	// Set Status

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

//TODO: error handling
func Add_User(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return 
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	floor_number, err := strconv.Atoi(r.FormValue("floor_number"))
	if err != nil {
		return
	}
	instruction := r.FormValue("instructionss")
	phone_number := r.FormValue("phone_number")

	store.AddUser(name, address, floor_number, instruction, phone_number)

}

func Start() {

	fileServer := http.FileServer(http.Dir("./static")) 

    http.Handle("/", fileServer) 

	// Func to Add User
	http.HandleFunc("/addUser", Add_User)
	
	// Func to Post Food
	http.HandleFunc("/request", RequestFood)

	// Func to offer
	http.HandleFunc("/offer", OfferFood)


	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}