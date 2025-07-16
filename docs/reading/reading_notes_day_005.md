# Kuberenetes Client Library

## Kind (K8s in Docker)
- ```kind create cluster``` to create a k8s cluster
    - prebuild k8s clusters node image
    - specify image using --image=...
    - default name is "kind" --name to assing a different name
    - --wait flag blocks till control pane is in ready state
        - --wait 5m
    - choses either docker podman or nerctl automatically
        - KIND_EXPERIMENTAL_PROVIDER-docker (or others) to stop auto detect
    - --kubconfig to set the config file to load
        - config stored in ~/.kube/config if $KUBECONFIG
- ```kind get clusters``` to list clusters created
- ```kubectl cluster-info --context kind-kind```
    - to interact with specified cluster
- ```kind delete cluster```
    - similar to create cluster will use the "kind"

- ```kind load docker-image my-custom-image-0 my-custom-image-1```
    - load docker images into cluster nodes
    - --name to use specific cluster name
- Workflows usually look like
    - docker build -t custom-image:tag ./myimagedir
    - kind load docker-image custome-image:tag
    - kubectl apply -f manifest-using-image.yaml
- ```docker exec -it my-node-name critcl imgaes```
    - get a list of images present on a cluster node
    - :latest tag shouldnt be used with images
        - sets default pull policy to Always which can mess up the
          configuration and pull duplicates
        - have to specify imagePullPolicy: IfNotPresent explicitlly
          if using images with latest
- kind runs a local k8s cluster using docker containers as nodes
    - uses node-image to run k8s aritifcats like kubeadm or kubelet
    - node-image is built of base-image
    - node-image uses $GOPATH/src/k8s.io/kubernetes using source
    - ```kind build node-image``` uses this source
        - add v1.30.0 or other version tags to the end to specify
          version
        - --type option to specify build type (url|file|release (version no.)|source)
- env variables HTTP_PROXY | HTTPS_PROXY | NO_PROXY
    - sets proxy which is passed to everything in kind nodes
- kind export logs

## Client-Go library
- rest.InClusterConfig()
    - creates the in-cluster config
- kubernetes.NewForConfig(config)
    - creates a client list for the given config
- clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions())
    - get all pods from the namespace


