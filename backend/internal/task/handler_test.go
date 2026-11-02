package task

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateReturns201WithLocation(t *testing.T) {
	h := NewHandler(NewStore())

	body := strings.NewReader(`{"title":"estudar go"}`)
	req := httptest.NewRequest(http.MethodPost, "/tasks", body)
	rec := httptest.NewRecorder()

	h.Create(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("status = %d, esperado %d", rec.Code, http.StatusCreated)
	}
	if got := rec.Header().Get("Location"); got != "/tasks/1" {
		t.Errorf("location = %q, esperado %q", got, "/tasks/1")
	}
}
