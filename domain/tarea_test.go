package domain

import "testing"

func TestNuevaTarea(t *testing.T) {
	tests := []struct {
		nombre      string
		titulo      string
		eperarError bool
	}{
		{"Titulo valido", "Aprender Go", false},
		{"Titulo vacio", "", true},
	}

	for _, tc := range tests {
		t.Run(tc.nombre, func(t *testing.T) {
			tarea, err := NuevaTarea("1", tc.titulo, "ddddd")
			if (err != nil) != tc.eperarError {
				t.Errorf("Esperaba error: %v, obtuve: %v", tc.eperarError, err)
			}

			if !tc.eperarError && tarea.Titulo != tc.titulo {
				t.Errorf("Esperaba titulo: %s, obtube %s", tc.titulo, tarea.Titulo)
			}
		})
	}
}
