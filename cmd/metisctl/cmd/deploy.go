package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var deployCmd = &cobra.Command {

    Use: "deploy",
    Short: "Deploy the service defined by the metis.yaml",
    Long: `Looks for the manifest yaml "metis.yaml", parses
           it and deploys the service as specified`,
    RunE: func(cmd *cobra.Command, args []string) error {

        fmt.Println("--> Running the deploy command.");
        return nil;
    },
}

func init() {

    rootCmd.AddCommand(deployCmd);
}
