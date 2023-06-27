#!/bin/sh
docker build . -t api-doc-preview
docker run --rm -it -p ${1:-8000}:8000 api-doc-preview
