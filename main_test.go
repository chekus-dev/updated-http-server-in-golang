package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomePage_Success(t *testing.T) {
	// Arrange: Create mock request
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rr := httptest.NewRecorder()

	// Act: Run the handler directly
	HomePage(rr, req)

	// Assert 1: Check that omitting WriteHeader still yields 200 OK
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	// Assert 2: Check standard Canonical header case
	expectedHeader := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedHeader {
		t.Errorf("Expected Content-Type %q, got %q", expectedHeader, contentType)
	}

	// Assert 3: Decode and verify body strings EXACTLY
	var actualResponse Message
	err = json.Unmarshal(rr.Body.Bytes(), &actualResponse)
	if err != nil {
		t.Fatalf("Failed to decode response JSON: %v", err)
	}

	expectedMessage := "welcome dear user to updated http server in go"
	if actualResponse.Message != expectedMessage {
		t.Errorf("Expected message %q, got %q", expectedMessage, actualResponse.Message)
	}

	if actualResponse.Reply != "thank you" {
		t.Errorf("Expected reply 'thank you', got %q", actualResponse.Reply)
	}
}
