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
	MateBlue
	SoftPink
)

type MetroStation struct {
	Id         int    `json:"id"`
	ExternalID int    `json:"externalID"`
	Title      string `json:"title"`
	IsClosed   bool   `json:"is_closed"`
	OnCross    bool   `json:"on_cross"`
	LineID     int    `json:"lineID"`
	Neighbors  []int  `json:"neighbors"`
}

type MetroLine struct {
	Id       int              `json:"id"`
	Title    string           `json:"title"`
	Color    Color            `json:"color"`
	Stations []int            `json:"stations"`
	Crosses  map[string][]int `json:"crosses"`
}

func BuildGraph() *Node {
	return &Node{}
}
