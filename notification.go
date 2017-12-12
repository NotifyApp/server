package main

import (
	"bytes"
	"encoding/json"
)

// Notification is the notifs channel
type Notification struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

// ToBytes convert struct to byte array
func (n Notification) ToBytes() []byte {
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(n)
	return buffer.Bytes()
}

func (n Notification) String() string {
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(n)
	return string(buffer.Bytes())
}
