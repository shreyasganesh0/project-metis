package cmd

import (
    "fmt"
    "os"
    "gopkg.in/yaml.v3"
    "github.com/spf13/cobra"
    "github.com/shreyasganesh0/project-metis/pkg/metis"
)

var deployCmd = &cobra.Command {

    Use: "deploy",
    Short: "Deploy the service defined by the metis.yaml",
    Long: `Looks for the manifest yaml "metis.yaml", parses
           it and deploys the service as specified`,
    RunE: func(cmd *cobra.Command, args []string) error {

        manifest_bytes, err := os.ReadFile("metis.yaml") //should be fine, the file isnt huge
        if err != nil {

            return fmt.Errorf("Couldnt read manifest file:%w",err);
        }

        var service metis.ServiceManifest

        if err := yaml.Unmarshal(manifest_bytes, &service); err != nil {

            return fmt.Errorf("Error unmarshalling file: %w", err);
        }

        fmt.Printf("--> Parsed service name: %s\n", service.Name)
        fmt.Printf("Port: %d, Language used: %s\n", service.Port, service.Language)

        return nil;
    },
}

func init() {

    rootCmd.AddCommand(deployCmd);
}
