# Engage Digital API Specifications

This folder contains API specifications for Engage Digital.

The reference document is the file `engage-digital_openapi3.yaml`.

Other specifications are auto-generated from this file.

## Postman Collection

Use `swaggman` to create the Postman 2.x collection from the OpenAPI 3 API Specification.

### Installation

The following will install the `swaggman` executable in the `~/go/bin` directory.

```bash
$ go get github.com/grokify/swaggman
```

### Usage

```bash
$ swaggman -c engage-digital_postman2.config.json -b engage-digital_postman2.base.json -o engage-digital_openapi3.yaml -p engage-digital_postman2.json
```

In Postman, set the following environment variables:

| Environment Variable | Example |
|----------------------|--------    -|
| `RINGCENTRAL_ENGAGE_SERVER_URL` | `https://{myaccount}.api.engagement.dimelo.com` |
| `RINGCENTRAL_ENGAGE_ACCESS_TOKEN` | `deadbeef0123456789abcdef` |

#### Testing

Try the "Get all Users" API via "Provisioning" > "Users" > "Getting all Users".