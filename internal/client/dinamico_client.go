package client

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/rbartolome/chatrooms/internal/entity"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	commons "gitlab.falabella.com/fif/integracion/forthehorde/commons/go-microservices-commons"
)

//MakeHTTPClientEndpoint crea un client para el GET character
func MakeHTTPClientCharacterEndpoint(url string, timeout time.Duration, logger log.Logger) endpoint.Endpoint {

	url_dinamica := url + "/character/:id"

	return commons.MakeHTTPClientBuilder("GET",
		url_dinamica,
		timeout,
		EncodeRequestCharacters,
		DecodeResponseCharacters,
		logger).
		Build()
}

func DecodeResponseCharacters(_ context.Context, r *http.Response) (interface{}, error) {

	response := new(entity.CharacterResponse)

	err := json.NewDecoder(r.Body).Decode(response)
	if err != nil {
		return nil, errors.New("Error en Decode Response Characters de client")
	}

	return response, nil
}

func EncodeRequestCharacters(_ context.Context, r *http.Request, req interface{}) error {

	request := req.(*entity.Request)
	id := strconv.Itoa(request.Id_character)

	r.URL.Path = strings.Replace(r.URL.Path, ":id", id, -1)

	return nil

}
