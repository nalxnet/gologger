package mock

import (
	"github.com/stretchr/testify/mock"
)

type LoggerMock struct {
	mock.Mock
}

func NewLoggerMock() *LoggerMock {
	return &LoggerMock{}
}

func (m *LoggerMock) Error(msg string) {
	m.Called(msg)
}

func (m *LoggerMock) Errorf(format string, v ...interface{}) {
	m.Called(format, v)
}

func (m *LoggerMock) ErrorErr(err error) {
	m.Called(err)
}

func (m *LoggerMock) Warn(msg string) {
	m.Called(msg)
}

func (m *LoggerMock) Warnf(format string, v ...interface{}) {
	m.Called(format, v)
}

func (m *LoggerMock) WarnErr(err error) {
	m.Called(err)
}

func (m *LoggerMock) Info(msg string) {
	m.Called(msg)
}

func (m *LoggerMock) Infof(format string, v ...interface{}) {
	m.Called(format, v)
}

func (m *LoggerMock) Debug(msg string) {
	m.Called(msg)
}

func (m *LoggerMock) Debugf(format string, v ...interface{}) {
	m.Called(format, v)
}
