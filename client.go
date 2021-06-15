package client

import "github.com/shopspring/decimal"

type GTPosition struct {
	Instrument string
	Quantity   decimal.Decimal
	CostBasis  decimal.Decimal
}

type GTOrderType int

const (
	MKT GTOrderType = iota
	LMT
)

type GTOrder struct {
	Instrument string
	Quantity   decimal.Decimal
	OrderType  GTOrderType
}

type GTClient interface {
	OpenPositions() map[GTPosition]bool
	OpenOrders() map[GTOrder]bool
	SubmitOrder(order *GTOrder) error
}

func Buy(client GTClient, instrument string, qty decimal.Decimal) error {
	o := GTOrder{instrument, qty, MKT}
	return client.SubmitOrder(&o)
}
