// main_test.go
package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

// Test for the homePage handler
func TestHomePage(t *testing.T) {
    req, err := http.NewRequest("GET", "/home", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(homePage)
    handler.ServeHTTP(rr, req)

    // Check if the response status code is 200 OK
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    // Check if the response content-type is "text/html; charset=utf-8"
    expected := "text/html; charset=utf-8"
    if contentType := rr.Header().Get("Content-Type"); contentType != expected {
        t.Errorf("handler returned wrong content-type: got %v want %v", contentType, expected)
    }
}

// Additional tests for other pages (courses, about, etc.) would go here
