package services

import (
	"WebDevelopmentTodd/Webservices/REST_MS_GO/golang-microservices/src/api/clients/rest_client"
	"WebDevelopmentTodd/Webservices/REST_MS_GO/golang-microservices/src/api/domain/repositories"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	rest_client.StartMockups()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidInputName(t *testing.T) {
	request := repositories.CreateRepoRequest{
		Name:        "",
		Description: "golang-tutorial",
	}
	response, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "Invalid Repository Name", err.Message())
}

func TestCreateRepoFailedError(t *testing.T) {
	rest_client.FlushMockups()
	rest_client.AddMockups(rest_client.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message":"Requires authentication","documentation_url":"https://developer.github.com"}`)),
		},
	})
	request := repositories.CreateRepoRequest{
		Name: "testing",
	}
	response, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.Status())
	assert.EqualValues(t, "Requires authentication", err.Message())
}

func TestCreateRepoSuccess(t *testing.T) {
	rest_client.FlushMockups()
	rest_client.AddMockups(rest_client.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id":123}`)),
		},
	})
	request := repositories.CreateRepoRequest{
		Name: "testing",
	}
	response, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.EqualValues(t, 123, response.Id)
	assert.EqualValues(t, "", response.Name)
}
