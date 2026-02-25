package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/lhbelfanti/ditto/database"
	"github.com/lhbelfanti/ditto/log"
	"github.com/lhbelfanti/ditto/setup"
	"github.com/rs/zerolog"

	"github.com/lhbelfanti/go-project-template/cmd/api/example"
)

var prodEnv bool

func init() {
	flag.BoolVar(&prodEnv, "prod", false, "Run in production environment")
	flag.Parse()
}

func main() {
	/* --- Dependencies --- */
	ctx := context.Background()

	logLevel := zerolog.DebugLevel
	if prodEnv {
		logLevel = zerolog.InfoLevel
	}

	log.NewCustomLogger(os.Stdout, logLevel)

	// Database
	pg := setup.Init(database.InitPostgres())
	defer pg.Close()
	db := pg.Database()

	// Services

	// GET /example/v1 dependencies
	collectExampleDAORows := database.MakeCollectRows[example.DAO](nil)
	selectAllExamples := example.MakeSelectAll(db, collectExampleDAORows)

	/* --- Router --- */
	log.Info(ctx, "Initializing router...")
	router := http.NewServeMux()
	router.HandleFunc("GET /example/v1", example.SelectAllHandlerV1(selectAllExamples))
	log.Info(ctx, "Router initialized!")

	/* --- Server --- */
	port := fmt.Sprintf(":%s", os.Getenv("API_PORT"))
	log.Info(ctx, fmt.Sprintf("github.com/lhbelfanti/go-project-template server is ready to receive request on port %s", port))
	if err := http.ListenAndServe(port, router); err != nil {
		log.Error(ctx, fmt.Sprintf("Server failed to start: %v", err))
		os.Exit(1)
	}
}
