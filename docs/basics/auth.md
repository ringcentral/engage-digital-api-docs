# Authenticate to the RingCX Digital API

Every API request must include an API access token. Send the token in the HTTP `Authorization` header using the Bearer authentication scheme:

```http
Authorization: Bearer <access-token>
```

For example:

```http
GET /1.0/interventions HTTP/1.1
Host: {account-name}.api.digital.ringcentral.com
Authorization: Bearer <access-token>
Accept: application/json
```

```bash
curl --request GET \
  --url "https://{account-name}.api.digital.ringcentral.com/1.0/interventions" \
  --header "Accept: application/json" \
  --header "Authorization: Bearer ${RINGCX_DIGITAL_ACCESS_TOKEN}"
```

Replace `{account-name}` with your RingCX Digital account name. Set `RINGCX_DIGITAL_ACCESS_TOKEN` in your local environment; do not place the token directly in source code.

!!! note "Permissions"
    RingCX Digital automatically associates each API access token with the account's default administrator user, which has all API permissions.

The API also accepts an `access_token` request parameter for compatibility with existing integrations. Use the Bearer header for new integrations because URLs can be stored in browser history, proxy logs, and server access logs.

To create a token, see [Obtain an API Access Token](access-token.md).
