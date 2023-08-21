package grpc

import (
	"fmt"
	"net"

	"github.com/aszanky/gofolderingproject/config"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func ServeGRPC() error {
	cfg, err := config.Load("./config/.env")
	if err != nil {
		log.Fatal().Err(err).Msg("error while starting http server")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0%s", cfg.PORT_GRPC))
	if err != nil {
		log.Fatal().Err(err).Msg("Listening HTTP was error")
	}
	// Dependency Injection

	//grpc
	srv := grpc.NewServer()

	// create health server
	hs := health.NewServer()
	healthpb.RegisterHealthServer(srv, hs)
	if err := srv.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("Starting grpc server has failed")
		return err
	}

	return nil
}
