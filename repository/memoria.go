package repository

import (
	"mini-proyecto-1/domain"
	"sync"
)

type RepositorioEnMemoria struct {
	db map[string]*domain.Tarea
	mu sync.RWMutex
}

func NuevoRepositorioEnMemoria() *RepositorioEnMemoria {
	return &RepositorioEnMemoria{
		db: make(map[string]*domain.Tarea),
	}
}

func (r *RepositorioEnMemoria) Crear(tarea *domain.Tarea) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, existe := r.db[tarea.ID]; existe {
		return domain.ErrTareaYaExiste
	}

	r.db[tarea.ID] = tarea
	return nil
}

func (r *RepositorioEnMemoria) ObtenerPorID(id string) (*domain.Tarea, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	tarea, existe := r.db[id]
	if !existe {
		return nil, domain.ErrTareaNoEncontrada
	}
	return tarea, nil
}

func (r *RepositorioEnMemoria) ListarTodas() ([]*domain.Tarea, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	tareas := make([]*domain.Tarea, 0, len(r.db))
	for _, tarea := range r.db {
		tareas = append(tareas, tarea)
	}
	return tareas, nil
}

func (r *RepositorioEnMemoria) MarcarCompletada(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	tarea, ok := r.db[id]
	if !ok {
		return domain.ErrTareaNoEncontrada
	}
	tarea.Completada = true
	return nil
}
