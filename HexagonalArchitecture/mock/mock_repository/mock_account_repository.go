package mock_repository

import (
	gomock "github.com/golang/mock/gomock"
)

// MockAccountRepository is a mock of AccountRepository interface.
type MockAccountRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAccountRepositoryMockRecorder
}

// MockAccountRepositoryMockRecorder is the mock recorder for MockAccountRepository.
type MockAccountRepositoryMockRecorder struct {
	mock *MockAccountRepository
}

// NewMockAccountRepository create a new instance
func NewMockAccountRepository(ctrl *gomock.Controller) *MockAccountRepository {
	mock := &MockAccountRepository{
		ctrl: ctrl,
	}
	mock.recorder = &MockAccountRepositoryMockRecorder{mock: mock}
	return mock
}
