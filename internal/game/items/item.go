package item

import (
	"project_m/assets"
	"project_m/internal/constants"
	"project_m/internal/engine"
)

type Item struct {
	ID     int
	Name   string
	Sprite *engine.Sprite
}

func (i *Item) Instantiate() *Item {
	item := *i

	if item.Sprite != nil {
		item.Sprite = item.Sprite.Instantiate()
	}

	return &item
}

func (i *Item) GetSprite() *engine.Sprite {
	if i.Sprite == nil {
		return nil
	}
	return i.Sprite
}

func (i *Item) GetName() string {
	return i.Name
}

func CreateItem(id int, name string, spritePath string) *Item {
	Image := assets.LoadImage(spritePath)

	return &Item{
		ID:   id,
		Name: name,
		Sprite: engine.NewSprite(
			name,
			name,
			Image,
		),
	}
}

type ItemManager struct {
	Items map[constants.ItemID]*Item
}

func NewItemManager() *ItemManager {
	manager := &ItemManager{
		Items: make(map[constants.ItemID]*Item),
	}

	manager.Items[constants.ItemCopper] = CreateItem(int(constants.ItemCopper), "Copper", "ores/copper.png")
	return manager
}

var Manager *ItemManager = NewItemManager()
