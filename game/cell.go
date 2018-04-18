package game

import "strconv"

type Cell struct {
	Index      int
	BombCount  int
	IsRevealed bool
	IsMine     bool
	IsFlagged  bool
}

func (c Cell) DispChar() rune {
	if c.IsFlagged {
		return 'F'
	} else if !c.IsRevealed {
		return '?'
	} else if c.IsMine {
		return 'M'
	} else {
		if c.BombCount == 0 {
			return ' '
		}
		return rune(strconv.Itoa(c.BombCount)[0])
	}
}
