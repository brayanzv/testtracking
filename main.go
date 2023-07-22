package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/endpoint", locationHandler)

	log.Println("Server started on localhost:8000")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func locationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decodificar el JSON del cuerpo de la solicitud en una estructura Location
	var loc Location
	err := json.NewDecoder(r.Body).Decode(&loc)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Utilizar la ubicación recibida en la estructura Location
	response := fmt.Sprintf("Ubicación recibida: Latitud: %f, Longitud: %f", loc.Latitude, loc.Longitude)
	fmt.Fprintf(w, response)
	fmt.Printf("Ubicación recibida: Latitud: %f, Longitud: %f", loc.Latitude, loc.Longitude)
}
