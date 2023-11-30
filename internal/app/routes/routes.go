package routes

import (
	"internal/app/routes/handlers"

	"github.com/gorilla/mux"
)

// Set up your router
var router = mux.NewRouter()

// Define your API routes
func InitializeRoutes() {
	router.HandleFunc("/api/meditation", handlers.GetMeditation).Methods("GET")
	router.HandleFunc("/api/meditation/{id}", handlers.GetMeditationByID).Methods("GET")
	router.HandleFunc("/api/meditation", handlers.CreateMeditation).Methods("POST")
	router.HandleFunc("/api/meditation/{id}", handlers.UpdateMeditation).Methods("PUT")
	router.HandleFunc("/api/meditation/{id}", handlers.DeleteMeditation).Methods("DELETE")

	// Add more routes as needed for your application
}

// GetRouter returns the configured router
func GetRouter() *mux.Router {
	return router
}
