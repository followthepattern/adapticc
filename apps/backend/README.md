# adapticc-backend

Adapticc Backend Repository

### Pre requirements

    go version go1.14.3 darwin/amd64 or later
    docker version 19.03.8

## Getting Started
Make sure you install all required dependencies and set up project config.
You can create project config file by renaming [config/config.yaml.tmpl](./config/config.yaml.tmpl) -> config

To build and run project:

```
make dev-run
```

Also, you can run this project using hot-reload daemon, to run it:

```
go get github.com/githubnemo/CompileDaemon

make hot-run
```

You can check your source code with Go linter. You can find how to install it locally [here](https://golangci-lint.run/usage/install/#local-installation).

To run linter checking, just execute:
```
make lint
```
