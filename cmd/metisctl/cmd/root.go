package cmd

import (
    "fmt"
    "os"
    "github.com/spf13/cobra"
)

var root_cmd = &cobra.Command {
    Use: "metisctl",
    Short: "A CLI for the Metis Internal Development Platform.",
    Long: `The Metis CLI is intended to be an Internal Development Platform tool.
            The tool will be useful to help build, deploy and operate microservices
            in a reliable and efficient manner. It is an abstraction over Kubernetes
            that provides a developer-friendly approach to deploying pods.`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Project Metis CLI (metisctl) starting...");
        fmt.Println("This is the root command for the Metis CLI.");
    },
}

func Execute() {

    if err := root_cmd.Execute(); err != nil {

        os.Exit(1);
    }
}






