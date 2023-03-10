#!/bin/sh
docker build . -t api-doc-gen-postman
docker run --rm -it -v .:/app api-doc-gen-postman
