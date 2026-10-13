package engine

type Camera struct {
	X, Y                      float64
	ScreenWidth, ScreenHeight int
}

func (c *Camera) Follow(x, y float64) {
	c.X = x
	c.Y = y
	
	if x < float64(c.ScreenWidth / 2) {
		c.X = (float64(c.ScreenWidth) / 2)
	} 
	if y < float64(c.ScreenHeight / 2) {
		c.Y = (float64(c.ScreenHeight) / 2)
	} 
		
	

}
