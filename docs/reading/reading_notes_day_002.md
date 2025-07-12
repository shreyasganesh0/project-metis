# Go Linker Flags and Cobra Commands

## [Go Linker Flags](https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications)

### Problem Statement
Need a way to add version, VCS commit id, build machine build time and other metadata.
Since this data is highly dynamic, having it in the source code and modifying is not
feasible and prone to breaking.

### Solution
- -ldflags with go build inserts this dynamic information into the binary
- ld uses the go toolchain linker cmd/link
- linker allows you to change packages at build time using command line args

## How To
go build -ldflags="-X 'package_path.varialble_name=new_value'"
used to write values to a vairable at link time

-X is one of many possible link flags that can be passed
the key-value pair represents the variable we want to change and the new value
    'package_path.varialble_name=new_value'

- Rules
    - the variable must be a package level variable of type string
      (maybe exported or unexported variable)
    - value cannot be const
    - cannot be set by the result of a function call

- Modifying subpackages
    - if we know the path we can just do 'relative/path/to/package.variable_name'
    - alternate way is to use the nm tool
        - go tool nm
        - outputs symbols involved with a given executable, obj file or lib
        - we can use grep after generating the symbol table from  nm to search
          for the variable
        - quickly get the path info.

## Cobra Init
When we call the cmd.Execute() function from the cobra.Command
it will go and call the init() function in all cobraCommands
within the funciton
it is a function of the Go init() function
