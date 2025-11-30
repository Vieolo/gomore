# gomore & go.yaml
CLI support for the go.yaml, centralized config file for the Go projects

- `go.yaml` is an optional centralized config file for go project, with no intention to replace `go.mod`
- `gomore` is a cli to provide common functionalities for `go.yaml`.
- `go.yaml` is independent of `gomore`
- `go.yaml` and `gomore` cli are not a new build system, and will use existing ecosystem under the hood

## go.yaml
A very useful convention of many languages is a centralized and extendible config file for a project that not only holds the project's meta data, but also the configuration of the dependencies and tools. 

Go has a go.mod that holds the module name, dependencies, and the version of Go used in the project. However, in many projects, there is a tangible need for storing extra meta data and configs.

`go.yaml` acts as the centralized config file for a Go project which contains some standard fields and can be used to contain any other configuration for any other reason, tool, or workflow. `go.yaml` is an standalone config file and does not need `gomore` cli to be usable. Any tool can use `go.yaml` for their configuration if desired.

So...
- JS/TS -> package.json
- Python -> pyproject.toml
- Rust -> Cargo.toml
- Dart -> pubspec.yaml
- ***Go -> go.yaml***


### Structure of go.yaml
`go.yaml` has a series of keys that are placed on the root of the file, and some root-level reserved objects. Beyond that, you add any arbitrary objects to the configuration. 

Here is the overview of the structure:

```yaml
## Root level reserved keys
##
name: myproject # The human-readable name of the project
description: my new project # [Optional] description of the project
version: 12.3.5 # version of the project
homepage: https://example.com # [Optional] homepage of the project
repository: https://github.com/example # [Optional] URL pointing to the project's source code
issue_tracker: https://github.com/example/issues # [Optional] URL pointing to an issue tracker for the project
documentation: https://example.com/docs # [Optional] URL pointing to documentation for the project
license: MIT # [Optional] the license of the project

## Root level reserved objects
##

# [Optional] pre-defined commands
# The commands object is optional, but if provided,
# should be a map of key-value pair (string: string). The commands can be anything you wish.
commands:
    build: go build .
    test: go test
    some-other-command: other command
    ...


# [Optional] output binaries
# The binaries object is optional, but if provided,
# should be a sequence of objects, each object keys being the os, and the value being the path to the binary
binaries:
    server:
        macos: out/server
        linux: out/server
    admin:
        macos: out/admin
        linux: out/admin
    ...
    
# [Optional] Reserved for future use
profiles:
    ...

# [Optional] Reserved for future use
workflows:
    ...

# [Optional] Reserved for future use
tools:
    ...

# [Optional] Reserved for future use
generate:
    ...
    
## Non-reserved and arbitrary config
##

# Any tool or project can add and use their configuration and settings
my-tool:
    some-flag: true
```

### Stability status
At the moment, `go.yaml` is not yet guaranteed to be stable and backward compatible. After we are confident that we have a solid structure by community feedback, `go.yaml` will reach a stable state and will avoid any breaking changes.


### Roadmap for go.yaml
- [ ] Introduce profiles
- [ ] Introduce workflows
- [ ] Introduce tools
- [ ] Introduce generate
- [ ] Locking for stability and avoid breaking changes


## gomore CLI
`gomore` is a cli tool to provide a series of common functionalities based on the standard configurations of the `go.yaml`.

### Install
```bash
go install github.com/vieolo/gomore@latest
```

Other ways to install the cli will come soon.

### usage
- `init` -> initializes a new `go.yaml` file
- `run` -> Runs one of the defined commands


### Roadmap for gomore
- [ ] profiles (prod/dev/test)
- [ ] homebrew install tap
- [ ] workflow runner
- [ ] environment schema validation
- [ ] validate command (`e.g., gomore doctor`)
