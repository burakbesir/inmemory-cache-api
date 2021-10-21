package controller_test

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ControllerSuite struct {
	suite.Suite
	cacheServiceMock *cacheServiceMock
}

func TestController(t *testing.T)  {
	suite.Run(t, new(ControllerSuite))
}

func (s *ControllerSuite) SetupSuite() {
	s.cacheServiceMock = new(cacheServiceMock)
}

func (s *ControllerSuite) TearDownTest() {
	s.cacheServiceMock = new(cacheServiceMock)
}

type cacheServiceMock struct {
	mock.Mock
	putCalledCount int
	key string
	value string
}

func (c *cacheServiceMock) Put(key string, value string) {
	c.putCalledCount++
	c.key = key
	c.value = value
}

func (c *cacheServiceMock) Get(key string) (string, error) {
	args := c.Called(key)
	return args.String(0), args.Error(1)
}