package inventory

import (
	sprite "project_m/internal/engine"
)

type itemInterface interface {
	GetSprite() *sprite.Sprite
	GetName() string
}

type InventorySlot struct {
	Quantity int
	Item     itemInterface
}

type Inventory struct {
	Items [10]InventorySlot
}

func (i *Inventory) AddItem(item itemInterface, quantity int) {
	for index, inventoryItem := range i.Items {
		switch inventoryItem.Item {
		case nil:
			i.Items[index] = InventorySlot{
				Quantity: quantity,
				Item:     item,
			}
			return
		case item:
			i.Items[index].Quantity += quantity
			return
		}
	}
}

func (i *Inventory) RemoveItem(item itemInterface, quantity int) {
	for index, inventoryItem := range i.Items {
		if inventoryItem.Item == item {
			if inventoryItem.Quantity > quantity {
				i.Items[index].Quantity -= quantity
			} else {
				i.Items[index] = InventorySlot{}
			}
			return
		}
	}
}