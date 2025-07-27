package cmd

import (
    "fmt"
    "os"
    "flag"
    "context"
    "path/filepath"
    "bufio"

    "gopkg.in/yaml.v3"
    "github.com/spf13/cobra"

    "github.com/shreyasganesh0/project-metis/pkg/metis"
    "github.com/shreyasganesh0/project-metis/internal/kubernetes"

    appsv1 "k8s.io/api/apps/v1"
    corev1 "k8s.io/api/core/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    clientk8s "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/util/homedir"
    "k8s.io/client-go/tools/clientcmd"

    "github.com/rs/zerolog/log"
)

var (

image_tag *string

clientset *clientk8s.Clientset

deployCmd = &cobra.Command {

    Use: "deploy",
    Short: "Deploy the service defined by the metis.yaml",
    Long: `Looks for the manifest yaml "metis.yaml", parses
           it and deploys the service as specified`,
    RunE: func(cmd *cobra.Command, args []string) error {

        _, err := kubernetes.NewClient()
        if err != nil {

            log.Fatal().Err(err).Msg("Failed to setup k8s cluster Client");
            return fmt.Errorf("Failed to setup k8s cluster: %w\n", err);
        }
        log.Debug().Msg("Successfully connected to kubernetes cluster.\n")

        manifest_bytes, err := os.ReadFile("metis.yaml") //should be fine, the file isnt huge
        if err != nil {

            log.Fatal().Err(err).Msg("Failed to read the manifest file\n");
            return fmt.Errorf("Couldnt read manifest file:%w",err);
        }


        var metis_service metis.ServiceManifest
        log.Info().Msg("--> Unmarshalling Metis Service Manifest..\n");

        if err := yaml.Unmarshal(manifest_bytes, &metis_service); err != nil {

            log.Error().Err(err).Msg("Couldnt unmarshal manifest bytes\n");
            return fmt.Errorf("Error unmarshalling file: %w", err);
        }

        log.Info().Str("serviceName", metis_service.Name).Int("port", metis_service.Port).Msg("--> Metis Service Manifest unmarshalled\n")

        if err := SetClientset(); err != nil {

            log.Error().Err(err).Msg("Couldnt create k8s clients\n");
            return fmt.Errorf("Error trying to create client set from config: %w\n", err)
        }

        log.Debug().Msg("--> Generating K8s Deployment\n");

        deployment := kubernetes.GenerateDeployment(&metis_service, image_tag)
        // _, err_dep := yaml.Marshal(deployment)
        // if err_dep != nil {
        //
        //     return fmt.Errorf("Error converting deployment to YAML: %w\n", err_dep);
        // }

        if err := CreateDeploymentResources(deployment); err != nil {

            log.Error().Err(err).Msg("Couldnt create deployment\n");
            return fmt.Errorf("Error while trying to create deployment in k8s: %w\n", err)
        }

        log.Info().Msg("Deployment Generated\n");
        log.Info().Msg("--> Generating K8s Service\n");

        service := kubernetes.GenerateService(&metis_service)
        if err := CreateServiceResources(service); err != nil {

            log.Error().Err(err).Msg("Couldnt create service\n");
            return fmt.Errorf("Error while trying to create service in k8s: %w\n", err)
        }
        log.Info().Msg("Service Generated\n");


        return nil;
    },
}
)

func init() {

    rootCmd.AddCommand(deployCmd);
    image_tag = deployCmd.Flags().StringP("image", "i", "", "provide image tag to be deployed.")

}

func SetClientset() error {

    var kubeconfig *string
    if home := homedir.HomeDir(); home != "" {
        kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "")
    } else {

        fmt.Println("Enter absolute path to .kube/config\n")
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan()
        if err := scanner.Err(); err != nil {

            return err
        }
        abs_path := scanner.Text()
        kubeconfig = flag.String("kubeconfig", "", abs_path);
    }
    flag.Parse()

    config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
    if err != nil {
        return err
    }

    clientSet, err := clientk8s.NewForConfig(config)
    if err != nil {
        return err
    }

    clientset = clientSet
    return nil;
}

func CreateServiceResources(service *corev1.Service) error {

    serviceClient := clientset.CoreV1().Services("default");

    log.Info().Msg("Creating K8s service...")

    result, err := serviceClient.Create(context.TODO(), service, metav1.CreateOptions{})
    if err != nil {

        return err
    }
    log.Info().Str("serviceName", result.GetObjectMeta().GetName()).Msg("Created Service\n")

    return nil;
}

func CreateDeploymentResources(deployment *appsv1.Deployment) error{

    deploymentsClient := clientset.AppsV1().Deployments(corev1.NamespaceDefault)

    log.Info().Msg("Creating K8s deployment...")

    result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
    if err != nil {
        return err
    }
    log.Info().Str("deploymentName", result.GetObjectMeta().GetName()).Msg("Created Deployment\n")

    return nil;

}
