package kubernetes

import (
    "github.com/shreyasganesh0/project-metis/pkg/metis"
    "k8s.io/apimachinery/pkg/util/intstr"
    appsv1 "k8s.io/api/apps/v1"
    corev1 "k8s.io/api/core/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


func GenerateDeployment(service *metis.ServiceManifest, image_tag *string) (*appsv1.Deployment) {

    labels := map[string]string{

        "app": service.Name,
    }

    replica_count := int32(1);

    return &appsv1.Deployment{

        //metav1.TypeMeta{},//optional

        ObjectMeta: metav1.ObjectMeta{

            Name: service.Name,
            Namespace: "default",
            Labels: labels,
        }, //optional maybe can ignore

        Spec: appsv1.DeploymentSpec{

            Replicas: &replica_count, //optional default 1

            Selector: &metav1.LabelSelector{

                MatchLabels: labels,
                //MatchExpressions: []LabelSelectorRequirements
            },

            Template: corev1.PodTemplateSpec{

                ObjectMeta: metav1.ObjectMeta{

                    Name:service.Name,
                    Labels: labels,
                },

                Spec: corev1.PodSpec{

                    Containers: []corev1.Container{
                        {

                            Name: service.Name,
                            Image: resolveTag(service, image_tag),
                            Ports: []corev1.ContainerPort{
                                {

                                    //HostPort: int32(service.Port),
                                    ContainerPort: int32(service.Port),
                                },
                            },
                        },
                    },
                },

            }, //template.spec.restartPolicy value must be "Always"

            //Strategy: appsv1.DeploymentStrategy{}, //optional pods replacement strategy

            //MinReadySeconds: int32(), //optional(default 0 if not specified)

            //RevisiontHistoryLimit: &int32(10), //optional(default 10 if not specified)

            //Paused: false, //optional indicates if deployment is false

            //ProgressDeadlineSeconds: &int32(600), //optional (default 600) time before deployment                                                  //considered failed if no progress made
        },

        //Status: appsv1.DeploymentStatus{},//optional
    }
}


func GenerateService(service *metis.ServiceManifest) (*corev1.Service) {

    labels := map[string]string {

        "app": service.Name,
    }

    return &corev1.Service{

        ObjectMeta: metav1.ObjectMeta{

            Name: service.Name,
            Labels: labels,
            Namespace: "default",
            Annotations: map[string]string {

                "prometheus.io/scrape": "true",
            },
        },

        Spec: corev1.ServiceSpec{

            Selector: labels,

            Ports: []corev1.ServicePort{
                {

                    Protocol: corev1.ProtocolTCP,
                    Port: 80,//internal to container
                    TargetPort: intstr.FromInt(service.Port),

                },
            },

            Type: corev1.ServiceTypeClusterIP, //default internal service type
        },
    }
}

func resolveTag(service *metis.ServiceManifest, image_tag *string) string {

    if *image_tag == "" {

        return service.Name + ":latest"
    }
    return *image_tag
}
