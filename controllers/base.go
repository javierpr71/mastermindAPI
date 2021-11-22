package controllers

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	gohandlers "github.com/gorilla/handlers"
	"github.com/javierpr71/mastermind/driver"
	ph "github.com/javierpr71/mastermind/handler"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	logger       *log.Logger
	connection   *driver.DB
	pGameHandler *ph.Game
}

func NewServer(l *log.Logger) *Server {
	return &Server{l, nil, nil}
}

func (server *Server) Initialize() error {
	var err error
	server.connection, err = driver.ConnectStorage(os.Getenv("REDIS"), "", 0)
	if err != nil {
		return err
	}
	server.logger.Info("Connected to dabasase.")
	server.pGameHandler = ph.NewMasterMindHandler(server.logger, server.connection)
	return nil
}

func (server *Server) Run(addr string) {

	r := server.initializeRoutes()

	// CORS
	cors := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	srv := &http.Server{
		Handler:      cors(r),
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		server.logger.Infof("Starting server on port %s", addr)

		err := srv.ListenAndServe()
		if err != nil {
			server.logger.Infof("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// Gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Waiting for signal
	sig := <-c
	log.Println("Signal recived:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
}
