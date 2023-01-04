package service

import (
	"errors"
	"ikuzports/features/category"
	"ikuzports/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	repo := new(mocks.CategoryRepo)

	t.Run("Success get all categories", func(t *testing.T) {
		data := []category.CategoryCore{{Name: "name"}}
		repo.On("GetAll").Return(data, nil).Once()

		srv := New(repo)
		res, err := srv.GetAll()
		assert.NoError(t, err)
		assert.Equal(t, data[0].Name, res[0].Name)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get all categories", func(t *testing.T) {
		repo.On("GetAll").Return(nil, errors.New("failed")).Once()

		srv := New(repo)
		_, err := srv.GetAll()
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}
