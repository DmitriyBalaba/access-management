package cmd

import (
	"access-management/pkg/config"
	auth "access-management/pkg/domain/auth/registration"
	company "access-management/pkg/domain/company/registration"
	user "access-management/pkg/domain/user/registration"
	"access-management/pkg/server"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCommand.AddCommand(
		&cobra.Command{
			Use:   "serve",
			Short: "Starts API server",
			Run: func(cmd *cobra.Command, args []string) {
				if err := serve(); err != nil {
					log.Fatal().Msg("Serve command failed: " + err.Error())
					return
				}
			},
		})
}

func serve() error {
	conf, err := config.NewConfig(configPath)
	if err != nil {
		return err
	}

	user.HttpRoutes(conf)
	company.HttpRoutes(conf)
	auth.HttpRoutes(conf)

	newServer := server.NewServer(conf)
	newServer.Run()
	return nil
}
