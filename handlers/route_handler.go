package handlers

import (
	"github.com/bodagovsky/metro-project/models"
)

type Storage interface {
	GetStationsByLineID(lineID int) ([]*models.MetroStation, error)
}

type Station struct{
	Id int
	StationId int
	LineId int
}

type GetRouteRequest struct {
	From Station
	To 	 Station
}

type GetRouteResponse struct {
	Path [][]*models.MetroStation
}

func GetRoute(request *GetRouteRequest, s Storage) (*GetRouteResponse, error) {
	response := &GetRouteResponse{
		Path: [][]*models.MetroStation{},
	}

	if request.From.LineId == request.To.LineId {
		var stationsOut []*models.MetroStation
		stations, err := s.GetStationsByLineID(request.From.LineId)
		if err != nil {
			return nil, err
		}
		if request.From.StationId > request.To.StationId {
			for _, station := range stations[request.To.StationId-1:request.From.StationId] {
				defer func(s *models.MetroStation) {
					stationsOut = append(stationsOut, s)
				}(station)
			}
		} else {
			stationsOut = append(stationsOut, stations[request.From.StationId-1:request.To.StationId]...)
		}
		response.Path = append(response.Path, stationsOut)
	}
	return response, nil
}