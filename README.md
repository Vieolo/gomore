# gomore
CLI tool to provide functionalities that complement the official Go tooling

`gomore` relies on `go.yaml` which you can read more about it [here](https://github.com/vieolo/godotyaml) which explains the schema and provides a light weight parser.

## Features
- Running pre-defined commands

## Philosophy
- `gomore` is not a replacement for `go` cli or any other part of the official go toolchain
- `gomore` has no overlap with the official tooling and such overlap is not planned for the future.


## Install
```bash
go install github.com/vieolo/gomore@latest
```

Other ways to install the cli will come soon.

## Commands
- `init` -> initializes a new `go.yaml` file
- `run` -> Runs one of the defined commands


## Roadmap
- [ ] homebrew install tap
- [ ] profiles (prod/dev/test)
- [ ] workflow runner
- [ ] environment schema validation
- [ ] validate command (`e.g., gomore doctor`)
