package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var (
    version = "dev"
    commit = "none"
    date = "unkown"
)

var versionCmd = &cobra.Command {

    Use: "version",
    Short: "Command to get version number and build details of metisctl.",
    Long: `This command will display version details of the binary
            It provides details of version number, commit hash and build date.`,
    RunE: func(cmd *cobra.Command, args []string) error{

        fmt.Printf("Metisctl Version: %s\nGit Commit: %s\nBuild Date: %s\n",
                    version, commit, date);
        return nil;
    },
}

func init() {

    rootCmd.AddCommand(versionCmd);
}
