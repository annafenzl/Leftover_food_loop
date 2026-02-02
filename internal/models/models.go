package models

import (
	"time"
)

type Role string
const (
	Participant Role = "Participant"
	// the persons distributing the food, aka those people are mobile
	Volunteer Role = "Volunteer"
)

type Status int
const (
	available Status = iota
	reserved
	expired
)

type UserId string

type User struct {
	ID					UserId
	FirstName			string

	Role				Role
	// pre registered, only people with Role Participiant need this
	Phone_number		string
	Address				string
	Floor_number		int
	Pickup_instructions	string

	// do i need this circle id??
	CircleID string
}

// how do i handle the feedback, ie we have a match, how do they know who picks up where
type Offer struct {
	ID        UserId
	Food      string
	Portions  int
	ExpiresAt time.Time
	CircleID  string
	Status    Status
}

type Date struct {
	Year	int
	Month	time.Month
	day		int
}

// “If food becomes available nearby, someone will bring it in a reasonable time
type Request struct {
	ID        UserId
	date      Date
}
