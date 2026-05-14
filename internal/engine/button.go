package engine

type Button struct {
	X, Y int
	Width, Height int
	Label string
	Sprite *Sprite
}

func (b *Button) IsClicked(mouseX, mouseY int, callback func()) bool {
	if mouseX >= b.X && mouseX <= b.X + b.Width &&
		mouseY >= b.Y && mouseY <= b.Y + b.Height {
		callback()
		return true
	}
	return false
}