package rest

import (
	"context"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/aszanky/gofolderingproject/config"
	"github.com/aszanky/gofolderingproject/internal/handler/resthandler"
	"github.com/aszanky/gofolderingproject/internal/repository"
	"github.com/aszanky/gofolderingproject/internal/usecase"
	"github.com/aszanky/gofolderingproject/pkg/db"
	"github.com/aszanky/gofolderingproject/pkg/logging"
	"github.com/aszanky/gofolderingproject/pkg/pgxdb"
	"github.com/aszanky/gofolderingproject/pkg/tracing"
)

func ServeREST() error {
	// Dependency Injection

	logger := logging.NewLogger()
	logger.Info("Logger initialized")

	tp, err := tracing.NewTracerProvider()
	if err != nil {
		logger.Error("failed to initialize tracer", "error", err)
		os.Exit(1)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			logger.Error("failed to shutdown tracer", "error", err)
		}
	}()
	logger.Info("Tracer initialized")

	cfg, err := config.Load("./config/.env")
	if err != nil {
		log.Fatal().Err(err).Msg("error while starting http server")
	}

	//SQLX
	sqlxDB, err := db.NewDatabase()
	if err != nil {
		log.Fatal().Err(err).Msg("error database")
	}
	defer sqlxDB.Close()

	//PGX
	dbpool, err := pgxdb.NewConnection(cfg.DatabaseURL)
	if err != nil {
		logger.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	logger.Info("successfully connected to PostgreSQL")

	allRepository := repository.NewRepository(repository.NewRepositoryParam{
		DB:    sqlxDB,
		PGXDB: dbpool,
	})
	allUsecase := usecase.NewUsecase(allRepository)
	allHandler := resthandler.NewHandler(allUsecase, cfg)

	// Run Server
	router := resthandler.SetupRouter(allHandler, logger)

	logger.Info("Starting server on port %s", cfg.SERVER_PORT)
	if err := router.Run(cfg.SERVER_PORT); err != nil {
		logger.Error("failed to start server", "error", err)
		os.Exit(1)
	}

	return nil
}
