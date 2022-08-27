package tests

import (
	"sort"
	"testing"

	"github.com/bodagovsky/metro-project/handlers"
	"github.com/bodagovsky/metro-project/models"
	"github.com/stretchr/testify/require"

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
	testCases := []struct {
		id       int
		title    string
		expected *models.MetroLine
	}{
		{
			id:    1,
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
			id:    2,
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

func TestStorage_GetStationByLineID(t *testing.T) {
	storage := database.New()
	stations, err := storage.GetStationsByLineID(1)
	require.NoError(t, err)
	expected := []*models.MetroStation{
		{
			Id:         1,
			ExternalID: 1,
			Title:      "Бульвар Рокоссовского",
			LineID:     1,
			Neighbors:  []int{},
		},
		{
			Id:         2,
			ExternalID: 2,
			Title:      "Черкизовская",
			LineID:     1,
			Neighbors:  []int{},
		},
		{
			Id:         3,
			ExternalID: 3,
			Title:      "Преображенская площадь",
			LineID:     1,
			Neighbors:  []int{},
		},
		{
			Id:         4,
			ExternalID: 4,
			Title:      "Сокольники",
			LineID:     1,
			Neighbors:  []int{},
		},
		{
			Id:         5,
			ExternalID: 5,
			Title:      "Красносельская",
			LineID:     1,
			Neighbors:  []int{},
		},
		{
			Id:         6,
			ExternalID: 6,
			Title:      "Комсомольская",
			LineID:     1,
			Neighbors:  []int{},
		},
		{
			Id:         7,
			ExternalID: 7,
			Title:      "Красные Ворота",
			LineID:     1,
			Neighbors:  []int{},
		},
		{
			Id:         8,
			ExternalID: 8,
			Title:      "Чистые пруды",
			LineID:     1,
			Neighbors:  []int{},
		},
		{
			Id:         9,
			ExternalID: 9,
			Title:      "Лубянка",
			LineID:     1,
			Neighbors:  []int{},
		},
		{
			Id:         10,
			ExternalID: 10,
			Title:      "Охотный Ряд",
			LineID:     1,
			Neighbors:  []int{38},
		},
		{
			Id:         11,
			ExternalID: 11,
			Title:      "Библиотека имени Ленина",
			LineID:     1,
			Neighbors:  []int{},
		},
		{
			Id:         12,
			ExternalID: 12,
			Title:      "Кропоткинская",
			LineID:     1,
			Neighbors:  []int{},
		},
		{
			Id:         13,
			ExternalID: 13,
			Title:      "Парк культуры",
			LineID:     1,
			Neighbors:  []int{},
		},
		{
			Id:         14,
			ExternalID: 14,
			Title:      "Фрунзенская",
			LineID:     1,
			Neighbors:  []int{},
		},
		{
			Id:         15,
			ExternalID: 15,
			Title:      "Спортивная",
			LineID:     1,
			Neighbors:  []int{},
		},
		{
			Id:         16,
			ExternalID: 16,
			Title:      "Воробьёвы горы",
			LineID:     1,
			Neighbors:  []int{},
		},
		{
			Id:         17,
			ExternalID: 17,
			Title:      "Университет",
			LineID:     1,
			Neighbors:  []int{},
		},
		{
			Id:         18,
			ExternalID: 18,
			Title:      "Проспект Вернадского",
			LineID:     1,
			Neighbors:  []int{},
		},
		{
			Id:         19,
			ExternalID: 19,
			Title:      "Юго-Западная",
			LineID:     1,
			Neighbors:  []int{},
		},
		{
			Id:         20,
			ExternalID: 20,
			Title:      "Тропарёво",
			LineID:     1,
			Neighbors:  []int{},
		},
		{
			Id:         21,
			ExternalID: 21,
			Title:      "Румянцево",
			LineID:     1,
			Neighbors:  []int{},
		},
		{
			Id:         22,
			ExternalID: 22,
			Title:      "Саларьево",
			LineID:     1,
			Neighbors:  []int{},
		},
		{
			Id:         23,
			ExternalID: 23,
			Title:      "Филатов Луг",
			LineID:     1,
			Neighbors:  []int{},
		},
		{
			Id:         24,
			ExternalID: 24,
			Title:      "Прокшино",
			LineID:     1,
			Neighbors:  []int{},
		},
		{
			Id:         25,
			ExternalID: 25,
			Title:      "Ольховая",
			LineID:     1,
			Neighbors:  []int{},
		},
		{
			Id:         26,
			ExternalID: 26,
			Title:      "Коммунарка",
			LineID:     1,
			Neighbors:  []int{},
		},
	}

	for i, station := range stations {
		assert.Equal(t, expected[i], station)
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
			Id:         1,
			ExternalID: 1,
			Title:      "Бульвар Рокоссовского",
			IsClosed:   false,
			LineID:     1,
			Neighbors:  []int{},
		},
		{
			Id:         2,
			ExternalID: 2,
			Title:      "Черкизовская",
			IsClosed:   false,
			LineID:     1,
			Neighbors:  []int{},
		},
		{
			Id:         3,
			ExternalID: 3,
			Title:      "Преображенская площадь",
			IsClosed:   false,
			LineID:     1,
			Neighbors:  []int{},
		},
		{
			Id:         4,
			ExternalID: 4,
			Title:      "Сокольники",
			IsClosed:   false,
			LineID:     1,
			Neighbors:  []int{},
		},
		{
			Id:         5,
			ExternalID: 5,
			Title:      "Красносельская",
			IsClosed:   false,
			LineID:     1,
			Neighbors:  []int{},
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
			Id:         5,
			ExternalID: 5,
			Title:      "Красносельская",
			IsClosed:   false,
			LineID:     1,
			Neighbors:  []int{},
		},
		{
			Id:         4,
			ExternalID: 4,
			Title:      "Сокольники",
			IsClosed:   false,
			LineID:     1,
			Neighbors:  []int{},
		},
		{
			Id:         3,
			ExternalID: 3,
			Title:      "Преображенская площадь",
			IsClosed:   false,
			LineID:     1,
			Neighbors:  []int{},
		},
		{
			Id:         2,
			ExternalID: 2,
			Title:      "Черкизовская",
			IsClosed:   false,
			LineID:     1,
			Neighbors:  []int{},
		},
		{
			Id:         1,
			ExternalID: 1,
			Title:      "Бульвар Рокоссовского",
			IsClosed:   false,
			LineID:     1,
			Neighbors:  []int{},
		},
	}
	route, err := handlers.GetRoute(&getRoute, storage)
	assert.NoError(t, err)

	for i, r := range route.Path[0] {
		assert.Equal(t, expected[i], r)
	}
}

func TestStorage_GetRoute_DiffLine(t *testing.T) {
	storage := database.New()

	getRoute := &handlers.GetRouteRequest{
		From: handlers.Station{
			Id:        1,
			StationId: 1,
			LineId:    1,
		},
		To: handlers.Station{
			Id:        27,
			StationId: 1,
			LineId:    2,
		},
	}

	path, err := handlers.GetRoute(getRoute, storage)
	assert.NoError(t, err)
	assert.Equal(t, len(path.Path), 2)
}
