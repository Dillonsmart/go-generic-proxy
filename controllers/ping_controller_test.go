package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	// Simulate a request to the Ping handler
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Call the Ping handler
	Ping(c)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	// Check the response body
	expectedBody := `{"message":"pong"}`
	if w.Body.String() != expectedBody {
		t.Errorf("Expected body %s, got %s", expectedBody, w.Body.String())
	}
}
