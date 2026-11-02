package task

import (
	"encoding/json"
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

func TestGetUnknownTaskReturns404JSON(t *testing.T) {
	h := NewHandler(NewStore())
	req := httptest.NewRequest(http.MethodGet, "/tasks/99", nil)
	req.SetPathValue("id", "99")
	rec := httptest.NewRecorder()

	h.GetTask(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("status = %d, esperado %d", rec.Code, http.StatusNotFound)
	}

	if rec.Header().Get("Content-Type") != "application/json" {
		t.Errorf("status = %q, esperado %q", rec.Header().Get("Content-Type"), "application/json")
	}

	var resp errorResponse
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("corpo não é JSON válido: %v", err)
	}

	if resp.Error != "tarefa não encontrada" {
		t.Errorf("error: %q, esperado %q", resp.Error, "tarefa não encontrada")
	}
}
