package store

import (
    "sync"
)

// Local placeholder types to avoid depending on an external module during build.
// Replace these with models.User / models.Offer / models.Request and the correct import
// once the module path / go.mod is fixed.
type User struct{}
type Offer struct{}
type Request struct{}

type Store struct {
	Users    map[string]*User
	Offers   map[string]*Offer
	Requests map[string]*Request
	mu       sync.Mutex
}

var store = &Store{
	Users:    make(map[string]*User),
	Offers:   make(map[string]*Offer),
	Requests: make(map[string]*Request),
}