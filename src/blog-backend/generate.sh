#!/bin/bash

oapi-codegen -generate types -o ports/openapi_types.gen.go -package ports openapi.yaml

oapi-codegen -generate gin,strict-server -o ports/openapi_api.gen.go -package ports openapi.yaml