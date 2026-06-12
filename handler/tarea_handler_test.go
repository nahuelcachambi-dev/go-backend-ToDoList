package handler

import (
	"bytes"
	"mini-proyecto-1/domain"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockRepositorio struct {
	crearErr error
}

func (m *MockRepositorio) Crear(tarea *domain.Tarea) error {
	return m.crearErr
}

func (m *MockRepositorio) ObtenerPorID(id string) (*domain.Tarea, error) { return nil, nil }
func (m *MockRepositorio) ListarTodas() ([]*domain.Tarea, error)         { return nil, nil }
func (m MockRepositorio) MarcarCompletada(id string) error               { return nil }

func TestCrearTareaHandler(t *testing.T) {
	tests := []struct {
		nombre         string
		bodyJSON       string
		mockCrearErr   error
		expectedStatus int
	}{
		{
			nombre:         "Creacion exitosa",
			bodyJSON:       `{"id":"1", "titulo":"Tarea valida", "descripcion":"test"}`,
			mockCrearErr:   nil,
			expectedStatus: http.StatusCreated,
		},
		{
			nombre:         "Tarea ya existe (Conflicto)",
			bodyJSON:       `{"id":"1", "titulo":"Tarea valida", "descripcion":"test"}`,
			mockCrearErr:   domain.ErrTareaYaExiste,
			expectedStatus: http.StatusConflict,
		},
		{
			nombre:         "Json invalido",
			bodyJSON:       `{"id":1, "titulo":"Tarea valida", "descripcion":"test"}`,
			mockCrearErr:   nil,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range tests {
		t.Run(tc.nombre, func(t *testing.T) {
			mockRepo := &MockRepositorio{crearErr: tc.mockCrearErr}

			handler := NuevoTareaHandler(mockRepo)
			req := httptest.NewRequest(http.MethodPost, "/api/tareas", bytes.NewBufferString(tc.bodyJSON))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			handler.CrearTarea(w, req)
			if w.Code != tc.expectedStatus {
				t.Errorf("Esperando status %d, obtenido %d, Body: %s", tc.expectedStatus, w.Code, w.Body.String())
			}
		})
	}
}
