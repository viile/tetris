package block

type BlockS struct {
	Te
}

func NewBlockS(w,h int) *BlockS {
	return &BlockS{Te{[]Block{
		Block{w, h, true},
		{w - 1, h, false},
		{w - 1, h-1, false},
		{w - 2, h-1, false}}}}
}