package game

import (
	"math/rand"
)

type Vec2i struct {
	x int
	y int
}

type Board struct {
	Cells []*Cell
	width int
}

func NewBoard(width int, prob float32) *Board {
	cellCount := width * width
	cells := make([]*Cell, cellCount, cellCount)
	board := Board{cells, width}
	board.InitializeBoard(prob)
	return &board
}

func (b *Board) CellCount() int {
	return b.width * b.width
}

func (b *Board) InitializeBoard(prob float32) {
	cellCount := b.CellCount()
	expectedCap := int(float32(cellCount) * prob)
	mineIndicies := make([]int, 0, expectedCap)
	for i := range b.Cells {
		b.Cells[i] = &Cell{
			BombCount:  0,
			IsFlagged:  false,
			IsMine:     false,
			IsRevealed: false,
			Index:      i,
		}

		chance := rand.Float32()
		if chance < prob {
			b.Cells[i].IsMine = true
			mineIndicies = append(mineIndicies, i)
		}
	}
	// go through all mines and increment bomb count for adjacent mines
	for i := range mineIndicies {
		index := mineIndicies[i]
		adjacent := b.GetAdjacent(index)
		for _, cell := range adjacent {
			cell.BombCount++
		}
	}
}

func (b *Board) GetAdjacent(index int) []*Cell {
	x, y := b.IndexToXY(index)
	adjacent := [8]struct {
		x int
		y int
	}{
		{x - 1, y},
		{x + 1, y},
		{x, y - 1},
		{x - 1, y - 1},
		{x + 1, y - 1},
		{x, y + 1},
		{x - 1, y + 1},
		{x + 1, y + 1},
	}
	cells := make([]*Cell, 0, 8)
	for j := range adjacent {
		other := adjacent[j]
		if other.x < 0 || other.x >= b.width || other.y < 0 || other.y >= b.width {
			continue
		}
		adj := (other.y * b.width) + other.x
		cells = append(cells, b.Cells[adj])
	}
	return cells
}

func (b *Board) ClampWidth(num int) int {
	if num < 0 {
		return 0
	} else if num >= b.width {
		return b.width - 1
	}
	return num
}

func (b *Board) IndexToXY(index int) (int, int) {
	x := index % b.width
	y := index / b.width
	return x, y
}

func (b *Board) XYToIndex(x int, y int) int {
	return (y * b.width) + x
}

func (b *Board) GetCell(x int, y int) *Cell {
	index := (y * b.width) + x
	return b.Cells[index]
}

func (b *Board) RevealCell(x int, y int) {
	queue := make([]*Cell, 0)
	queue = append(queue, b.GetCell(x, y))
	for len(queue) != 0 {
		cell := queue[0]
		queue = queue[1:]
		cell.IsRevealed = true
		adj := b.GetAdjacent(cell.Index)
		for _, c := range adj {
			inQueue := false
			for _, o := range queue {
				if o == c {
					inQueue = true
					break
				}
			}
			if !inQueue && !c.IsMine && !c.IsRevealed && !c.IsFlagged {
				if c.BombCount == 0 {
					queue = append(queue, c)
				} else {
					c.IsRevealed = true
				}
			}
		}
	}
}

func (b *Board) IsWin() bool {
	for _, c := range b.Cells {
		if !c.IsRevealed {
			if c.IsMine && c.IsFlagged {
				continue
			} else {
				return false
			}
		}
	}
	return true
}

func (b *Board) IsLose() bool {
	for _, c := range b.Cells {
		if c.IsRevealed && c.IsMine {
			return true
		}
	}
	return false
}

func (b *Board) FlagCell(x int, y int, flag bool) {
	b.GetCell(x, y).IsFlagged = flag
}
