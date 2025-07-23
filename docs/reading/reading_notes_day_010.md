# Kubectl Rollout and Port Forwarding

## Rollout
Manage rollout of one or more resources
    - useful for depoyments, daemonsets and statefulsets
- kubectl rollout SUBCOMMAND
    - history
        - kubectl rollout history (TYPE NAME | TYPE) [flags]
    - kubectl rollout resource_command RESOURCE
        - pause
            - new updates to the resource wont have affect
            - current state will continue
        - restart
            - restart rollout
        - resume
            - resume rollout
        - status
            - --watch flag for continously checking till finished
        - undo
- --selector string
    - this flag can be used to filter on label
    - supports: = == != in notin
    - matching objets must satisfy all of the label constraints

## Port-Forward
```
kubectl port-forward TYPE/NAME [options] [LOCAL_PORT:]REMOTE_PORT [...[LOCAL_PORT_N:]REMOTE_PORT_N]
```

- port forward one or more local ports to a pod
- using resource type/name like deployment/my_deployment
    - default resource is pod
- will select pod automatically if multiple match
- forwarding ends on pod termination and will have to rerun the command
