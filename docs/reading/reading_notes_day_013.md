# Helm

## Helm
Helm is a tool used to manage k8s packages

### Three Big Concepts
- Chart is a Helm package
    - contain all resource definitions needed to run an application
      or service inside the k8s cluster
    - kind of like the k8s equivalent of a homebrew formula or APK dpkg
- Repository where charts can be collected and shared
- Release instance of a chart running in a k8s cluster
    - one chart can be installed multiple times in a k8s cluster

### Finding Charts
helm search hub - used to search aritifact hub which has charts from many repositories
    - helm search hub lists the URL to the location on aritfathub.io but
      not actual helm repo
    - helm search hub --list-repo-url exposes the actual helm repo URL

helm search repo - used to search for charts in locally added (helm repo add)
- helm search uses fuzzy string matching

### Installing a package
install a new package

helm install <release_name> <chart_name>
- can use the --generate-name option to generate release name

helm status
- uses to read config info


### Customizing chart
helm show values <chart_name>
- show what values are configurable
- override any of the settings using a YAML
helm install -f values.yaml <char_name> --generate-name
    - can also use the --set flag to override on the command line
    - --set name=value
    - if both are used then the --set flag options take precedence
    - can also do other yaml based stuff using set
        - --set outer.inner=value
        - --set name={a,b,c}
        - --set servers[0].port=80

### Upgrading and Recovering on Failure
helm upgrade -f config.yaml <release> <chart>
- used to upgrade an existing release
helm get values <release>
- to check values
helm rollback <release> <version>
- to rollback to the previous version

### Working with Repo
helm repo list
- to see configured repos
helm repo add <name> <url>
- add repo

### Create charts
helm create deis-workflow
helm package deis-workflow
- package chart for distribution
-
