package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestOfferFood(t *testing.T) {
	form := url.Values{}
	form.Add("name", "Anna")
	form.Add("address", "123 Main St")
	form.Add("food", "Soup")
	form.Add("portions", "3")

	req := httptest.NewRequest(
		http.MethodPost,
		"/offer",
		strings.NewReader(form.Encode()),
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()

	OfferFood(rr, req)

	res := rr.Result()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}

	bodyStr := string(body)

	if !strings.Contains(bodyStr, "POST request successful") {
		t.Fatalf("unexpected response: %s", bodyStr)
	}

	if !strings.Contains(bodyStr, "Anna") {
		t.Fatalf("name not found in response: %s", bodyStr)
	}
}
