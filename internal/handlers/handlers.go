package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Sample data structure for a meditation
type Meditation struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Define a slice to store meditation data (simulated database)
var meditations []Meditation

// Handler for getting all meditations
func GetMeditations(w http.ResponseWriter, r *http.Request) {
	// Serialize meditations to JSON and send as response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(meditations)
}

// Handler for getting a specific meditation by ID
func GetMeditationByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	meditationID := vars["id"]

	// Find the meditation with the given ID (simplified for demonstration)
	var foundMeditation Meditation
	for _, meditation := range meditations {
		if meditation.ID == meditationID {
			foundMeditation = meditation
			break
		}
	}

	// Serialize the found meditation to JSON and send as response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foundMeditation)
}

// Handler for creating a new meditation
func CreateMeditation(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to extract meditation data (simplified for demonstration)
	var newMeditation Meditation
	_ = json.NewDecoder(r.Body).Decode(&newMeditation)

	// Generate a unique ID for the new meditation (you may use a UUID library)
	newMeditation.ID = "1" // Simplified ID generation

	// Append the new meditation to the slice (simulated database)
	meditations = append(meditations, newMeditation)

	// Serialize the new meditation to JSON and send as response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newMeditation)
}

// Handler for updating an existing meditation
func UpdateMeditation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	meditationID := vars["id"]

	// Parse the request body to extract updated meditation data (simplified for demonstration)
	var updatedMeditation Meditation
	_ = json.NewDecoder(r.Body).Decode(&updatedMeditation)

	// Find the index of the meditation with the given ID (simplified for demonstration)
	var meditationIndex int
	for i, meditation := range meditations {
		if meditation.ID == meditationID {
			meditationIndex = i
			break
		}
	}

	// Update the meditation at the specified index (simulated database)
	meditations[meditationIndex] = updatedMeditation

	// Serialize the updated meditation to JSON and send as response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedMeditation)
}

// Handler for deleting a meditation by ID
func DeleteMeditation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	meditationID := vars["id"]

	// Find the index of the meditation with the given ID (simplified for demonstration)
	var meditationIndex int
	for i, meditation := range meditations {
		if meditation.ID == meditationID {
			meditationIndex = i
			break
		}
	}

	// Remove the meditation from the slice (simulated database)
	if meditationIndex >= 0 {
		meditations = append(meditations[:meditationIndex], meditations[meditationIndex+1:]...)
	}

	// Respond with a success message (or appropriate response)
	w.WriteHeader(http.StatusNoContent)
}
