FROM golang:1.18

RUN go install github.com/grokify/spectrum@v1.15.0 # same as github tests

RUN mkdir /app
WORKDIR /app

CMD spectrum --config engage-digital_postman2.config.json --basePostmanFile engage-digital_postman2.base.json --openapiFile engage-digital_openapi3.yaml --postmanFile engage-digital_postman2.json
