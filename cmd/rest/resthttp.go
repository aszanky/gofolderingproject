package rest

import (
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/aszanky/gofolderingproject/config"
)

func ServeREST() error {
	// Dependency Injection
	cfg, err := config.Load("./config/.env")
	if err != nil {
		log.Fatal().Err(err).Msg("error while starting http server")
	}

	// Run Server
	log.Print("Starting server on ", cfg.PORT)
	err = http.ListenAndServe(cfg.PORT, nil)
	if err != nil {
		return err
	}

	return nil
}
