package domain

import "errors"

type Tarea struct {
	ID          string
	Titulo      string
	Descripcion string
	Completada  bool
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
