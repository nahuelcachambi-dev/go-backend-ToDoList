package domain

import (
	"errors"
	"strings"
)

type Tarea struct {
	ID          string
	Titulo      string
	Descripcion string
	Completada  bool
}

func NuevaTarea(id, titulo, descripcion string) (*Tarea, error) {
	if strings.TrimSpace(titulo) == "" {
		return nil, errors.New("el titulo de la tarea no puede estar vacio")
	}

	if len(titulo) < 3 {
		return nil, errors.New("el titulo deebe tener al menos 3 caracteres")
	}

	return &Tarea{
		ID:          id,
		Titulo:      titulo,
		Descripcion: descripcion,
		Completada:  false,
	}, nil
}

type RepositorioTareas interface {
	Crear(tarea *Tarea) error
	ObtenerPorID(id string) (*Tarea, error)
	ListarTodas() ([]*Tarea, error)
	MarcarCompletada(id string) error
}

var (
	ErrTareaNoEncontrada = errors.New("la tarea no existe")
	ErrTareaYaExiste     = errors.New("la tarea ya existe")
)
