package cmd

import (
	"os"
	"path"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const (
	DEBUG       = "debug"
	DEBUG_SHORT = "d"
	DEBUG_USE   = "server debug mode"

	CONFIG       = "config"
	CONFIG_SHORT = "c"
	CONFIG_USE   = "config file path"
)

var (
	debugMode  = false
	configPath = path.Base(os.Args[0]) + ".yaml"
)

var rootCommand = &cobra.Command{
	Use:   "Access management",
	Short: "Access Management (c)",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// showing debug messages in debug mode only
		if debugMode {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		} else {
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}
		log.Debug().Msg("Access management is in debug mode: you are able to see debug messages")
	},
}

func init() {
	// configure global app flags
	rootCommand.PersistentFlags().StringVarP(&configPath, CONFIG, CONFIG_SHORT, configPath, CONFIG_USE)
	rootCommand.PersistentFlags().BoolVarP(&debugMode, DEBUG, DEBUG_SHORT, debugMode, DEBUG_USE)
}

// Execute runs command configurations and exits on its finish
func Execute() {
	if err := rootCommand.Execute(); err != nil {
		log.Fatal().Msg("Execution failed: " + err.Error())
		os.Exit(-1)
	}
	os.Exit(0)
}
