// cmd/main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"doctor-reservation-system/config"
	"doctor-reservation-system/internal/auth"
	"doctor-reservation-system/internal/doctors"
	"doctor-reservation-system/internal/handlers"
	"doctor-reservation-system/internal/notifications"
	"doctor-reservation-system/internal/reservations"
	"doctor-reservation-system/pkg/database"
	m "doctor-reservation-system/pkg/database/migrations"
	"doctor-reservation-system/pkg/middleware"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Connect to database
	db, err := database.NewPostgresConnection(cfg.Database)
	if err != nil {
		log.Fatalf("Cannot connect to database: %v", err)
	}
	defer db.Close()

	// Initialize repositories
	authRepo := auth.NewRepository(db)
	doctorRepo := doctors.NewRepository(db)
	reservationRepo := reservations.NewRepository(db)

	// Initialize services
	authService := auth.NewService(authRepo)
	doctorService := doctors.NewService(doctorRepo)
	reservationService := reservations.NewService(reservationRepo, authRepo)
	notificationService := notifications.NewService()

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	doctorHandler := handlers.NewDoctorHandler(doctorService)
	// In main.go
	reservationHandler := handlers.NewReservationHandler(
		reservationService,
		notificationService,
		db, // Pass the database connection
	)
	m.RunMigrations()
	// Create router
	r := mux.NewRouter()

	// Public routes
	r.HandleFunc("/register", authHandler.Register).Methods("POST")
	r.HandleFunc("/login", authHandler.Login).Methods("POST")

	// Protected routes
	protectedRoutes := r.PathPrefix("/api").Subrouter()
	protectedRoutes.Use(middleware.AuthMiddleware(authService))
	{
		// Doctor routes
		protectedRoutes.HandleFunc("/doctors", doctorHandler.Search).Methods("GET")

		// Reservation routes
		protectedRoutes.HandleFunc("/reservations", reservationHandler.Create).Methods("POST")
		protectedRoutes.HandleFunc("/reservations", reservationHandler.List).Methods("GET")
	}

	// Start server
	port := cfg.Server.Port
	log.Printf("Server starting on :%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
