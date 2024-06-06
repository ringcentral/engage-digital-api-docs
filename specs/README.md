# Engage Digital API Specifications

This folder contains API specifications for Engage Digital.

Use the following files:

* `engage-digital_openapi3.yaml`: this is the reference spec.
* `engage-digital_postman2.json`: this spec is auto-generated from the OpenAPI 3 specification, along with `engage-digital_postman2.config.json` and `engage-digital_postman2.base.json` using [`spectrum`](https://github.com/grokify/spectrum).

The following files are used to generate the Postman collection and are not designed to be used on their own:

* `engage-digital_postman2.config.json`: configuration file for Spectrum Postman Collection generator.
* `engage-digital_postman2.base.json`:

## Postman Collection

### Script

Run `./gen_postman.sh`

or proceed with the manual installation as described below

Use `spectrum` to create the Postman 2.x collection from the OpenAPI 3 API Specification.

### Installation

The following will install the `spectrum` executable in the `~/go/bin` directory.

```bash
$ go install github.com/grokify/spectrum@v1.10.3
```

This approach requires `go` 1.16 be installed on your system. See more here: [https://golang.org/](https://golang.org/).

### Usage

The following example writes the output to `engage-digital_postman2.json`.

```bash
$ ~/go/bin/spectrum --config engage-digital_postman2.config.json --basePostmanFile engage-digital_postman2.base.json --openapiFile engage-digital_openapi3.yaml --postmanFile engage-digital_postman2.json
```

In Postman, set the following environment variables:

| Environment Variable | Example |
|----------------------|---------|
| `ENGAGE_DIGITAL_SERVER_URL` | `https://{myaccount}.api.engagement.dimelo.com` |
| `ENGAGE_DIGITAL_ACCESS_TOKEN` | `deadbeef0123456789abcdef` |

#### Testing

Try the "Get all Users" API via "Provisioning" > "Users" > "Getting all Users".
