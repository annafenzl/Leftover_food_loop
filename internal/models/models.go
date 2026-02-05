package models

import (
	"fmt"
	"time"
)

type Role string

const (
	Participant Role = "Participant"
	// the persons distributing the food, aka those people are mobile
	Volunteer Role = "Volunteer"
)

//go:generate go run golang.org/x/tools/cmd/stringer@latest -type=Status
type Status int

const (
	Available Status = iota
	Reserved
	Expired
)

type UserId int
type OfferId int
type RequestId int

// Preference is a bitmask describing whether a user is willing to
// pick up and/or deliver food.
type Preference uint8

const (
	PrefPickup Preference = 1 << iota
	PrefDelivery
)

func (p Preference) String() string {
	if p == 0 {
		return "None"
	}
	s := ""
	if p&PrefPickup != 0 {
		s += "Pickup"
	}
	if p&PrefDelivery != 0 {
		if s != "" {
			s += "|"
		}
		s += "Delivery"
	}
	return s
}

// what about a bitmap if the person would do
// [] pickup for food
// [] delivery
type User struct {
	ID                 UserId
	PhoneNumber        string
	FirstName          string
	Address            string
	FloorNumber        int
	PickupInstructions string
	Preferences        Preference
}

type Food struct {
	Food      string
	Portions  int
	CreatedAt time.Time
}

// how do i handle the feedback, ie we have a match, how do they know who picks up where
type Offer struct {
	ID     UserId
	Status Status
	Food
}

// “If food becomes available nearby, someone will bring it in a reasonable time
type Request struct {
	ID     UserId
	Time   time.Time
	Status Status
}

func (u User) String() string {
	return fmt.Sprintf("%s, id %d role: %s, prefs: %s", u.FirstName, u.ID, u.Preferences, u.Preferences)
}

func (o Offer) String() string {
	return fmt.Sprintf("OFFER ID %d %s, %d portions, status: %s, created at: %s", o.ID, o.Food.Food, o.Portions, o.Status, o.CreatedAt)
}

func (r Request) String() string {
	return fmt.Sprintf("REQUEST ID %d, time: %s, status: %s", r.ID, r.Time, r.Status)
}
