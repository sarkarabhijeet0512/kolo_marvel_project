package marvel

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {
	var s *Service
	t.Run("Success", func(t *testing.T) {
		mockUserResp := &MarvelCharacterDetails{}

		mockUserRepository := new(MockUserService)
		us := NewService(s.conf, s.log, s.CacheService)
		payload := Payload{
			Page: 1,
		}
		mockUserRepository.On("FetchCharacterDetails", mock.Anything, payload).Return(mockUserResp, nil)

		// ctx := context.TODO()
		u, err := us.FetchCharacterDetails(&payload)

		assert.NoError(t, err)
		assert.Equal(t, u, mockUserResp)
		mockUserRepository.AssertExpectations(t)
	})
}
