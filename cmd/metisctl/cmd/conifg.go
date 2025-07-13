package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var configCmd = &cobra.Command {

    Use: "config",
    Short: "Manage Metisctl configuration.",
    Long: `Use the various subcommands to manage configuration values of metisctl`,
}

var configViewCmd = &cobra.Command {

    Use: "view <key>",
    Short: "Display a single configuraiton value by key",
    Args: cobra.ExactArgs(1), // since it takes key as args
    RunE: func(cmd *cobra.Command, args []string) error {

        key := args[0]
        val := viper.Get(key)
        fmt.Printf("%s: %v\n", key, val);
        return nil
    },
}

func init() {

    configCmd.AddCommand(configViewCmd)
    rootCmd.AddCommand(configCmd)
}
