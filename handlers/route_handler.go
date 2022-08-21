package handlers

import (
	"errors"
	"github.com/bodagovsky/metro-project/models"
	"strconv"
)

type Storage interface {
	GetStationsByLineID(lineID int) ([]*models.MetroStation, error)
	GetLineByID(lineID int) (*models.MetroLine, error)
}

type Station struct {
	Id        int
	StationId int
	LineId    int
}

type GetRouteRequest struct {
	From Station
	To   Station
}

type GetRouteResponse struct {
	Path [][]*models.MetroStation
}

var visitedLines = map[int]bool{}

func GetRoute(request *GetRouteRequest, s Storage) (*GetRouteResponse, error) {
	var stationsOut []*models.MetroStation
	var err error

	response := &GetRouteResponse{
		Path: [][]*models.MetroStation{},
	}

	if request.From.LineId == request.To.LineId {
		stationsOut, err = getRouteSameLine(request, s)
		if err != nil {
			return nil, err
		}
	} else {
		stationsOut, err = getRouteDifferentLane(request, s)
		if err != nil {
			return nil, err
		}
	}
	response.Path = append(response.Path, stationsOut)
	return response, nil
}

func getRouteSameLine(request *GetRouteRequest, s Storage) ([]*models.MetroStation, error) {
	var stationsOut []*models.MetroStation
	stations, err := s.GetStationsByLineID(request.From.LineId)
	if err != nil {
		return nil, err
	}
	if request.From.StationId > request.To.StationId {
		for _, station := range stations[request.To.StationId-1 : request.From.StationId] {
			//todo get rid of defer statement
			defer func(s *models.MetroStation) {
				stationsOut = append(stationsOut, s)
			}(station)
		}
	} else {
		stationsOut = append(stationsOut, stations[request.From.StationId-1:request.To.StationId]...)
	}
	return stationsOut, nil
}

func getRouteDifferentLane(request *GetRouteRequest, s Storage) ([]*models.MetroStation, error) {
	var stationsOut []*models.MetroStation
	var found bool

	start, err := s.GetLineByID(request.From.LineId)
	if err != nil {
		return nil, err
	}
	_, found = Traverse(start, request.To.LineId, s)
	if !found {
		return nil, errors.New("не нашелся путь от линии %s до %s")
	}

	return stationsOut, nil
}

func Traverse(node *models.MetroLine, target int, s Storage) ([]*models.MetroLine, bool) {
	if _, ok := visitedLines[node.Id] ; ok {
		return nil, false
	}
	visitedLines[node.Id] = true
	if node.Id == target {
		return []*models.MetroLine{node}, true
	}
	for lineID, _ := range node.Crosses {
		id, err := strconv.Atoi(lineID)
		if err != nil {
			panic(err)
		}
		line, err := s.GetLineByID(id)
		if err != nil {
			panic(err)
		}
		nodes, found := Traverse(line, target, s)
		if found {
			return append(nodes, node), found
		}
	}
	return nil, false
}
