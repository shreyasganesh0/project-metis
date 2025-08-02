# Prometheus Scraping

## Kuberenetes SD Config
- scrape k8s using the k8s REST API
- syncs with cluster state
- "role" types
- "node"
    - discovers one target per cluster node with address defaulting to k8s http port
        - order is NodeInternalIP, NodeExternalIP, NodeLegacyHostIP, NodeHostName
    - instance label of the node will be set to node name from the API server

- "service"
    - used to discover a service port for each service
    - useful for blackbox monitoring of a service
    - address will be set to k8s DNS name of the service and service port
- "pod"
    - discovers all pods and exposes their containers as targets
    - single target per port per container
    - if no port present in container a port-free target is generated
- "endpoints"
    - listed endopints of a service
    - one target per port per endpoint
    - if enpoint backed by a pod all additional ports of the
      pod not bound to an endpoint port are discovered as well
- "endpointslice"
    - each endpoint address referenced in the endpointslice object
    - if endpoint backed by a pod all ports of the pod are discovered
- "ingress"
    - blackbox monitoring of an ingress
    - the address set to the host specified in the ingress spec
