package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHandleAny(t *testing.T) {
	// mock the env
	err := os.Setenv("FORWARD_TO", "http://example.com")
	if err != nil {
		return
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/proxy/test", nil)
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	HandleAny(c)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}
}
