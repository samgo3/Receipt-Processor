package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepository_GetById(t *testing.T) {

	// Initialize a new mock repository
	mockRepo := NewKVRepo()
	generateId, _ := mockRepo.AddEntry(78)

	tests := []struct {
		name     string
		id       string
		expected interface{}
	}{
		{
			name:     "RepoGetById Valid ID Test",
			id:       generateId,
			expected: int64(78),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			point, _ := mockRepo.GetById(test.id)
			assert.Equal(t, test.expected, point, "Test Failed")
		})
	}
	t.Run("RepoGetById Invalid id Test", func(t *testing.T) {
		_, err := mockRepo.GetById("Invalid")
		assert.Error(t, err, "Expected an error for invalid ID")
	})

}

func TestRepository_Create(t *testing.T) {
	mockRepo := NewKVRepo()
	id1, _ := mockRepo.AddEntry(3)
	id2, _ := mockRepo.AddEntry(14)
	t.Run("RepoAddEntry Test", func(t *testing.T) {
		assert.NotEmpty(t, id1, "Expected a non-empty ID")
		assert.NotEmpty(t, id2, "Expected a non-empty ID")
	})

	t.Run("RepoAddEntry Unique Id test", func(t *testing.T) {
		assert.NotEqual(t, id1, id2, "Expected different IDs for different creations")
	})

}
