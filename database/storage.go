package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"

	"github.com/bodagovsky/metro-project/models"
)

const (
	linesPath    = "/Users/aabodagovskiy/metro-project/metro-project/database/lines.json"
	stationsPath = "/Users/aabodagovskiy/metro-project/metro-project/database/stations.json"
)

var _ Storage = &storage{}

type Storage interface {
	GetStationsByLineID(lineID int) ([]*models.MetroStation, error)
	GetLineByID(lineID int) (*models.MetroLine, error)
	GetStationByID(stationID int) (*models.MetroStation, error)
}

type storage struct {
	lines    map[string]*models.MetroLine
	stations map[string]*models.MetroStation
	//key - lineID, value - []metroStation
	stationsByLine map[int][]*models.MetroStation
	graph          map[int]*models.Node
}

func New() *storage {
	linesData := make(map[string]*models.MetroLine)
	stationsData := make(map[string]*models.MetroStation)

	linesFile, err := os.Open(linesPath)
	if err != nil {
		panic(err)
	}
	defer linesFile.Close()
	stationsFile, err := os.Open(stationsPath)
	if err != nil {
		panic(err)
	}
	defer stationsFile.Close()

	err = mapData(&linesData, linesFile)
	if err != nil {
		panic(fmt.Errorf("failed to init storage: %s", err.Error()))
	}
	err = mapData(&stationsData, stationsFile)
	if err != nil {
		panic(fmt.Errorf("failed to init storage: %s", err.Error()))
	}

	stationsByLine := make(map[int][]*models.MetroStation)

	for _, station := range stationsData {
		if _, ok := stationsByLine[station.LineID]; !ok {
			stationsByLine[station.LineID] = make([]*models.MetroStation, 0)
		}
		stationsByLine[station.LineID] = append(stationsByLine[station.LineID], station)
	}

	for lineID, stations := range stationsByLine {
		sort.Slice(stations, func(i, j int) bool {
			return stations[i].Id < stations[j].Id
		})
		stationsByLine[lineID] = stations
	}
	return &storage{
		lines:          linesData,
		stations:       stationsData,
		stationsByLine: stationsByLine,
		graph:          map[int]*models.Node{},
	}
}

func (s *storage) GetStationsByLineID(lineID int) ([]*models.MetroStation, error) {
	if stations, ok := s.stationsByLine[lineID]; ok {
		return stations, nil
	}
	return nil, errors.New(fmt.Sprintf("no line found with id %d", lineID))
}
func (s *storage) GetLineByID(lineID int) (*models.MetroLine, error) {
	id := strconv.Itoa(lineID)
	if line, ok := s.lines[id]; ok {
		return line, nil
	}
	return nil, errors.New(fmt.Sprintf("no line found with id %d", lineID))
}

func (s *storage) GetStationByID(stationID int) (*models.MetroStation, error) {
	id := strconv.Itoa(stationID)
	if station, ok := s.stations[id]; ok {
		return station, nil
	}
	return nil, errors.New(fmt.Sprintf("no station found with id %d", stationID))
}

func (s *storage) BuildGraph() *models.Node {
	lines, err := s.GetStationsByLineID(1)
	if err != nil {
		panic(err)
	}
	return s.traverse(0, nil, lines)
}

func (s *storage) traverse(i int, parent *models.Node, stations []*models.MetroStation) *models.Node {
	if i == len(stations) {
		return nil
	}
	station := stations[i]

	node := models.NewNode()
	s.graph[station.ExternalID] = node

	node.Id = station.ExternalID
	node.Title = station.Title
	node.IsClosed = station.IsClosed
	node.LineID = station.LineID
	if parent != nil {
		node.Next = append(node.Next, parent)
	}

	nextNode := s.traverse(i+1, node, stations)

	if nextNode != nil {
		node.Next = append(node.Next, nextNode)
	}

	for _, id := range station.Neighbors {
		if _, ok := s.graph[id]; ok {
			continue
		}

		nextStation, err := s.GetStationByID(id)
		if err != nil {
			panic(err)
		}
		nextStations, err := s.GetStationsByLineID(nextStation.LineID)
		if err != nil {
			panic(err)
		}
		i := nextStation.Id - 1
		leftPart := nextStations[:i]
		rightPart := nextStations[i+1:]

		models.Reverse(leftPart)

		nextNode = models.NewNode()

		nextNode.Id = nextStation.ExternalID
		nextNode.Title = nextStation.Title
		nextNode.IsClosed = nextStation.IsClosed
		nextNode.LineID = nextStation.LineID

		leftNode := s.traverse(0, nextNode, leftPart)
		rightNode := s.traverse(0, nextNode, rightPart)

		if leftNode != nil {
			nextNode.Next = append(nextNode.Next, leftNode)
		}

		if rightNode != nil {
			nextNode.Next = append(nextNode.Next, rightNode)
		}

		nextNode.Next = append(nextNode.Next, node)

		node.Next = append(node.Next, nextNode)
	}

	return node
}

func mapData(destination interface{}, source *os.File) error {
	switch t := destination.(type) {
	case *map[string]*models.MetroLine:
	case *map[string]*models.MetroStation:
	default:
		return errors.New(fmt.Sprintf("wrong destination type: %v", t))
	}
	body, err := io.ReadAll(source)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, destination)
	if err != nil {
		return err
	}
	return nil
}
