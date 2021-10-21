package controller

import (
	"encoding/json"
	"errors"
	internalerror "github.com/rahmanbesir/inmemory-cache-api/error"
	"github.com/rahmanbesir/inmemory-cache-api/model"
	"github.com/rahmanbesir/inmemory-cache-api/service"
	"net/http"
)

type Controller interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	cacheService service.CacheService
}

func New(cacheService service.CacheService) Controller {
	return &controller{cacheService: cacheService}
}

func (c *controller) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		key := r.PostFormValue("key")
		val := r.PostFormValue("value")
		c.cacheService.Put(key, val)
		w.WriteHeader(http.StatusOK)
	case http.MethodGet:
		key, exist := r.URL.Query()["key"]
		w.Header().Set("Content-Type", "application/json")

		if !exist {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(convertErrToByteResponse(internalerror.CreateValidationError(errors.New("key url parameter doesn't exist"))))
			return
		}

		val, err := c.cacheService.Get(key[0])
		if err != nil {
			errorsResponse := err.(*internalerror.ErrorResponse)
			w.WriteHeader(errorsResponse.StatusCode)
			w.Write(convertErrToByteResponse(err))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(createCacheResponseAsBytes(val))

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
	}
}

func convertErrToByteResponse(err error) []byte {
	return []byte(err.Error())
}

func createCacheResponseAsBytes(val string) []byte {
	resp := model.CacheResponse{Value: val}
	bytes, _ := json.Marshal(resp)
	return bytes
}
