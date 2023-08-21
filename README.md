# gofolderingproject

GO Foldering Project is a project structure that I often use for my development using clean architecture (combination a few architecture, pattern that we think make our development better) approach.

## Getting started
### Before start, please run this
```
go mod install
go mod tidy
go mod vendor
```

### How To Try to Run
```make start-public```

## Structure of Project
- cmd: Application dependency injection and entrypoint
    - rest
    - grpc
- config: Read configuration from any sources (such as .env, yaml, etc) and maintain our config with Viper
- init: manually sql schema execution
- internal: internal logic
    - delivery: handler / port
    - helper: this folder is used for helper function
    - middleware: any middleware functions are here
    - model: data structure or struct
    - repository: is a folder using a repository pattern to connect with database (NOSQL or SQL) or another service
    - usecase: business logic
- pkg: folder to connect with other package or wrap other library that we want to use
- Dockerfile
- .gitlab-ci.yml
- main.go
- Makefile
- README.md

Notes: SSL configuration (key & certificate) must be saved on the save place. We use it to make secure connection both HTTP and gRPC web server

## Tech Stack
- Golang - Programming Language
- Gin (HTTP REST Framework - https://github.com/gin-gonic/gin)
- sqlx (https://github.com/jmoiron/sqlx) - library to connect with database/sql
- zerolog (https://github.com/rs/zerolog) - library to make a great logging standart
- OpenTracing (Still used open tracing, it is need to be updated to Open Telemetry https://github.com/open-telemetry/opentelemetry-go) - as tracer library
- Viper (https://github.com/spf13/viper) - part of sf13, library to manage configuration
- Cobra (https://github.com/spf13/cobra) - part of sf13, used to modern CLI applications

## Integrate with your tools

- [ ] [Set up project integrations](https://github.com/aszanky/gofolderingproject/-/settings/integrations)

## How to Contribute
Please flow Github Flow <br />
https://docs.github.com/en/get-started/quickstart/github-flow <br />

![image githubflow](/assets/github-flow-1.png)