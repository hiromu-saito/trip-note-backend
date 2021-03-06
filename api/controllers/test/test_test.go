package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestTest(t *testing.T) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(
		http.MethodGet,
		"/",
		nil,
	)

	// Act ---
	Test(c)

	// Assert ---
	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.Equal(t, "Test!", response.Body.String())
}
