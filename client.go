package client

import "github.com/shopspring/decimal"

type Position struct {
	Instrument string
	Quantity   decimal.Decimal
	CostBasis  decimal.Decimal
}

type GTClient interface {
	CurrentPositions() map[Position]bool
}
