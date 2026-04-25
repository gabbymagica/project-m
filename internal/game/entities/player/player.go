package player

import (
	"project_m/internal/game/entities"
	"project_m/internal/inventory"
)

type Player struct {
	entities.Entity
	Inventory *inventory.Inventory
}
