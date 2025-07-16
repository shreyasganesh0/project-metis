package kubernetes

import (
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
)

func NewClient() (*kubernetes.Clientset, error) {

    loadingRules := clientcmd.NewDefaultClientConfigLoadingRules();

    kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules,
                    &clientcmd.ConfigOverrides{});

    config, err := kubeConfig.ClientConfig();
    if err != nil {

        return nil, err
    }

    return kubernetes.NewForConfig(config);
}
