package inventory

type itemInterface interface { } 

type InventoryItem struct {
	Quantity int
	Item itemInterface
}

type Inventory struct {
	Items [10]InventoryItem
}

func (i *Inventory) AddItem(item itemInterface, quantity int) {
	for index, inventoryItem := range i.Items {
		switch inventoryItem.Item {
			case nil:
				i.Items[index] = InventoryItem{
					Quantity: quantity,
					Item: item,
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
		if inventoryItem.Item == item  {
			if inventoryItem.Quantity > quantity {
				i.Items[index].Quantity -= quantity 
			} else {
				i.Items[index] = InventoryItem{}
			}
			return
		}
	}
}