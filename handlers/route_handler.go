package handlers

import (
	"errors"

	"github.com/bodagovsky/metro-project/database"
	"github.com/bodagovsky/metro-project/models"
)

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

func GetRoute(request *GetRouteRequest, s database.Storage) (*GetRouteResponse, error) {
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
		stationsOut, err = getRouteDifferentLine(request, s)
		if err != nil {
			return nil, err
		}
	}
	response.Path = append(response.Path, stationsOut)
	return response, nil
}

func getRouteSameLine(request *GetRouteRequest, s database.Storage) ([]*models.MetroStation, error) {
	var stationsOut []*models.MetroStation
	stations, err := s.GetStationsByLineID(request.From.LineId)
	if err != nil {
		return nil, err
	}
	i, j := minMax(request.From.Id, request.To.Id)

	stationsOut = append(stationsOut, stations[i-1:j]...)
	if request.From.Id > request.To.Id {
		models.Reverse(stationsOut)
	}

	return stationsOut, nil
}

func getRouteDifferentLine(request *GetRouteRequest, s database.Storage) ([]*models.MetroStation, error) {
	var stationsOut []*models.MetroStation
	var found bool

	root, err := s.GetNodeById(request.From.Id)
	if err != nil {
		return nil, err
	}

	path := root.TraverseDFS(request.To.Id)
	if !found {
		return nil, errors.New("не нашелся путь от линии %s до %s")
	}

	//todo extract the best path
	stationsOut = buildPath(path[0].Root)

	return stationsOut, nil
}

func buildPath(path []*models.Node) []*models.MetroStation {
	var outPath = make([]*models.MetroStation, 0, len(path))

	i := len(path) - 1
	for i >= 0 {
		var station = &models.MetroStation{}
		station.Id = path[i].Id
		station.LineID = path[i].LineID
		station.Title = path[i].Title
		outPath = append(outPath, station)
		i--
	}

	return outPath
}

func minMax(i, j int) (min int, max int) {
	if i > j {
		return j, i
	}
	return i, j
}
