package metis

import (
)

type ServiceManifest struct {

    ApiVersion string `yaml:"apiVersion"`
    Kind string `yaml:"king"`
    Name string `yaml:"name"`
    Language string `yaml:"go"`
    Port int `yaml:"port"`
}
