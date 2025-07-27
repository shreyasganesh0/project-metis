package cmd

import (
    "fmt"
    "os"
    "time"
    "context"
    "syscall"
    "os/signal"
    "path/filepath"

    "github.com/spf13/cobra"
    "github.com/spf13/viper"

    "github.com/rs/zerolog/log"

	"github.com/shreyasganesh0/project-metis/internal/logging"
)

var rootCmd = &cobra.Command {
    Use: "metisctl",
    Short: "A CLI for the Metis Internal Development Platform.",
    Long: `The Metis CLI is intended to be an Internal Development Platform tool.
            The tool will be useful to help build, deploy and operate microservices
            in a reliable and efficient manner. It is an abstraction over Kubernetes
            that provides a developer-friendly approach to deploying pods.`,
    RunE: func(cmd *cobra.Command, args []string) error{
		log.Info().Msg("Project Metis CLI (metisctl) is running. Press Ctrl+C to exit...\n");

        ctx, cancel := context.WithCancel(context.Background());

        sig_ch := make(chan os.Signal, 1);
        signal.Notify(sig_ch, os.Interrupt, syscall.SIGTERM);

        go func() {

            <-sig_ch
            cancel();
        }() // wait for ctrl+c signal

        select {
        case <-ctx.Done():
            log.Debug().Msg("Recieved SIGINT. Cleaning up..");
            time.Sleep(2 * time.Second);
        }
        log.Info().Msg("Cleanup complete. Shutting down.\n");

        return nil;
    },

    PersistentPreRun: func(cmd *cobra.Command, args []string) {

		logging.Init();

        home_dir, err := os.UserHomeDir();
        if err != nil {

			log.Fatal().Err(err).Msg("Error while getting home directory");
			return;
        }

        viper.AddConfigPath(filepath.Join(home_dir, ".metis"))
        viper.SetConfigName("config")
        viper.SetConfigType("yaml")
        viper.SetEnvPrefix("METIS")
        //viper.AutomaticEnv() doesnt work reliably for nested keys like user.name need to explicitlly bind
        if err_bind := viper.BindEnv("user.name", "METIS_USER_NAME"); err_bind != nil {
			log.Error().Err(err_bind).Msg("Failed to bind key\n");
			return;
        }

        if err_read := viper.ReadInConfig(); err_read != nil {

			log.Error().Err(err_read).Msg("Error while reading config file\n");
            panic(fmt.Errorf("Error while reading config file: %w", err_read));
        }
        log.Info().Msgf("Using config file: %s", viper.ConfigFileUsed());
    },
}

func Execute() {

    if err := rootCmd.Execute(); err != nil {

        os.Exit(1);
    }

}
