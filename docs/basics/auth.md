# Authenticating to the RingCX API

## Access Tokens

Every request must provide an access token to authenticate properly.

<a class="btn btn-primary" href="../access-token/">Obtain an Access Token</a>

!!! note "Access Token Permissions"
    Different API endpoints require different permissions. The permissions assocated with an access token are inherited from the associated user. Read about [creating an access token](../access-token/) to learn how to associate an access token with a user. 

## Transmitting an Access Token

### Via Form Parameter

An access token can be specified in a request parameter named `access_token`.

#### Example

To get all interventions on the source accessible by the token’s users, URL will looks like:

`https://[YOUR DOMAIN].api.digital.ringcentral.com/1.0/interventions?access_token=abc42`

### Via HTTP Header

In order to be compliant with OAuth 2.0 standards an access token can also be specified via the `Authorization` request header where value respects following format:

`Authorization: Bearer <access_token_value>`

#### Example

To get all interventions on the source accessible by the token’s users, you’ll need to build your request with Authorization request header with proper value. HTTP request will looks like:

```http
GET /1.0/interventions HTTP/1.1
Host: test.api.digital.ringcentral.com
Authorization: Bearer abc42
```

!!! warning "Keep access token secure"
    Do not publish your access token publicly. The access token is **unencrypted**, and possession of it by a third-party will give them access to your account. 



