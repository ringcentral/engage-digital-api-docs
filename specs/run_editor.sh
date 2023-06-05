#!/bin/sh
docker pull swaggerapi/swagger-editor
docker run -p ${1:-8042}:8080 -v $(pwd):/tmp -e SWAGGER_FILE=/tmp/engage-digital_openapi3.yaml swaggerapi/swagger-editor
