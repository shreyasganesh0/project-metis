# Viper and YAML

## Viper
- Configuration management done via Viper
    - setting defaults
    - live watching and re reading from config file
    - Uses configuration files either TOML, YAML, JSON, HCL, INI, envfile
      and Java Properties files
    - readin from remote config system etcd and Consul
    - reading from buffer
    - setting explicit values
    - reading from command line flags
- Registry for configuration basically

- Reading Config Files
    - single viper instance supports a single config file
    - does not default to any paths so need to set it explicitly in app
    - SetConfigName("config") // name without extension
    - SetConfigType("yaml")
    - AddConfigPath("/etc/appname/") // can set this multiple times
    - ReadInConfig()
        - produces ConfigFileNotFoundError if file not found

- Working with Enviornment variables
    - enable out of the box 12 factor applications
    - ENV varialbles are case sensitive

    - SetEnvPrefix(string)
        - tell viper to prefix while reading from the env variable
        - used by BindEnv(string...) error and AutomaticEnv

    - BindEnv(string...) :error
        - first param is the keyname
        - rest are name of env variables to bind to this key
        - will take precendence in specified order
        - if ENV variable name not provided
            - viper will match prefix+"_"+key name in all caps
        - if ENV variable name specified (second param)
            - doesnt automatically add the prefix
        - value is read each time it is accessed

    - AutomaticEnv
        - viper will check for a variable name on every viper.Get
        - check for env variable with anem matching the key and prefix with EnvPrefix

    - SetEnvKeyReplacer
        - strings.Replacer object to rewrite Env keys
    - EnvKeyReplacer along with NewWithOptions factory function will allow you
      to write custome replacer logic
    - default env variables are unset to empty env variables as set use AllowEmptyEnv

    SetEnvPrefix("spf") // capitalized automatically
    BindEnv("id")

    os.SetEnv("SPF_ID", "13") //done outside the app usually
    id := Get("id") //13

## YAML
- YAML Aint Markup Language
- starts with ---
    - each set of --- in a yaml file is recognized as the start of a new document
- key: value pairs
    - strings; doe: "is a deer"
        - strings can be single double or no quotes
    - floating; points pi: 3.14159
    - boolean; xmax: true
    - integer; french-hens: 3
    - array; calling-birds:
                - huey
                - dewey
                - louie
                - fred
    - nesting; xmas-fifth-day:
                    calling-birds: 4
                    french-hens: 3
                    partridges:
                        count: 1
                        location: "pear tree"
- newlines are end of field
- no tabs for indentations
- anything after # comments in a yaml
- recognizes numeric types
    - bar: 0x12d4 //hex
    - foo: 0232 //octal
- floating point exponents
    - bar: 12.3015e+05
- NAN and INF
    - foo: .inf
    - foo: -.Inf
    - foo: .NAN
- string to greater than 1 line using >
    - bar: >
        this is not a normal
        string it goes across two
        lines
    - can also use the |
        bar: |
            multiline string
            used in yaml
- null fields represent with ~ or null
- booleans represent with
    - foo: True
    - foo: False
    - foo: On
    - foo: Off
    - doesnt need to be capitalized to work
- lists and arrays
    - items: [1, 2] //single line
    - list:
        - 1
        - 2 //multiple lines
- dictionaries
    - foo: {thing1: bar, thing2: huey} //single line
    - exa:
        bar:
            - foo
            - item2 //mulitple line dict with a array in it
- chomp modifiers
    - for multiline values if we want to strip or preseve tailing whitespace
      save the last character
    - bar: >+
        multiline stirn
        with
        taliing whitespace saved

    - bar: |-
        the tailing
        whitespace
        will be removed
        in this

- documents start with --- and ends with ...
    - .. is usually optional


