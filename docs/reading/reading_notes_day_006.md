# CRUD a K8s Deployment

## Setup
- get the kubeconfig
    - config = clientcmd.BuildConfigFromFlags("", *kubconfig)
        - here the kubeconfig is build using flag.String("kubeconfig", "path/to/kube/config")
- build a clientset from kubeconfig
    - clientset = kuberenetes.NewFromConfig(config)

- deploymentsClient = clientset.AppV1().Deployments(apiv1.NamespaceDefault)

## Create
deploymentsClient.Create(ctx, deployment, metav1.CreateOptions{})//default metadata
- create the deployment using deployments client
- deployment is the &appv1.Deployment{}

## Update
deploymentsClient.Update(ctx, deployment, metav1.CreateOptions{})
- 2 ways to do it
    - modify the deployment and call Update(deployment)
        - overrwrites the changes made by other clients between Create and Update
        - this is because its the same as udpating the reference
    - get the result from deploymentClient.Get()
        - update the result with new values
        - deploymentClient.Update()
        - retry until it succeeds retry.RetryOnConflict(retry.DefaultRetry, func () error {})
        - all this done ina lambda function passed to retry.RetryOnConflict
            - has exponential backoff built into this util
        - this is better because its not an udpate to the reference can iwll only apply
          the changes if possible

## List
deploymentsClient.List(ctx, metav1.CreateOptions{});
- returns a list with .Items
- .Name, .Spec.Replicas properites of the .Item

## Delete
deploymentsClient.Delete(ctx, deployment.ObjectMeta.Name, metav1.DeleteOptions{})

## metav1.XXXMeta{}
- each command has some metadata passed to it
- different types are defined in the k8s.io/apimachinery/pkg/apis/meta/v1
- most are optional fields that get populated automatically so we just need to pass the
  default object



