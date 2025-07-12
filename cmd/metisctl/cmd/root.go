package cmd

import (
    "fmt"
    "os"
    "time"
    "context"
    "syscall"
    "os/signal"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command {
    Use: "metisctl",
    Short: "A CLI for the Metis Internal Development Platform.",
    Long: `The Metis CLI is intended to be an Internal Development Platform tool.
            The tool will be useful to help build, deploy and operate microservices
            in a reliable and efficient manner. It is an abstraction over Kubernetes
            that provides a developer-friendly approach to deploying pods.`,
    RunE: func(cmd *cobra.Command, args []string) error{
        fmt.Println("Project Metis CLI (metisctl) is running. Press Ctrl+C to exit...");

        ctx, cancel := context.WithCancel(context.Background());

        sig_ch := make(chan os.Signal, 1);
        signal.Notify(sig_ch, os.Interrupt, syscall.SIGTERM);

        go func() {

            <-sig_ch
            cancel();
        }() // wait for ctrl+c signal

        select {
        case <-ctx.Done():
            fmt.Println("Recieved SIGINT. Cleaning up..");
            time.Sleep(2 * time.Second);
        }
        fmt.Println("Cleanup complete. Shutting down.");

        return nil;
    },
}

func Execute() {

    if err := rootCmd.Execute(); err != nil {

        os.Exit(1);
    }
}






