package metis

import (
)

type ServiceManifest struct {

    ApiVersion string `yaml:"apiVersion"`
    Kind string `yaml:"kind"`
    Name string `yaml:"name"`
    Language string `yaml:"language"`
    Port int `yaml:"port"`
}
