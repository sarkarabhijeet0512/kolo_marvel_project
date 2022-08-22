package marvel

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// MockUserService is a mock type for model.UserService
type MockUserService struct {
	mock.Mock
}

// ClearProfileImage is a mock of UserService.ClearProfileImage
func (m *MockUserService) FetchCharacterDetails(ctx context.Context, payload Payload) (mcd *MarvelCharacterDetails, err error) {
	ret := m.Called(ctx, payload)

	// first value passed to "Return"
	var r0 *MarvelCharacterDetails
	if ret.Get(0) != nil {
		// we can just return this if we know we won't be passing function to "Return"
		r0 = ret.Get(0).(*MarvelCharacterDetails)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}
