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


type User struct {
	ID                   UserId
	PhoneNumber          string
	FirstName            string
	Role                 Role
	Address              string
	FloorNumber          int
	PickupInstructions   string
}

type Food struct {
	Food      string
	Portions  int
	ExpiresAt time.Time
}

// how do i handle the feedback, ie we have a match, how do they know who picks up where
type Offer struct {
	ID        UserId
	Status    Status
	Food
}

// “If food becomes available nearby, someone will bring it in a reasonable time
type Request struct {
	ID        UserId
	Time      time.Time
	Status    Status
}

func (u User)String() string {
	return fmt.Sprintf("%s, id %d role: %s ", u.FirstName, u.ID, u.Role)
}

func (o Offer) String() string {
	return fmt.Sprintf("OFFER ID %s %s, %d portions, status: %s, expires at: %s", o.ID, o.Food, o.Portions, o.Status, o.ExpiresAt)
}

func (r Request) String() string {
	return fmt.Sprintf("REQUEST ID %s, date %d %s %d, status: %s", r.ID, r.date.Year, r.date.Month, r.date.day, r.Status)
}