package game

import "time"

var ChooseModeOptions = []string{
	"Start Game",
	"Exit Game",
}

var ChooseDifficultyOptions = []string{
	"Easy",
	"Medium",
	"Hard",
}

func (g *Game) EnterMainMenuState() {
	g.State = MainMenu
	g.SubState = ChooseMode
	g.Cursor = 0
	g.Options = ChooseModeOptions
}

func (g *Game) EnterStartState(diff Difficulty) {
	g.State = StartGame
	width := diff.getWidth()
	prob := diff.getProb()
	board := NewBoard(width, prob)
	g.Board = board
	g.Time = 0
	g.StartTime = time.Now()
	g.Cursor = 0
	g.EnterRunState()
}

func (g *Game) EnterRunState() {
	g.State = RunGame
}

func (g *Game) handleEventRun(k KeyboardEvent) {
	x := 0
	y := 0
	if k.Input == LEFT {
		x--
	} else if k.Input == RIGHT {
		x++
	} else if k.Input == UP {
		y--
	} else if k.Input == DOWN {
		y++
	}
	nx, ny := g.Board.IndexToXY(g.Cursor)
	nx = g.Board.ClampWidth(nx + x)
	ny = g.Board.ClampWidth(ny + y)

	g.Cursor = g.Board.XYToIndex(nx, ny)

	if k.Input == CONFIRM {
		g.Board.RevealCell(nx, ny)
	}
}

func (g *Game) handleEventMainMenu(k KeyboardEvent) {
	if k.Input == LEFT {
		g.decrementCursor()
	} else if k.Input == RIGHT {
		g.incrementCursor()
	}

	switch g.SubState {
	case ChooseMode:
		if k.Input == CONFIRM {
			g.Options = ChooseDifficultyOptions
			g.SubState = ChooseDifficulty
			g.Cursor = 0
		}
	case ChooseDifficulty:
		if k.Input == BACK {
			g.Options = ChooseModeOptions
			g.SubState = ChooseMode
			g.Cursor = 0
		} else if k.Input == CONFIRM {
			g.EnterStartState(Difficulty(g.Cursor))
		}
	}
}

func (g *Game) incrementCursor() {
	g.Cursor = (g.Cursor + 1) % len(g.Options)
}

func (g *Game) decrementCursor() {
	g.Cursor = (g.Cursor - 1)
	if g.Cursor < 0 {
		g.Cursor = len(g.Options) - 1
	}
}
