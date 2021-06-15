package client

import "github.com/shopspring/decimal"

type LocalClient struct {
	positions map[GTPosition]bool
	orders    map[GTOrder]bool
}

func NewLocalClient() LocalClient {
	positions := make(map[GTPosition]bool)
	orders := make(map[GTOrder]bool)
	return LocalClient{positions, orders}
}

func (lc *LocalClient) OpenPositions() map[GTPosition]bool {
	return lc.positions
}

func (lc *LocalClient) OpenOrders() map[GTOrder]bool {
	return lc.orders
}

func (lc *LocalClient) SubmitOrder(order *GTOrder) error {
	pos := GTPosition{order.Instrument, order.Quantity, decimal.NewFromInt(1)}
	lc.positions[pos] = true
	return nil
}
