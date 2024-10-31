package main

import (
	"eco-smart-api/auth"
	"eco-smart-api/config"
	"eco-smart-api/controller"
	"eco-smart-api/core"
	"eco-smart-api/handler"
	"eco-smart-api/repository"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
}

func NewApp(cfg *config.Config) *App {
	core.InitializeDatabase("user=" + cfg.Database.User + " password=" + cfg.Database.Password + " dbname=" + cfg.Database.Name + " sslmode=disable")

	app := &App{
		Router: mux.NewRouter(),
	}
	app.Router.Use(auth.JSONMiddleware)
	app.Router.Use(auth.LoggingMiddleware)

	authRouter := app.Router.PathPrefix("/auth/v1").Subrouter()
	v1 := app.Router.PathPrefix("/api/v1").Subrouter()

	v1.Use(auth.AuthecationMiddleware)

	sessionRepo := repository.NewSessionRepository(core.GetDB())
	sessionController := controller.NewSessionController(sessionRepo)
	sessionHandler := handler.NewSessionHandler(sessionController)
	sessionHandler.RegisterRoutes(authRouter)

	userRepo := repository.NewUserRepository(core.GetDB())
	userController := controller.NewUserController(userRepo)
	userHandler := handler.NewUserHandler(userController)
	userHandler.RegisterRoutes(v1)

	collectionPointRepo := repository.NewCollectionPointRepository(core.GetDB())
	collectionPointController := controller.NewCollectionPointController(collectionPointRepo)
	collectionPointHandler := handler.NewCollectionPointHandler(collectionPointController)
	collectionPointHandler.RegisterRoutes(v1)

	return app
}

func main() {
	cfg := config.LoadConfig()

	app := NewApp(cfg)

	addr := cfg.ServerAddress
	// CORS options
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	log.Printf("Server listening on %s", addr)
	if err := http.ListenAndServe(addr, handlers.CORS(originsOk, headersOk, methodsOk)(app.Router)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
