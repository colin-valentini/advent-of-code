package day2

type handShape int

const (
	rock handShape = iota
	paper
	scissors
)

func (s handShape) value() int {
	return int(s) + 1
}

func (s handShape) against(opponent handShape) outcome {
	// Go handles negative modulo operations unintuitively (to me).
	// This is why we're adding 3 to the difference between this
	// play and the opponent's play before taking the modulus which
	// prevents us from having a negative modulus operand.
	// See https://stackoverflow.com/questions/43018206/modulo-of-negative-integers-in-go.
	return outcome(((int(s-opponent)+3)%3 + 1) % 3)
}

func (s handShape) winsAgainst() handShape {
	return handShape((s - 1 + 3) % 3)
}

func (s handShape) losesAgainst() handShape {
	return handShape((s + 1 + 3) % 3)
}

func (s handShape) drawsAgainst() handShape {
	return s
}

type outcome int

const (
	loss outcome = iota
	draw
	win
)

func (o outcome) value() int {
	return int(o) * 3
}
