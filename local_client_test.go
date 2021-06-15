package client

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestLocalClient_OpenPositions(t *testing.T) {
	lc := NewLocalClient()

	expectedPosition := GTPosition{"GME", decimal.NewFromInt(100), decimal.NewFromInt(1)}
	expectedGTPositionMap := make(map[GTPosition]bool)
	expectedGTPositionMap[expectedPosition] = true

	order := GTOrder{"GME", decimal.NewFromInt(100), MKT}
	lc.SubmitOrder(&order)

	for pos := range lc.OpenPositions() {
		if inst := pos.Instrument; inst != "GME" {
			t.Errorf("LocalClient.OptionPositions() instrument: %v, want GME", inst)
		}
	}
}
