package game

import (
	"runtime"
	"time"

	"github.com/nsf/termbox-go"
)

type GameState int
type GameSubState int

const (
	MainMenu GameState = iota
	StartGame
	RunGame
	PauseGame
	WinLoseGame
	ExitGame
)

// MainMenuState
const (
	ChooseMode GameSubState = iota
	ChooseDifficulty
)

type Game struct {
	StartTime time.Time
	Time      time.Duration
	Board     *Board
	State     GameState
	SubState  GameSubState
	Options   []string
	Cursor    int

	ticker             *time.Ticker
	renderTicker       *time.Ticker
	keyboardEventsChan chan KeyboardEvent
}

func NewGame() *Game {
	ticker := time.NewTicker(1 * time.Second)
	renderTicker := time.NewTicker(time.Second / 30)
	keyboardEventsChan := make(chan KeyboardEvent)

	return &Game{
		Time:               0,
		Board:              nil,
		State:              MainMenu,
		ticker:             ticker,
		renderTicker:       renderTicker,
		keyboardEventsChan: keyboardEventsChan,
	}
}

func (g *Game) Init() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	go listenToKeyboard(g.keyboardEventsChan)

	g.EnterMainMenuState()

mainloop:
	for {
		select {
		case k := <-g.keyboardEventsChan:
			if k.Input == ESC {
				break mainloop
			}
			g.handleEvent(k)
		case <-g.ticker.C:
			if g.State == RunGame {
				g.Time += 1 * time.Second
			}
		case t := <-g.renderTicker.C:
			if err := g.render(t); err != nil {
				panic(err)
			}
		default:
			if g.State == ExitGame {
				break mainloop
			}
			runtime.Gosched()
		}
	}
}

func (g *Game) handleEvent(k KeyboardEvent) {
	switch g.State {
	case MainMenu:
		g.handleEventMainMenu(k)
	case RunGame:
		g.handleEventRun(k)
	}
}
