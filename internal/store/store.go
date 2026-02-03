package store
import (
	"annafenzl/leftoverfoodloop/internal/models"
	"sync"
)

type Store struct {
	// TODO: check for existing users, functions to look up phone numbers...
	Usercount	models.UserId
    Users    map[models.UserId]*models.User
    Offers   map[string]*models.Offer
    Requests map[string]*models.Request
    mu       sync.Mutex
}

var store = &Store{
    Users:    make(map[models.UserId]*models.User),
    Offers:   make(map[string]*models.Offer),
    Requests: make(map[string]*models.Request),
}


func AddUser(name string, address string, floor_number int, instruction string, phone_number string) {
	store.mu.Lock()
	defer store.mu.Unlock()

	if _, exists := store.Users[store.Usercount]; exists {
		return 
	}

	store.Users[store.Usercount] = &models.User{
		ID: store.Usercount,
		PhoneNumber: phone_number,
		FirstName: name,
		Role: "Participant",
		Address: address,
		FloorNumber: floor_number,
		PickupInstructions: instruction,
	}

	store.Usercount++
}

func AddOffer(UserId models.UserId,  ) {
	store.mu.Lock()
	defer store.mu.Unlock()


	store.Offers[]


	store.Usercount++
}


func (s *Store)Match() {

}