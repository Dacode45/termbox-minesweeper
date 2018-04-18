package game

type Difficulty int

const (
	EASY = iota
	MEDIUM
	HARD
)

func (d Difficulty) getWidth() int {
	switch d {
	case MEDIUM:
		return 20
	case HARD:
		return 30
	default:
		return 10
	}
}

func (d Difficulty) getProb() float32 {
	switch d {
	case MEDIUM:
		return 0.05
	case HARD:
		return 0.1
	default:
		return 0.15
	}
}
