package models

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNode_TraverseDFS(t *testing.T) {
	graph := Node{
		Id:       1,
		LineID:   1,
		Title:    "test_0",
		IsClosed: false,
		Next: []*Node{
			{
				Id:     2,
				LineID: 1,
				Title:  "test_1",
				Next: []*Node{
					{
						Id:     4,
						LineID: 1,
						Title:  "test_3",
					},
				},
			},
			{
				Id:     3,
				LineID: 2,
				Title:  "test_2",
				Next:   []*Node{},
			},
		},
		visited: false,
	}

	expected := []*Node{
		{
			Title: "test_3",
		},
		{
			Title: "test_1",
		},
		{
			Title: "test_0",
		},
	}

	path := graph.TraverseDFS(4)
	require.NotNil(t, path)

	for i, node := range expected {
		assert.Equal(t, node.Title, path[0].Root[i].Title)
	}
}
