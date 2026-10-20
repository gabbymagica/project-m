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
	BuildTest1
	BuildTest2
	BuildTest3
)

type ActionID int

const (
	Idling ActionID = iota
	PreviewBuilding
	Building
	Mining
	PreviewDestroying
	Destroying
)
