package constants

type LayerID int


type TileID int

const (
	TileOreCopper TileID = iota
	TileOreIron
	TileOreGold

	TileGrass
)

type ItemID int

const (
	ItemCopper ItemID = ItemID(TileOreCopper)
)

type BuildID int

const (
	BuildCopperMine BuildID = iota
)
