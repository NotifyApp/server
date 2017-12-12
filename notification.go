package main

import (
	"encoding/json"
	"log"
)

// Notification is the notifs channel
type Notification struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

// ToBytes convert struct to byte array
func (n Notification) ToBytes() []byte {
	b, err := json.Marshal(&n)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func (n Notification) String() string {
	b, err := json.Marshal(&n)
	if err != nil {
		log.Fatal(err)
	}
	return string(b[:])
}
