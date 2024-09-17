package business

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"urbanmedia/go-cli-boilerplate/config"
)

// Mock config for testing
func mockConfig(apiURL string) {
	config.AppConfig.ApiURL = apiURL
}

func TestFetchInfo_Success(t *testing.T) {
	// Mock server with a plain text response (e.g., IP address)
	mockResponse := "8.8.8.8"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	// Update config to point to the mock server
	mockConfig(server.URL)

	// Call the function under test
	result, err := FetchInfo()

	// Assert no error was returned
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Assert the plain text response is as expected
	if result != mockResponse {
		t.Errorf("Expected %s, got %s", mockResponse, result)
	}
}

func TestFetchInfo_ErrorResponse(t *testing.T) {
	// Mock server that returns a 500 error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}))
	defer server.Close()

	// Update config to point to the mock server
	mockConfig(server.URL)

	// Call the function under test
	_, err := FetchInfo()

	// Assert that an error was returned
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	// Optionally, check the error message
	if !strings.Contains(err.Error(), "failed to fetch info") {
		t.Errorf("Expected error to contain 'failed to fetch info', got %v", err)
	}
}

func TestFetchInfo_InvalidPlainTextResponse(t *testing.T) {
	// Mock server with an invalid plain text response
	invalidResponse := "invalid-ip-response"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(invalidResponse))
	}))
	defer server.Close()

	// Update config to point to the mock server
	mockConfig(server.URL)

	// Call the function under test
	result, err := FetchInfo()

	// Assert no error (since we're not validating the IP format)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Assert the invalid plain text response is returned as-is
	if result != invalidResponse {
		t.Errorf("Expected %s, got %s", invalidResponse, result)
	}
}
