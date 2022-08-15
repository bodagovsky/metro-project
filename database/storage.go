package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/bodagovsky/metro-project/models"
)
const (
	linesPath = "/database/lines.json"
	stationsPath = "/database/stations.json"
)

type storage struct {
	lines map[string]*models.MetroLine
	stations map[string]*models.MetroStation
}

type Storage interface {
	Type() string
}

func New() *storage {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	linesData := make(map[string]*models.MetroLine)
	stationsData := make(map[string]*models.MetroStation)

	linesFile, err := os.Open(wd + linesPath)
	if err != nil {
		panic(err)
	}
	defer linesFile.Close()
	stationsFile, err := os.Open(wd + stationsPath)
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
	return &storage{
		linesData,
		stationsData,
	}
}

func (s *storage) Type() string {
	return "json storage"
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
