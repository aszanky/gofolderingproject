package cmd

import (
	"github.com/aszanky/gofolderingproject/cmd/grpc"
	"github.com/aszanky/gofolderingproject/cmd/rest"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	rootCmd = cobra.Command{
		Use: "project-structure",
	}

	httpPublicCmd = &cobra.Command{
		Use: "rest-public",
		Run: func(cmd *cobra.Command, args []string) {
			if err := rest.ServeREST(); err != nil {
				log.Fatal().Err(err).Msg("error while starting http server")
			}
		},
	}

	grpcCmd = &cobra.Command{
		Use: "grpc",
		Run: func(cmd *cobra.Command, args []string) {
			if err := grpc.ServeGRPC(); err != nil {
				log.Fatal().Err(err).Msg("error while starting grpc server")
			}
		},
	}
)

func Execute() {
	rootCmd.AddCommand(httpPublicCmd)
	rootCmd.AddCommand(grpcCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("failed to execute command")
	}
}
