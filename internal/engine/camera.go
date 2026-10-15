package engine

import "math"

type Camera struct {
	X, Y                      float64
	Tilesize                  int
	ScreenWidth, ScreenHeight int
}

func (c *Camera) Follow(x, y float64) {
	c.X = x
	c.Y = y

	if x < float64(c.ScreenWidth/2) {
		c.X = (float64(c.ScreenWidth) / 2)
	}
	if y < float64(c.ScreenHeight/2) {
		c.Y = (float64(c.ScreenHeight) / 2)
	}
}

func (c *Camera) WorldToScreen(worldX, worldY float64) (screenX, screenY float64) {
	originX := c.X - float64(c.ScreenWidth)/2
	originY := c.Y - float64(c.ScreenHeight)/2

	screenX = worldX - originX
	screenY = worldY - originY
	return screenX, screenY
}

func (c *Camera) ScreenToWorld(screenX, screenY float64) (worldX, worldY float64) {
	originX := c.X - float64(c.ScreenWidth)/2
	originY := c.Y - float64(c.ScreenHeight)/2

	worldX = screenX + originX
	worldY = screenY + originY
	return worldX, worldY
}

func (c *Camera) ScreenToTile(screenX, screenY float64, tileSize int) (tileX, tileY int) {
	worldX, worldY := c.ScreenToWorld(screenX, screenY)
	tileX = int(math.Floor(worldX / float64(tileSize)))
	tileY = int(math.Floor(worldY / float64(tileSize)))
	return tileX, tileY
}

func (c *Camera) VisibleTileBounds(tileSize, mapSizeX, mapSizeY int) (minX, minY, maxX, maxY int) {
	ts := float64(tileSize)
	originX := c.X - float64(c.ScreenWidth)/2
	originY := c.Y - float64(c.ScreenHeight)/2

	minX = int(math.Floor(originX / ts))
	minY = int(math.Floor(originY / ts))
	maxX = int(math.Floor((originX+float64(c.ScreenWidth))/ts)) + 1
	maxY = int(math.Floor((originY+float64(c.ScreenHeight))/ts)) + 1

	if minX < 0 {
		minX = 0
	}
	if minY < 0 {
		minY = 0
	}
	if maxX > mapSizeX {
		maxX = mapSizeX
	}
	if maxY > mapSizeY {
		maxY = mapSizeY
	}

	return minX, minY, maxX, maxY
}
