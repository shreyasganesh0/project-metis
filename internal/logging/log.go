package logging

import (
    "os"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)

func Init() {

    output := zerolog.ConsoleWriter{Out: os.Stderr
    log.Logger = log.Output(output)

    zerolog.SetGlobalLevel(zerolog.InfoLevel)

    log.Info().Msg("Logger Initialzied.");
}
