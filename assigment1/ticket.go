package main

import "time"

type Ticket struct {
	price     float64
	createdAt time.Time
	car       *Car
}

func NewTicket(car *Car, price float64, parking *ParkSlot) *Ticket {
	t := Ticket{createdAt: time.Now(), car: car, price: price}
	return &t
}

func (t *Ticket) Pay(hours float64) float64 {
	if hours >= 2 {
		t.price = 10 + (hours-2)*10
	} else {
		t.price = 5 * hours
	}

	return t.price
}
