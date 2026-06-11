package main

import (
	"fmt"
	"log"
	"mini-proyecto-1/handler"
	"mini-proyecto-1/repository"
	"net/http"
)

func main() {
	repo := repository.NuevoRepositorioEnMemoria()

	tareaHandler := handler.NuevoTareaHandler(repo)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/tareas", tareaHandler.CrearTarea)
	mux.HandleFunc("GET /api/tareas/{id}", tareaHandler.ObtenerTarea)
	mux.HandleFunc("GET /api/tareas", tareaHandler.ListarTodas)
	mux.HandleFunc("PUT /api/tareas/{id}", tareaHandler.MarcarCompletada)
	fmt.Println("Servidor corriendo en http://localhost:8080")
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Fallo crítico: %v", err)
	}
}
