package mocks

import "github.com/stretchr/testify/mock"

type EnvironmentLoaderMock struct {
	mock.Mock
}

func (m *EnvironmentLoaderMock) Load(filenames ...string) (err error) {
	args := m.Called(filenames)
	return args.Error(0)
}
