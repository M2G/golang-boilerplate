package handlers

import (
  "net/http"
  "net/http/httptest"
  "strings"
  "testing"
)

func TestBarHandler(t *testing.T) {
  req := httptest.NewRequest(http.MethodGet, "/bar", nil)
  w := httptest.NewRecorder()

  BarHandler(w, req)

  res := w.Result()

  if res.StatusCode != http.StatusOK {
    t.Fatalf("expected status 200, got %d", res.StatusCode)
  }

  body := w.Body.String()
  if !strings.Contains(body, "Hello World") {
    t.Fatalf("unexpected body: %s", body)
  }
}

func TestBarByIDHandler(t *testing.T) {
  req := httptest.NewRequest(http.MethodGet, "/bar/123", nil)

  req.SetPathValue("id", "123")

  w := httptest.NewRecorder()

  BarByIDHandler(w, req)

  res := w.Result()

  if res.StatusCode != http.StatusOK {
    t.Fatalf("expected status 200, got %d", res.StatusCode)
  }

  body := w.Body.String()
  if !strings.Contains(body, "123") {
    t.Fatalf("unexpected body: %s", body)
  }
}
