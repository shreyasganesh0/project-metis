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
)

var (

clientset *clientk8s.Clientset

deployCmd = &cobra.Command {

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

        if err := SetClientset(); err != nil {

            return fmt.Errorf("Error trying to create client set from config: %w\n", err)
        }

        fmt.Println("--> Generating K8s Deployment\n");

        deployment := kubernetes.GenerateDeployment(&metis_service)
        // _, err_dep := yaml.Marshal(deployment)
        // if err_dep != nil {
        //
        //     return fmt.Errorf("Error converting deployment to YAML: %w\n", err_dep);
        // }

        if err := CreateDeploymentResources(deployment); err != nil {

            return fmt.Errorf("Error while trying to create deployment in k8s: %w\n", err)
        }

        fmt.Println("Deployment Generated\n");
        fmt.Println("--> Generating K8s Service\n");

        service := kubernetes.GenerateService(&metis_service)
        if err := CreateServiceResources(service); err != nil {

            return fmt.Errorf("Error while trying to create service in k8s: %w\n", err)
        }
        // _, err_serv := yaml.Marshal(service)
        // if err_serv != nil {
        //
        //     return fmt.Errorf("Error converting service to YAML: %w\n", err_serv);
        // }
        fmt.Println("Service Generated\n");

        // fmt.Println("---")
        // fmt.Println(string(dep_byts));
        // fmt.Println("...")
        // fmt.Println("---")
        // fmt.Println(string(serv_byts));
        // fmt.Println("...")

        return nil;
    },
}
)

func init() {

    rootCmd.AddCommand(deployCmd);
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

    fmt.Println("Creating K8s service...")

    result, err := serviceClient.Create(context.TODO(), service, metav1.CreateOptions{})
    if err != nil {

        return err
    }
    fmt.Printf("Created Service %q.\n", result.GetObjectMeta().GetName())

    return nil;
}

func CreateDeploymentResources(deployment *appsv1.Deployment) error{

    deploymentsClient := clientset.AppsV1().Deployments(corev1.NamespaceDefault)

    fmt.Println("Creating K8s deployment...")

    result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
    if err != nil {
        return err
    }
    fmt.Printf("Created Deployment %q,\n", result.GetObjectMeta().GetName())

    return nil;

}
