# Engage Digital API Specifications

This folder contains API specifications for Engage Digital.

Use the following files:

* `engage-digital_openapi3.yaml`: this is the reference spec.
* `engage-digital_postman2.json`: this spec is auto-generated from the OpenAPI 3 specification, along with `engage-digital_postman2.config.json` and `engage-digital_postman2.base.json` using [`spectrum`](https://github.com/grokify/spectrum).

The following files are used to generate the Postman collection and are not designed to be used on their own:

* `engage-digital_postman2.config.json`: configuration file for Spectrum Postman Collection generator.
* `engage-digital_postman2.base.json`:

## Postman Collection

### Generate Postman Collection file

There are two ways to generate postman collections described below.

#### 1. Script (recommended)

Run `./gen_postman.sh`

#### 2. Manual installation

Use `spectrum` to create the Postman 2.x collection from the OpenAPI 3 API Specification.

##### 2.1 Installation

Be sure to have the version of go specified in the Dockerfile. If you have the latest version of go installed, you can install an oldest version with :

```bash
$ go install golang.org/dl/go1.18@latest
```

See more here: [https://golang.org/](https://golang.org/).

The following will install the `spectrum` executable in the `~/go/bin` directory.

```bash
$ ~/go/bin/go1.16 install github.com/grokify/spectrum@v1.15.0
```

##### 2.2 Usage

The following example writes the output to `engage-digital_postman2.json`.

```bash
$ ~/go/bin/spectrum --config engage-digital_postman2.config.json --basePostmanFile engage-digital_postman2.base.json --openapiFile engage-digital_openapi3.yaml --postmanFile engage-digital_postman2.json
```

### Testing

In Postman, set the following environment variables:

| Environment Variable | Example |
|----------------------|---------|
| `ENGAGE_DIGITAL_SERVER_URL` | `https://{myaccount}.api.engagement.dimelo.com` |
| `ENGAGE_DIGITAL_ACCESS_TOKEN` | `deadbeef0123456789abcdef` |

Try the "Get all Users" API via "Provisioning" > "Users" > "Getting all Users".
