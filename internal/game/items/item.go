package item

import (
	"project_m/assets"
	"project_m/internal/engine"
	sprite "project_m/internal/engine"
)

type Item struct {
	ID 	string
	Name   string
	Sprite *sprite.Sprite
}

func (i *Item) Instantiate() *Item {
	item := *i

	if item.Sprite != nil {
		item.Sprite = item.Sprite.Instantiate()
	}

	return &item
}

func (i *Item) GetSprite() *sprite.Sprite {
	if i.Sprite == nil {
		return nil
	}
	return i.Sprite
}

func (i *Item) GetName() string {
	return i.Name
}

func CreateItem(id string, name string, spritePath string) *Item {
	Image := assets.LoadImage(spritePath)

	return &Item{
		ID:   id,
		Name: name,
		Sprite: engine.NewSprite(
			id,
			name,
			Image,
		),
	}
}

type ItemManager struct {
	Cobre *Item
}

func NewItemManager() *ItemManager {
	return &ItemManager{
		Cobre: CreateItem("cobre", "Cobre", "ores/cobre.png"),
	}
}