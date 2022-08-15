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

var colorsMap = map[Color]string{
	Red:        "красная",
	Green:      "зеленая",
	DarkBlue:   "",
	Blue:       "",
	Brown:      "",
	Orange:     "",
	Purple:     "",
	Yellow:     "",
	YellowA:    "",
	Gray:       "",
	GrassGreen: "",
	Turquoise:  "",
	TurquoiseA: "",
	PinkGray:   "",
	Pink:       "",
	MateBlue:   "",
	SoftPink:   "",
}

var LinesStorage = make(map[string]MetroLine)

type MetroStation struct {
	Id        int            `json:"id"`
	Title     string         `json:"title"`
	IsClosed  bool           `json:"is_closed"`
	OnCross   bool           `json:"on_cross"`
	Line      int            `json:"line"`
	Neighbors []int          `json:"neighbors"`
}

type MetroLine struct {
	Id       int            `json:"id"`
	Title    string         `json:"title"`
	Color    Color          `json:"color"`
	Stations []int 	        `json:"stations"`
	Crosses  []int          `json:"crosses"`
}

type Path struct {
}

func FindPath(from, to *MetroStation) Path {
	if shareSameLine(from, to) {

	}
	return Path{}
}

func shareSameLine(station1, station2 *MetroStation) bool {
	return station1.Line == station2.Line
}

func colorTitle(c Color) string {
	if color, ok := colorsMap[c]; ok {
		return color
	}
	return ""
}
