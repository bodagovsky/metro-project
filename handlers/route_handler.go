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
		reverse(stationsOut)
	}

	return stationsOut, nil
}

func getRouteDifferentLine(request *GetRouteRequest, s database.Storage) ([]*models.MetroStation, error) {
	var stationsOut []*models.MetroStation
	var found bool

	_, err := s.GetLineByID(request.From.LineId)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, errors.New("не нашелся путь от линии %s до %s")
	}

	return stationsOut, nil
}

func reverse(stations []*models.MetroStation) {
	i := 0
	j := len(stations) - 1
	for i < j {
		stations[i], stations[j] = stations[j], stations[i]
		i++
		j--
	}
}

func minMax(i, j int) (min int, max int) {
	if i > j {
		return j, i
	}
	return i, j
}
