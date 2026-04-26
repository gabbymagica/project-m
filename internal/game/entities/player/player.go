package player

import (
	"project_m/internal/engine"
	"project_m/internal/inventory"
)

type Player struct {
	engine.Entity
	Inventory *inventory.Inventory
}
