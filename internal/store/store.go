package store

import (
	"annafenzl/leftoverfoodloop/internal/models"
	"sync"
	"time"
)

type Store struct {

	nextUserId    models.UserId
	nextOfferId   models.OfferId
	nextRequestId models.RequestId

	Users    map[models.UserId]*models.User
	Offers   map[models.OfferId]*models.Offer
	Requests map[models.RequestId]*models.Request
	mu       sync.Mutex
}

var store = &Store{
	Users:    make(map[models.UserId]*models.User),
	Offers:   make(map[models.OfferId]*models.Offer),
	Requests: make(map[models.RequestId]*models.Request),
}


func AddParticipant(name string, address string, floor_number int, instruction string, phone_number string) {
	store.mu.Lock()
	defer store.mu.Unlock()

	if _, exists := store.Users[store.nextUserId]; exists {
		return 
	}

	// TODO: add preferences
	store.Users[store.nextUserId] = &models.User{
		ID: store.nextUserId,
		PhoneNumber: phone_number,
		FirstName: name,
		Address: address,
		FloorNumber: floor_number,
		PickupInstructions: instruction,
	}
	store.nextUserId++
}

func AddVolunteer(name string, address string, floor_number int, instruction string, phone_number string) {
	store.mu.Lock()
	defer store.mu.Unlock()

	if _, exists := store.Users[store.nextUserId]; exists {
		return 
	}

	store.Users[store.nextUserId] = &models.User{
		ID: store.nextUserId,
		PhoneNumber: phone_number,
		FirstName: name,
		Address: address,
		FloorNumber: floor_number,
		PickupInstructions: instruction,
	}
	store.nextUserId++
}


func AddOffer(UserId models.UserId, food models.Food) {
	store.mu.Lock()
	defer store.mu.Unlock()

	store.Offers[store.nextOfferId] = &models.Offer{
		ID: UserId,
		Status: models.Available,
		Food: food,
	}

	store.nextOfferId++
}

func AddRequest(UserId models.UserId) {

	store.mu.Lock()
	defer store.mu.Unlock()

	store.Requests[store.nextRequestId] = &models.Request{
		ID: UserId,
		Time: time.Now(),
		Status: models.Available,
	}

	store.nextRequestId++
}

func (s *Store)Match() {

}