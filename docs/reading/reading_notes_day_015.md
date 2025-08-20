# K8s Annotations And Prometheus

## K8s Annotations
- attach arbitrary non-identifying metadata to objects
- used by tools (our case prometheus) to retrieve the metadata
    - use either labels or annotations to attach metadata
        - labels - used to select objects and find collections of objects
        - annotations - large/small structured/unstructured and include chars not allowed
          by labels
    - "metadata" {
        "annotations {
            "key1": "value1",
            "key2": "value2"
         }
      }
    - keys and values must be strings
    - ex:
        - fields managed by a declarative configuration layer
            - used to differentiate those values from default client and server set values
            - also auto generated fields from auto scaling and sizing systems
        - build, release or image information like timestamps, release IDs, git branch
          PR numbers, Image hashes, registry address
        - pointers to logging monitoring analytics or audit repos
        - tool info for debugging (versions, name, build info)
        - user/tool provenance info like urls of objects from ecosystem components
        - rollout tool metadata: config or checkpoints
        - directives from the end user to implementations to modify behaviour or engage
          non standard features
    - could be stored in a db but then it would be hard for tools and libraries to use

- syntax and char set
    - key-value pairs
        - key
            - [prefix]/name
            - 63 chars or less
            - [a-z0-9A-Z]
            - dashses(-), underscores(_), dots(.)
            - prefix if specified must be a DNS subdomain
                - series of DNS labels seperated by .
                - not longer than 253
            - if prefix omitted the key is private to user
            - automated system components (kube-scheduler, kubectl etc.)
              must add prefix
            - k8s.io/ and kubernetes.io/ prefixes reserved for kuberenetes core components
