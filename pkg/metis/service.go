package metis

import (
    "gopkg.in/yaml.v3"
)

type MetisManifest struct {

    ApiVersion string `yaml:"apiVersion"`
    Kind string `yaml:"king"`
    Name string `yaml:"name"`
    Language string `yaml:"go"`
    Port int `yaml:"port"`
}
