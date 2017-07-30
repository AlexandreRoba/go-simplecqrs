package main

import (
	eh "github.com/looplab/eventhorizon"
)

type InventoryItem struct {
	*eh.AggregateBase
	
}
