package repositories_controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateRepoInvalidJSONRequest(t *testing.T) {
	resp := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(resp)
	request, _ := http.NewRequest(http.MethodPost, "/repo", strings.NewReader(``))
	c.Request = request 
	CreateRepo(c)
	assert.EqualValues(t, http.StatusBadRequest, resp.Code)
}
