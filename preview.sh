#!/bin/sh
docker build . -t api-doc-preview
docker run --rm -it -p 8000:8000 api-doc-preview
