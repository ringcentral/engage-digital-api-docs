# Authenticating to the Engage API

Every request must provide an access token to authenticate properly.

!!! warning "Keep access token secure"
    Do not publish your access token publicly. The access token is **unencrypted**, and possession of it by a third-party will give them access to your account. 

## Obtaining an Access Token

Contact your project manager to have an access token provisioned to you. Each access token is associated to an existing agent -- all contents, interventions you make will be published as associated agent.

Note that some API methods requires authorization. It depends from the token’s user permissions. Authorization is described on all API methods below.

## Transmitting an Access Token

### Via Form Parameter

An access token can be specified in a request parameter named `access_token`.

#### Example

To get all interventions on the source accessible by the token’s users, URL will looks like:

`https://[YOUR DOMAIN].api.engagement.dimelo.com/1.0/interventions?access_token=abc42`

### Via HTTP Header

In order to be compliant with OAuth 2.0 standards an access token can also be specified via the `Authorization` request header where value respects following format:

`Authorization: Bearer <access_token_value>`

#### Example

To get all interventions on the source accessible by the token’s users, you’ll need to build your request with Authorization request header with proper value. HTTP request will looks like:

```http
GET /1.0/interventions HTTP/1.1
Host: test.api.engagement.dimelo.com
Authorization: Bearer abc42
```


