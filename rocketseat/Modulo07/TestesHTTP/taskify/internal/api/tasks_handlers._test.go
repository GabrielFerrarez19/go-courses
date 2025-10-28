package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleCreateTasck(t *testing.T) {
	api := Application{}

	payload := map[string]any{
		"title":       "Learn TDD",
		"description": "Get hands-on exp with TDD in Go!",
		"priority":    8000,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatal("Failed to parse our request payload")
	}

	req := httptest.NewRequest("POST", "/api/v1/tasks", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(api.HandleCreateTask)

	handler.ServeHTTP(rec, req)

	t.Logf("Rec body %s\n", rec.Body.Bytes())

	if rec.Code != http.StatusCreated {
		t.Errorf("Status code differs; got %d | want %d", rec.Code, http.StatusCreated)
	}

	var ResBody map[string]any

	err = json.Unmarshal(rec.Body.Bytes(), &ResBody)
	if err != nil {
		t.Fatalf("Failed to parse response body: %s\n", err.Error())
	}

	if ResBody["title"] != payload["title"] {
		t.Errorf("title differs; got: %q | want %q", ResBody["title"], payload["title"])
	}
}
