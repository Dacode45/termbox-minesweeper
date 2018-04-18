package game

import (
	"github.com/nsf/termbox-go"
)

type KeyboardInput int

type KeyboardEvent struct {
	OriginalEvent termbox.Event
	Input         KeyboardInput
}

const (
	NOOP KeyboardInput = iota
	LEFT
	UP
	RIGHT
	DOWN
	CONFIRM
	BACK
	ACTION_1
	ESC
)

func keyToEvent(ev termbox.Event) KeyboardEvent {
	input := NOOP
	switch ev.Key {
	case termbox.KeyArrowLeft:
		input = LEFT
	case termbox.KeyArrowUp:
		input = UP
	case termbox.KeyArrowRight:
		input = RIGHT
	case termbox.KeyArrowDown:
		input = DOWN
	case termbox.KeySpace:
		input = CONFIRM
	case termbox.KeyBackspace:
		fallthrough
	case termbox.KeyBackspace2:
		input = BACK
	case termbox.KeyEsc:
		input = ESC
	default:
		switch ev.Ch {
		case 'w':
			input = UP
		case 's':
			input = DOWN
		case 'a':
			input = LEFT
		case 'd':
			input = RIGHT
		case 'f':
			input = ACTION_1
		}
	}
	return KeyboardEvent{ev, input}
}

func listenToKeyboard(evChan chan KeyboardEvent) {
	termbox.SetInputMode(termbox.InputEsc)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			evChan <- keyToEvent(ev)
		case termbox.EventError:
			panic(ev.Err)
		}

	}
}
