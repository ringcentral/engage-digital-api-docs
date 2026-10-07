# RingCX API Requests

RingCX Digital provides a REST JSON API to retrieve, create, and manipulate data from third-party applications.

!!! tip "Using this Guide"
    Endpoint URLs use the `{account-name}` placeholder. Replace it with your RingCX Digital account name.

## Building an HTTP Request

### Base URL

Send API requests over HTTPS using your account-specific API hostname:

```text
https://{account-name}.api.digital.ringcentral.com
```

For example, if your account name is `example`, the base URL is:

```text
https://example.api.digital.ringcentral.com
```

Use the hostname assigned to your account if it differs from this format.

### Versioned path

Append the versioned path shown in the API reference. Most endpoints in this reference use the `/1.0` prefix:

```text
https://{account-name}.api.digital.ringcentral.com/1.0/users/me
```

Do not substitute one API version for another. Request and response contracts can differ between versions.

### HTTP method

Use the HTTP method shown for the operation in the API reference. Depending on the operation, the method can be `GET`, `POST`, `PUT`, `PATCH`, or `DELETE`.

### Headers

Send the API access token in the `Authorization` header. For requests with a JSON body, also send `Content-Type: application/json`.

```http
Accept: application/json
Authorization: Bearer <access-token>
Content-Type: application/json
```

### Array parameters

Some query and form parameters accept multiple values. Append `[]` to the parameter name and repeat the parameter for each value.

Examples:

* `?firstname=john&category_ids[]=4242&category_ids[]=2854`
* `?tag_ids[]=1&tag_ids[]=2`
* `?custom_field_values[multiple_custom_field_key][]=value1&custom_field_values[multiple_custom_field_key][]=value2`

## Authentication

See [Authenticate to the RingCX Digital API](auth.md).
