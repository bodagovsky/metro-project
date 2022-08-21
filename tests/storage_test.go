package tests

import (
	"github.com/bodagovsky/metro-project/handlers"
	"github.com/bodagovsky/metro-project/models"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bodagovsky/metro-project/database"
)

func TestStorage_StationsSortedById(t *testing.T) {
	storage := database.New()
	stations, err := storage.GetStationsByLineID(1)
	assert.NoError(t, err, "failed to initiate storage")
	assert.True(t, sort.SliceIsSorted(stations, func(i, j int) bool {
		return stations[i].Id < stations[j].Id
	}))
}

func TestStorage_GetLineByID(t *testing.T) {
	storage := database.New()
	testCases := []struct{
		id int
		title string
		expected *models.MetroLine
	}{
		{
			id: 1,
			title: "Сокольническая",
			expected: &models.MetroLine{
				Id:       1,
				Title:    "Сокольническая",
				Color:    0,
				Stations: []int{},
				Crosses: map[string][]int{
					"2": {10},
				},
			},
		},
		{
			id: 2,
			title: "Замоскворецкая",
			expected: &models.MetroLine{
				Id:       2,
				Title:    "Замоскворецкая",
				Color:    1,
				Stations: []int{},
				Crosses: map[string][]int{
					"1": {38},
				},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.title, func(t *testing.T) {
			line, err := storage.GetLineByID(test.id)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, line)
		})
	}

}

func TestStorage_GetRoute_SameLine(t *testing.T) {
	storage := database.New()

	getRoute := handlers.GetRouteRequest{
		From: handlers.Station{
			Id:        1,
			StationId: 1,
			LineId:    1,
		},
		To: handlers.Station{
			Id:        5,
			StationId: 5,
			LineId:    1,
		},
	}

	expected := []*models.MetroStation{

		{
			Id:       1,
			Title:    "Бульвар Рокоссовского",
			IsClosed: false,
			LineID:   1,
		},
		{
			Id:       2,
			Title:    "Черкизовская",
			IsClosed: false,
			LineID:   1,
		},
		{
			Id:       3,
			Title:    "Преображенская площадь",
			IsClosed: false,
			LineID:   1,
		},
		{
			Id:       4,
			Title:    "Сокольники",
			IsClosed: false,
			LineID:   1,
		},
		{
			Id:       5,
			Title:    "Красносельская",
			IsClosed: false,
			LineID:   1,
		},
	}
	route, err := handlers.GetRoute(&getRoute, storage)
	assert.NoError(t, err)

	for i, r := range route.Path[0] {
		assert.Equal(t, expected[i], r)
	}
}

func TestStorage_GetRoute_SameLine_Inverted(t *testing.T) {
	storage := database.New()

	getRoute := handlers.GetRouteRequest{
		From: handlers.Station{
			Id:        5,
			StationId: 5,
			LineId:    1,
		},
		To: handlers.Station{
			Id:        1,
			StationId: 1,
			LineId:    1,
		},
	}

	expected := []*models.MetroStation{
		{
			Id:       5,
			Title:    "Красносельская",
			IsClosed: false,
			LineID:   1,
		},
		{
			Id:       4,
			Title:    "Сокольники",
			IsClosed: false,
			LineID:   1,
		},
		{
			Id:       3,
			Title:    "Преображенская площадь",
			IsClosed: false,
			LineID:   1,
		},
		{
			Id:       2,
			Title:    "Черкизовская",
			IsClosed: false,
			LineID:   1,
		},
		{
			Id:       1,
			Title:    "Бульвар Рокоссовского",
			IsClosed: false,
			LineID:   1,
		},
	}
	route, err := handlers.GetRoute(&getRoute, storage)
	assert.NoError(t, err)

	for i, r := range route.Path[0] {
		assert.Equal(t, expected[i], r)
	}
}
