package main

import "strconv"

type ParkSlot struct {
	id             int
	total_money    float64
	current_ticket *Ticket
}

func NewParking(id int, total_money float64) *ParkSlot {
	ps := ParkSlot{id: id, total_money: total_money, current_ticket: nil}
	return &ps
}

func (p *ParkSlot) AddTicket(current_ticket *Ticket) {
	p.current_ticket = current_ticket
}

func (p *ParkSlot) RemoveTicket() {
	p.current_ticket = nil
}

func (p *ParkSlot) SumTotal() {
	p.total_money += p.current_ticket.price
}

func (p *ParkSlot) GetStatus() string {
	var statusParking string
	if p.current_ticket == nil {
		statusParking = "Available"
	} else {
		statusParking = "Disavailable"
	}

	return statusParking
}

func (p *ParkSlot) GetMoney() string {
	convertTotalMoney := strconv.FormatFloat(p.total_money, 'f', 2, 64)
	return convertTotalMoney
}

func (p *ParkSlot) GetId() string {
	convertId := strconv.Itoa(p.id)
	return convertId
}

func (p *ParkSlot) RenderInfo() string {
	convertId := strconv.Itoa(p.id)
	convertTotalMoney := strconv.FormatFloat(p.total_money, 'f', 2, 64)
	var statusParking string

	if p.current_ticket == nil {
		statusParking = "Available"
	} else {
		statusParking = "Disavailable"
	}

	return "Id:" + convertId + " Total_money:" + convertTotalMoney + " Status:" + statusParking
}
