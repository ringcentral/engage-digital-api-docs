# RingCX Digital API Responses

API responses use JSON. For example:

```json
{
  "title": "Hello World",
  "created_at": "2012-05-21T01:19:49Z",
  "available": true,
  "comments_count": 4
}
```

## Response format

Responses use UTF-8 encoding and the following content type:

```http
Content-Type: application/json; charset=utf-8
```

## Errors

API errors use a non-2xx HTTP status and the following JSON structure:

```json
{
  "error": "error_identifier",
  "message": "A description of the error",
  "status": 400
}
```

Common error statuses include:

| Status | Meaning |
|--------|---------|
| `400` | The request is malformed or contains an invalid parameter. |
| `403` | Authentication is required, or the token's user is not authorized to perform the operation. |
| `404` | The requested resource does not exist or is not accessible to the token's user. |
| `409` | The request conflicts with the current resource state. |
| `422` | The request is valid JSON but cannot be processed with the supplied values. |
| `429` | The applicable request limit has been exceeded. |

## Rate limits

The default account-level limit is 500 API requests per minute. An account can be configured with a different limit, and selected operations can have a token-specific limit. When a token-specific limit applies, requests using that token are counted separately for those operations.

When the applicable limit is exceeded, the API returns `429 Too Many Requests`. The error message states the limit that was applied:

```json
{
  "error": "rate_limit_exceeded",
  "message": "Rate limit exceeded (500 requests per minute max)",
  "status": 429
}
```

Applications should limit request concurrency, avoid unnecessary polling, and wait before retrying a rate-limited request.
