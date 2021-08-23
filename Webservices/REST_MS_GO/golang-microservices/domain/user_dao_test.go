package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNotFound(t *testing.T) {
	user, err := GetUser(0)
	assert.NotNil(t, err)
	assert.Nil(t, user, "We were not expecting an user with ID 0")

	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "User 0 is not found", err.Message)
}

func TestGetUserFound(t *testing.T) {
	user, err := GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Hari", user.FirstName)
	assert.EqualValues(t, "Prathap", user.LastName)
	assert.EqualValues(t, "hfiery@gmail.com", user.Email)
}
