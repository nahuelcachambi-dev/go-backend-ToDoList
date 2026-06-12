package repository

import (
	"mini-proyecto-1/domain"
	"testing"
)

// Test para crear tarea
func TestRepositorio_Crear(t *testing.T) {
	tests := []struct {
		nombre       string
		nuevaTarea   *domain.Tarea
		esperarError error
	}{
		{
			nombre: "Creacion exitosa de tarea nueva",
			nuevaTarea: &domain.Tarea{
				ID:     "1",
				Titulo: "Aprender Tests",
			},
			esperarError: nil,
		},
		{
			nombre: "Fallo crear tarea con ID existente",
			nuevaTarea: &domain.Tarea{
				ID:     "1",
				Titulo: "Aprender Tests",
			},
			esperarError: domain.ErrTareaYaExiste,
		},
	}
	for _, tc := range tests {
		t.Run(tc.nombre, func(t *testing.T) {
			repo := NuevoRepositorioEnMemoria()

			if tc.esperarError == domain.ErrTareaYaExiste {
				repo.Crear(&domain.Tarea{ID: "1", Titulo: "Tarea Base"})
			}

			err := repo.Crear(tc.nuevaTarea)
			if err != tc.esperarError {
				t.Errorf("Esperado error: %v, obtenido: %v", tc.esperarError, err)
			}
		})
	}
}

// Test para obtener ID

func TestRepositorio_ObtenerPorId(t *testing.T) {
	repo := NuevoRepositorioEnMemoria()
	tareaBase := &domain.Tarea{ID: "10", Titulo: "Tarea de prueba"}
	repo.Crear(tareaBase)
	tests := []struct {
		nombre       string
		ID           string
		esperarError error
	}{
		{
			nombre:       "Tarea encontrada",
			ID:           "10",
			esperarError: nil,
		},
		{
			nombre:       "Tarea no encontrada",
			ID:           "99",
			esperarError: domain.ErrTareaNoEncontrada,
		},
	}

	for _, tc := range tests {
		t.Run(tc.nombre, func(t *testing.T) {
			tarea, err := repo.ObtenerPorID(tc.ID)
			if err != tc.esperarError {
				t.Errorf("Esperar error: %v, obtenido: %v", tc.esperarError, err)
			}
			if tc.esperarError == nil {
				if tarea.ID != tc.ID {
					t.Errorf("Esperado ID: %s, obtenido %s", tc.ID, tarea.ID)
				}

			}
		})
	}

}

// Test para marcar completada

func TestRepositorio_MarcarCompletada(t *testing.T) {
	repo := NuevoRepositorioEnMemoria()
	repo.Crear(&domain.Tarea{ID: "20", Titulo: "Tarea pendiente", Completada: false})
	tests := []struct {
		nombre       string
		idBuscar     string
		errorEsperar error
	}{
		{
			nombre:       "Tarea existente",
			idBuscar:     "20",
			errorEsperar: nil,
		},
		{
			nombre:       "Tarea inexistente",
			idBuscar:     "99",
			errorEsperar: domain.ErrTareaNoEncontrada,
		},
	}
	for _, tc := range tests {
		t.Run(tc.nombre, func(t *testing.T) {
			err := repo.MarcarCompletada(tc.idBuscar)
			if err != tc.errorEsperar {
				t.Errorf("Esperar Error: %v, Obtenido: %v", tc.errorEsperar.Error(), err)
			}

			if err == nil {
				tarea, _ := repo.ObtenerPorID(tc.idBuscar)
				if !tarea.Completada {
					t.Error("Se esperaba que la tarea estuviera marcada como completada (true)")
				}
			}
		})
	}
}
