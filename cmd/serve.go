package cmd

import (
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
	return nil
}
