package service_test

import (
	"github.com/rahmanbesir/inmemory-cache-api/error"
	"github.com/rahmanbesir/inmemory-cache-api/service"
	testifyAssert "github.com/stretchr/testify/assert"
	"testing"
)

func Test_it_should_get(t *testing.T) {
	// Given
	var (
		assert = testifyAssert.New(t)
	)

	cacheService := service.NewCacheService()
	cacheService.Put("k1", "v1")

	//When
	val , err := cacheService.Get("k1")

	//Then
	assert.Nil(err)
	assert.Equal("v1", val)
}


func Test_it_should_return_err_if_key_not_exist(t *testing.T) {
	// Given
	var (
		assert = testifyAssert.New(t)
	)

	cacheService := service.NewCacheService()

	//When
	_ , err := cacheService.Get("k1")

	//Then
	errorResponse := err.(*error.ErrorResponse)
	assert.Equal(404, errorResponse.StatusCode)
	assert.Equal("NotFoundError", errorResponse.ErrorName)
}
