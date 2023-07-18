package main

import (
	"fmt"
	"go-app/user-service/db"
	"go-app/user-service/endpoint"
	"go-app/user-service/server"
	"go-app/user-service/service"
	"go-app/user-service/utils"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	logger := *utils.InitLogger()
	repo := db.NewTransactionRepo(logger)

	logger.Log().Info("Starting Transaction Service Application")

	// Init Servcie
	srv := service.NewTransactionService(repo, logger)

	// Init Endpoint
	endpoints := endpoint.CreateEndpoints(srv, logger)

	errCh := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errCh <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log().Info("Listening on port : " + "8084")
		handler := server.NewGorillaMuxServer(endpoints)
		errCh <- http.ListenAndServe(":8084", handler)
	}()

	logger.Log().Info("Exit", <-errCh)

}
