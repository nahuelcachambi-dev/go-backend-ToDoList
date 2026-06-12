package handler

import (
	"encoding/json"
	"fmt"
	"mini-proyecto-1/domain"
	"net/http"
)

type TareaHandler struct {
	repo domain.RepositorioTareas
}

func NuevoTareaHandler(repo domain.RepositorioTareas) *TareaHandler {
	return &TareaHandler{repo: repo}
}

func (h *TareaHandler) CrearTarea(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		ID          string `json:"id"`
		Titulo      string `json:"titulo"`
		Descripcion string `json:"descripcion"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	tarea, err := domain.NuevaTarea(req.ID, req.Titulo, req.Descripcion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.repo.Crear(tarea); err != nil {
		if err == domain.ErrTareaYaExiste {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, "Error interno", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Conten-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tarea)
}

func (h *TareaHandler) ObtenerTarea(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}

	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "El parametro 'id' es requerido", http.StatusBadRequest)
		return
	}

	tarea, err := h.repo.ObtenerPorID(id)
	if err != nil {
		if err == domain.ErrTareaNoEncontrada {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tarea)
}

func (h *TareaHandler) ListarTodas(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Metodo invalido", http.StatusMethodNotAllowed)
		return
	}

	tareas, err := h.repo.ListarTodas()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tareas)
}

func (h *TareaHandler) MarcarCompletada(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "El parametro 'id' es requerido", http.StatusBadRequest)
		return
	}
	err := h.repo.MarcarCompletada(id)
	if err != nil {
		if err == domain.ErrTareaNoEncontrada {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"mensaje": fmt.Sprintf("Tarea %s marcada como completada", id),
	})
}
