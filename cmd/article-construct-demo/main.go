package main

import (
	"fmt"
	"github.com/user/article-construct-demo/internal/database"
	"github.com/user/article-construct-demo/internal/repository"
	"github.com/user/article-construct-demo/internal/service"
	"net/http"
	"os"

	"github.com/user/article-construct-demo/internal/log"
	"go.uber.org/zap"

	_ "go.uber.org/automaxprocs"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "an error occurred: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	logger, err := log.NewAtLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		return err
	}

	defer func() {
		err = logger.Sync()
	}()

	// create a database connection
	db := database.Connect()

	// create a repository
	repository := repository.NewRepository(db)

	// create an http server
	service := service.NewConstructHandler(repository)

	http.Handle("/construct", http.HandlerFunc(service.GetConstructForIan))
	http.Handle("/item", http.HandlerFunc(service.GetItemForIan))
	http.Handle("/case", http.HandlerFunc(service.GetCaseForIan))
	http.Handle("/variant", http.HandlerFunc(service.GetVariantForIan))
	err = http.ListenAndServe(":8181", nil)
	if err != nil {
		logger.Error("Error occurred")
	}

	logger.Info("Hello world!", zap.String("location", "world"))

	return err
}
