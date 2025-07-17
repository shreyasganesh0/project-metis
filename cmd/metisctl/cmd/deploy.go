package cmd

import (
    "fmt"
    "os"
    "gopkg.in/yaml.v3"
    "github.com/spf13/cobra"
    "github.com/shreyasganesh0/project-metis/pkg/metis"
    "github.com/shreyasganesh0/project-metis/internal/kubernetes"
)

var deployCmd = &cobra.Command {

    Use: "deploy",
    Short: "Deploy the service defined by the metis.yaml",
    Long: `Looks for the manifest yaml "metis.yaml", parses
           it and deploys the service as specified`,
    RunE: func(cmd *cobra.Command, args []string) error {

        _, err := kubernetes.NewClient()
        if err != nil {

            return fmt.Errorf("Failed to setup k8s cluster: %w\n", err);
        }
        fmt.Println("Successfully connected to kubernetes cluster.")

        manifest_bytes, err := os.ReadFile("metis.yaml") //should be fine, the file isnt huge
        if err != nil {

            return fmt.Errorf("Couldnt read manifest file:%w",err);
        }

        var metis_service metis.ServiceManifest
        fmt.Println("--> Unmarshalling Metis Service Manifest..\n");

        if err := yaml.Unmarshal(manifest_bytes, &metis_service); err != nil {

            return fmt.Errorf("Error unmarshalling file: %w", err);
        }
        fmt.Println("--> Metis Service Manifest unmarshalled");
        fmt.Printf("Metis Service name: %s\n\n", metis_service.Name);

        fmt.Println("--> Generating K8s Deployment\n");

        deployment := kubernetes.GenerateDeployment(&metis_service)
        dep_byts, err_dep := yaml.Marshal(deployment)
        if err_dep != nil {

            return fmt.Errorf("Error converting deployment to YAML: %w\n", err_dep);
        }

        fmt.Println("Deployment Generated\n");
        fmt.Println("--> Generating K8s Service\n");

        service := kubernetes.GenerateService(&metis_service)
        serv_byts, err_serv := yaml.Marshal(service)
        if err_serv != nil {

            return fmt.Errorf("Error converting service to YAML: %w\n", err_serv);
        }
        fmt.Println("Service Generated\n");

        fmt.Println("---")
        fmt.Println(string(dep_byts));
        fmt.Println("...")
        fmt.Println("---")
        fmt.Println(string(serv_byts));
        fmt.Println("...")

        return nil;
    },
}

func init() {

    rootCmd.AddCommand(deployCmd);
}
