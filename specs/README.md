# Engage Digital API Specifications

This folder contains API specifications for Engage Digital.

Use the following files:

* `engage-digital_openapi3.yaml`: this is the reference spec.
* `engage-digital_postman2.json`: this spec is auto-generated from the OpenAPI 3 specification, along with `engage-digital_postman2.config.json` and `engage-digital_postman2.base.json` using [`swaggman`](https://github.com/grokify/swaggman).

The following files are used to generate the Postman collection and are not designed to be used on their own:

* `engage-digital_postman2.config.json`: configuration file for Swaggman Postman Collection generator.
* `engage-digital_postman2.base.json`: 

## Postman Collection

Use `swaggman` to create the Postman 2.x collection from the OpenAPI 3 API Specification.

### Installation

The following will install the `swaggman` executable in the `~/go/bin` directory.

```bash
$ go get github.com/grokify/swaggman
```

This approach requires `go` be installed on your system. See more here: [https://golang.org/](https://golang.org/).

### Usage

The following example writes the output to `engage-digital_postman2.json`.

```bash
$ ~/go/bin/swaggman --config engage-digital_postman2.config.json --basePostmanFile engage-digital_postman2.base.json --openapiFile engage-digital_openapi3.yaml --postmanFile engage-digital_postman2.json
```

In Postman, set the following environment variables:

| Environment Variable | Example |
|----------------------|---------|
| `ENGAGE_DIGITAL_SERVER_URL` | `https://{myaccount}.api.engagement.dimelo.com` |
| `ENGAGE_DIGITAL_ACCESS_TOKEN` | `deadbeef0123456789abcdef` |

#### Testing

Try the "Get all Users" API via "Provisioning" > "Users" > "Getting all Users".
