package handlers

import (
	"github.com/bodagovsky/metro-project/models"
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

	return stationsOut, nil
}
