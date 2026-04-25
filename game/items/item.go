package item

import (
	"project_m/assets"
	sprite "project_m/engine"
)


type Item struct {
	ID 	string
	Name   string
	Sprite *sprite.Sprite
}

func InstanceItem(id string, name string, spritePath string) *Item {
	Image := assets.LoadImage("items/" + spritePath)

	return &Item{
		ID:   id,
		Name: name,
		Sprite: &sprite.Sprite{
			ID: id,
			Name: name,
			Image: Image,
		},
	}
}

type ItemManager struct {
	Cobre *Item
}

func NewItemManager() *ItemManager {
	return &ItemManager{
		Cobre: InstanceItem("cobre", "Cobre", "cobre.png"),
	}
}