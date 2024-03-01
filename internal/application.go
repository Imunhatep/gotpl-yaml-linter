package internal

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func NewApp() *cli.App {
	cmd := cli.NewApp()
	cmd.EnableBashCompletion = true
	cmd.Name = "yamltpl_linter"
	cmd.Usage = "GoLang template for yaml formatting and linting tool"
	cmd.UsageText = "bin/yamltpl_linter_{os}-{arch} [command] [subcommand] [command options]"
	cmd.Description = "https://github.com/imunhatep/gotpl-yaml-linter/README.md"
	cmd.Before = func(ctx *cli.Context) error {
		verbose := ctx.Int("verbose")
		setLogLevel(verbose)

		return nil
	}
	cmd.Flags = []cli.Flag{
		&cli.IntFlag{
			Name:     "verbose",
			Aliases:  []string{"v"},
			EnvVars:  []string{"APP_DEBUG"},
			Usage:    "Log verbosity",
			Required: false,
			Value:    3,
		},
	}

	return cmd
}

func setLogLevel(level int) {
	switch level {
	case 0:
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case 1:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case 2:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case 3:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case 4:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	}

	log.Debug().Msgf("logging level: %s", zerolog.GlobalLevel().String())
}
