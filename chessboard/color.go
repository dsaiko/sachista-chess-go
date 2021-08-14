package chessboard

type Color int

const (
	White Color = 0
	Black Color = 1
)

func (c Color) String() string {
	switch c {
	case White:
		return "w"
	case Black:
		return "b"
	default:
		return "?"
	}
}
