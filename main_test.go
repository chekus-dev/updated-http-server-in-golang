package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomePage(t *testing.T) {
	// 1. Setup the mock request and response recorder
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	rr := httptest.NewRecorder()

	// 2. Execute your handler
	HomePage(rr, req)

	// 3. Assert: Check for a successful 200 OK status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	// 4. Assert: Verify the Content-Type header
	// Go automatically normalizes "content-Type" to "application/json"
	expectedHeader := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedHeader {
		t.Errorf("Expected Content-Type %q, got %q", expectedHeader, contentType)
	}

	// 5. Assert: Decode the JSON and verify the fields exactly
	var user Message
	if err := json.Unmarshal(rr.Body.Bytes(), &user); err != nil {
		t.Fatalf("Failed to parse JSON response body: %v", err)
	}

	if user.ID != 1 {
		t.Errorf("Expected ID 1, got %d", user.ID)
	}

	if user.Name != "chekus joseph" {
		t.Errorf("Expected Name 'chekus joseph', got %q", user.Name)
	}

	if user.Email != "okaforchekus@gmail.com" {
		t.Errorf("Expected Email 'okaforchekus@gmail.com', got %q", user.Email)
	}
}
