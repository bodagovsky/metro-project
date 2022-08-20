package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/bodagovsky/metro-project/models"
)
const (
	linesPath = "/Users/aabodagovskiy/metro-project/metro-project/database/lines.json"
	stationsPath = "/Users/aabodagovskiy/metro-project/metro-project/database/stations.json"
)

type storage struct {
	lines map[string]*models.MetroLine
	stations map[string]*models.MetroStation
	//key - lineID, value - []metroStation
	stationsByLine map[int][]*models.MetroStation
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
		lines:linesData,
		stations: stationsData,
		stationsByLine: stationsByLine,
	}
}

func (s *storage) GetStationsByLineID(lineID int) ([]*models.MetroStation, error) {
	if stations, ok := s.stationsByLine[lineID]; ok {
		return stations, nil
	}
	return nil, errors.New("no such line")
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
