package block

type BlockZ struct {
	Te
}

func NewBlockZ(w,h int) BlockI {
	return BlockI{Te{[]Block{
		Block{w, h, true},
		{w + 1, h, false},
		{w + 1, h-1, false},
		{w + 2, h-1, false}}}}
}