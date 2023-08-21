FROM golang:1.17-alpine AS build

WORKDIR /src/

COPY main.go go.* /src/

RUN CGO_ENABLED=0 go build -o /bin/demo


# Multistage build to get a small size image
FROM scratch

COPY --from=build /bin/demo /bin/demo

ENTRYPOINT ["/bin/demo"]