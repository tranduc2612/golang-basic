package main

type Bill struct {
	cost float64
}

func newBill(cost float64) *Bill {
	b := Bill{cost: cost}
	return &b
}

func (b *Bill) updateBill(cost float64) {
	b.cost = cost
}
