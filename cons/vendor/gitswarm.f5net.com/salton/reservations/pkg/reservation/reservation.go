package reservation

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

const (
	OPNEW string = "New_Reservation"
	OPREM string = "Remove_Reservation"
	OPCHG string = "Update_Reservation"
)

// Reservation is the main object of the system
type Reservation struct {
	ID    int    `json:"id,omitempty"`
	Date  string `json:"date"`
	Name  string `json:"name"`
	Hour  int    `json:"hour"`
	Party int    `json:"party"`
}

// Bytes converts a Reservation instance into an array of bytes.
func (r Reservation) Bytes() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(r); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Print prints a reservation details
func (r Reservation) Print() {
	fmt.Println("RESERVATION : ")
	fmt.Printf(" Name : %s\n", r.Name)
	fmt.Printf(" Party : %d\n", r.Party)
	fmt.Printf(" Date : %s\n", r.Date)
	fmt.Printf(" Hour : %d\n", r.Hour)
}

// Decode convert array of bytes into a new instance of a reservation.
func Decode(b []byte) (Reservation, error) {
	var r Reservation
	dec := gob.NewDecoder(bytes.NewReader(b))
	err := dec.Decode(&r)
	return r, err
}
