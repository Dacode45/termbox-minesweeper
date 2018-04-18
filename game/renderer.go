package game

import (
	"fmt"
	"time"

	runewidth "github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

const (
	ColorDefault    = termbox.ColorDefault
	ColorBackground = termbox.ColorDefault
	ColorSelected   = termbox.ColorCyan
	ColorDeselected = termbox.ColorWhite
	ColorCursor     = termbox.ColorBlue
	ColorCursorAlt  = termbox.ColorCyan
)

type Window struct {
	w  int
	h  int
	hw int
	hh int
}

func (g *Game) render(t time.Time) error {
	termbox.Clear(ColorDefault, ColorDefault)
	w, h := termbox.Size()
	wnd := Window{w, h, w / 2, h / 2}

	switch g.State {
	case MainMenu:
		g.renderMainMenu(wnd, t)
	case RunGame:
		g.renderRun(wnd, t)
	case WinLoseGame:
		g.renderWinLoose(wnd, t)
	}

	return termbox.Flush()
}

func (g *Game) renderDebug(wnd Window, t time.Time) {
	tbPrint(0, 0, ColorDefault, ColorBackground, fmt.Sprintf("%v %v", g, wnd))
}

func (g *Game) renderWinLoose(wnd Window, t time.Time) {
	msg := "You won!"
	if g.SubState == Lose {
		msg = "You lost!"
	}
	tbPrint(0, 0, ColorDefault, ColorBackground, msg)
}

func (g *Game) renderMainMenu(wnd Window, t time.Time) {
	// divide the arena into cells where the options will be
	cellWidth := wnd.w / (2*len(g.Options) + 1)
	for i, option := range g.Options {
		y := wnd.hh
		x := (1 + (2 * i)) * cellWidth
		color := ColorDeselected
		if i == g.Cursor {
			color = ColorSelected
		}
		tbPrint(x, y, color, ColorBackground, option)
	}
}

func (g *Game) renderRun(wnd Window, t time.Time) {
	boardLeft := (wnd.w - g.Board.width) / 2
	boardTop := (wnd.h - g.Board.width) / 2
	// draw the time
	minutes := g.Time / time.Minute
	seconds := (g.Time % (60 * time.Second)) / (time.Second)
	tbPrint(0, 0, ColorDefault, ColorBackground, fmt.Sprintf("%02d:%02d", minutes, seconds))

	// draw the board
	for i, cell := range g.Board.Cells {
		x, y := g.Board.IndexToXY(i)
		termbox.SetCell(boardLeft+x, boardTop+y, cell.DispChar(), ColorDefault, ColorBackground)
	}

	// draw cursor
	oddSecond := seconds%2 == 1
	cursorColor := ColorCursor
	if oddSecond {
		cursorColor = ColorCursorAlt
	}
	cx, cy := g.Board.IndexToXY(g.Cursor)
	termbox.SetCell(boardLeft+cx, boardTop+cy, '+', cursorColor, ColorBackground)

}

func tbPrint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}
