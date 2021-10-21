package controller_test

import (
	"encoding/json"
	"github.com/rahmanbesir/inmemory-cache-api/controller"
	"github.com/rahmanbesir/inmemory-cache-api/error"
	"github.com/rahmanbesir/inmemory-cache-api/model"
	testifyAssert "github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
)

func (s *ControllerSuite) Test_it_should_put_value_with_key() {
	// Given
	var (
		assert           = testifyAssert.New(s.T())
		cacheServiceMock = s.cacheServiceMock
	)
	w := httptest.NewRecorder()
	cont := controller.New(cacheServiceMock)
	req := httptest.NewRequest(http.MethodPost, "/test", strings.NewReader("key=k1&value=v1"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// When
	cont.Handle(w, req)

	// Then
	assert.Equal(1, cacheServiceMock.putCalledCount)
	assert.Equal("k1", cacheServiceMock.key)
	assert.Equal("v1", cacheServiceMock.value)
}

func (s *ControllerSuite) Test_it_should_get_value() {
	// Given
	var (
		assert           = testifyAssert.New(s.T())
		cacheServiceMock = s.cacheServiceMock
	)
	w := httptest.NewRecorder()
	cont := controller.New(cacheServiceMock)
	req := httptest.NewRequest(http.MethodGet, "/test?key=k1", nil)

	cacheServiceMock.On("Get", "k1").Return("v1", nil)

	// When
	cont.Handle(w, req)

	// Then
	res := w.Result()
	defer res.Body.Close()
	var cacheResponse model.CacheResponse

	bytes, _ := io.ReadAll(res.Body)
	_ = json.Unmarshal(bytes, &cacheResponse)

	assert.Equal(200, res.StatusCode)
	assert.Equal("v1", cacheResponse.Value)
}

func (s *ControllerSuite) Test_it_should_return_error_if_cache_service_returns_error() {
	// Given
	var (
		assert           = testifyAssert.New(s.T())
		cacheServiceMock = s.cacheServiceMock
	)
	w := httptest.NewRecorder()
	cont := controller.New(cacheServiceMock)
	req := httptest.NewRequest(http.MethodGet, "/test?key=k1", nil)

	err := &error.ErrorResponse{
		StatusCode: 404,
	}
	cacheServiceMock.On("Get", "k1").Return("", err)

	// When
	cont.Handle(w, req)

	// Then
	res := w.Result()
	defer res.Body.Close()
	assert.Equal(404, res.StatusCode)
}

func (s *ControllerSuite) Test_it_should_return_bad_request_if_key_parameter_not_exist() {
	// Given
	var (
		assert           = testifyAssert.New(s.T())
		cacheServiceMock = s.cacheServiceMock
	)
	w := httptest.NewRecorder()
	cont := controller.New(cacheServiceMock)
	req := httptest.NewRequest(http.MethodGet, "/test", nil)

	// When
	cont.Handle(w, req)

	// Then
	res := w.Result()
	defer res.Body.Close()
	var errorResponse error.ErrorResponse

	bytes, _ := io.ReadAll(res.Body)
	_ = json.Unmarshal(bytes, &errorResponse)

	assert.Equal(400, res.StatusCode)
	assert.Equal("ValidationError", errorResponse.ErrorName)
}

func (s *ControllerSuite) Test_it_should_return_method_not_allowed() {
	// Given
	var (
		assert           = testifyAssert.New(s.T())
		cacheServiceMock = s.cacheServiceMock
	)
	w := httptest.NewRecorder()
	cont := controller.New(cacheServiceMock)
	req := httptest.NewRequest(http.MethodPut, "/test", nil)

	// When
	cont.Handle(w, req)

	// Then
	res := w.Result()
	defer res.Body.Close()

	assert.Equal(405, res.StatusCode)
}
