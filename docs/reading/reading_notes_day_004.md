# YAML Parsing And Declarative Model

## YAML Parsing in Go

### UnMarshal
- Unmarshal(in []byte, out interface{}) (err error)
- in
    - the in is a byte slice
    - the first document inside the byte slice is unmarshalled
    - if an internal pointer
- out
    - accepts Maps and pointers to int struct and string
    - types of decoded values must be compatible with values in out
    - decoding continues and error *yaml.TypeError returned for missing values
    - struct fields must be Exported aka have their Name capitalized
        - unmarshalled using the field name lowercased
    - custome keys defined via the "yaml" name in the field tag
        - content before the first , is the key and
        - yaml:"[<key>][,<flag1>[,<flag2>]]"
            - flag types
            - omitempty - only set field if its not set to zero value
            - flow - flow style useful for structs maps
            - inline - inlienthe field must be a struct or a map
            type StructA struct {
                A string `yaml:"a"`

                }
            type StructB struct {

                StructA `yaml:",inline"`
                B string `yaml:"b"`
            }

## Declarative Model
spec:
    containers:
    - name: nginx
      image: nginx:1.14.2
      ports:
      - containerPort: 80
- this is the declarative model that kubernetes deployment yamls use
- declare the state you want  the system to be in and the system makes it so
