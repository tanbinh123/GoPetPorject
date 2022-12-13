package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Brigant/GoPetPorject/app/handlers"
	"github.com/Brigant/GoPetPorject/app/repositories/pg"
	"github.com/Brigant/GoPetPorject/app/usecases"
	"github.com/Brigant/GoPetPorject/configs"
)

type Server struct {
	httpServer *http.Server
}

func main() {
	if err := SetupAndRun(); err != nil {
		log.Fatalf("error while SetupAndRun server: %s", err.Error())
	}

}

// SetupAndRun function binds all layers together and starts the server.
func SetupAndRun() error {
	cfg, err := configs.InitConfig()
	if err != nil {
		return fmt.Errorf("cannot read config: %w", err)
	}

	db, err := pg.NewPostgresDB(cfg)
	if err != nil {
		return fmt.Errorf("error while creating connection to database: %w", err)
	}

	repo := pg.NewRepository(db)

	usecase := usecases.NewUsecase(repo)

	handlers := handlers.NewHandler(usecase)

	routes := handlers.InitRouter(cfg.Server.Mode)

	server := new(Server)
	if err := server.Run(cfg.Server.Port, routes); err != nil {
		return fmt.Errorf("cannot run server: %w", err)
	}

	return nil
}

// Run function runs the http server.
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	return s.httpServer.ListenAndServe()
}

// Shutdown function stops the http server.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
