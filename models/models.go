package models

type Color int

const (
	Red Color = iota
	Green
	DarkBlue
	Blue
	Brown
	Orange
	Purple
	Yellow
	YellowA
	Gray
	GrassGreen
	Turquoise
	TurquoiseA
	PinkGray
	Pink
)

type MetroStation struct {
	id int
	title string
	isClosed bool
	onCross bool
	line MetroLine
	neighbors []MetroStation
}

type MetroLine struct {
	id int
	color Color
	stations []MetroStation
	crosses []MetroLine
}

