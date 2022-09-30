# DSV Go SDK

The purpose of this application is to provide a simple service for storing and getting secrets.

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Generate SDK

1. Install [swagger-codegen](https://github.com/swagger-api/swagger-codegen).
2. Move `swagger.json` to the current folder.
3. Generate models: `swagger-codegen generate -i ./path/to/swagger.json -l go -o ./models --model-name-prefix DSV -Dmodels`.
4. Generate API client: `swagger-codegen generate -i ./path/to/swagger.json -l go -o ./client --model-name-prefix DSV`.

### Generate with docker

```
docker run --rm -v ${PWD}:/local swaggerapi/swagger-codegen-cli generate \
    -i /local/path/to/swagger.json \
    -l go \
    -o /local/client \
    --model-name-prefix DSV
```

