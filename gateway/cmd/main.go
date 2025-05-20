package main

import (
	_ "gateway/docs"
	"gateway/internal/handlers"
	"gateway/middleware"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"path/filepath"
)

// @title APP FOR MODELS API
// @version 1.0
// @description GATEWAY for models site application

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	router := mux.NewRouter()

	// Создаем подмаршрутизатор для API
	apiRouter := router.PathPrefix("/api").Subrouter()

	// Инициализация обработчиков
	h := handlers.NewHandlers()

	// Добавляем Swagger UI
	router.PathPrefix("/swagger").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // Путь к swagger.json
		httpSwagger.DeepLinking(true),
	)).Methods("GET")

	// Добавляем статический обработчик для swagger.json
	swaggerDir := filepath.Join("..", "docs") // Путь к папке docs
	fs := http.FileServer(http.Dir(swaggerDir))
	router.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", fs))

	authRouter := apiRouter.PathPrefix("/auth").Subrouter()

	// Маршруты для HTTP API Gateway
	authRouter.HandleFunc("/register", h.Register).Methods("POST")
	authRouter.HandleFunc("/login", h.Login).Methods("POST")
	authRouter.HandleFunc("/is-admin/{user_id}", h.IsAdmin).Methods("GET")

	// Защищенные эндпоинты
	protectedAPIRouter := apiRouter.PathPrefix("/auth").Subrouter()
	protectedAPIRouter.Use(middleware.AuthMiddleware)
	protectedAPIRouter.HandleFunc("/refresh-token", h.RefreshToken).Methods("POST")
	protectedAPIRouter.HandleFunc("/logout", h.Logout).Methods("GET")
	protectedAPIRouter.HandleFunc("/session", h.Session).Methods("GET")

	protectedProfileRouter := apiRouter.PathPrefix("/profile").Subrouter()
	protectedProfileRouter.Use(middleware.AuthMiddleware)
	protectedProfileRouter.HandleFunc("", h.GetProfile).Methods("GET")
	protectedProfileRouter.HandleFunc("", h.UpdateProfile).Methods("PUT")
	protectedProfileRouter.HandleFunc("", h.DeleteProfile).Methods("DELETE")

	catalogRouter := apiRouter.PathPrefix("/catalog").Subrouter()
	catalogRouter.HandleFunc("", h.GetCatalog).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "PUT", "DELETE", "UPDATE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Cookie"},
	})

	handler := c.Handler(router)

	// Запуск сервера
	log.Println("Starting API Gateway on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
