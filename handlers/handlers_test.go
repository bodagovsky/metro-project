package handlers

import (
	"github.com/bodagovsky/metro-project/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	input := []*models.MetroStation{
		{
			Id:    1,
			Title: "Бульвар Рокоссовского",
		},
		{
			Id:    2,
			Title: "Черкизовская",
		},
		{
			Id:    3,
			Title: "Преображенская площадь",
		},
	}
	expected := []*models.MetroStation{
		{
			Id:    3,
			Title: "Преображенская площадь",
		},
		{
			Id:    2,
			Title: "Черкизовская",
		},
		{
			Id:    1,
			Title: "Бульвар Рокоссовского",
		},
	}
	reverse(input)
	for i, st := range input {
		assert.Equal(t, expected[i], st)
	}
}
