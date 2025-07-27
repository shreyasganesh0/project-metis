package logging

import (
    "os"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
    "github.com/spf13/viper"
)

func Init() {

    output := zerolog.ConsoleWriter{Out: os.Stderr}
    log.Logger = log.Output(output)

    logLevel := viper.GetString("log.level");
    if logLevel == "" {

        logLevel = "info"
        log.Warn().Msg("Couldnt get log level value from config file\n");
    }
    level, err  := zerolog.ParseLevel(logLevel)
    if err != nil {

        level = zerolog.InfoLevel
        log.Warn().Err(err).Msg("Couldnt get log level value from config file\n");
    }
    zerolog.SetGlobalLevel(level)

    log.Info().Str("logLevel", logLevel).Msg("Logger Initialzied.\n");
}
