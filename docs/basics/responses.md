# Engage API Responses

All responses are formatted in JSON, with the exception of a few errors. Here is an example:

```json
{
  "title": "Hello World",
  "created_at": "2012-05-21T01:19:49Z",
  "available": true,
  "comments_count": 4
}
```

### Encoding

All responses are formatted using UTF-8 encoding.

### Content type

The returned content-type is : `application/json; charset=utf-8`.

## Errors

In case of a fatal error, a response is sent in JSON (application/json; charset=utf-8 content type) with an HTTP status code different than 200.

All errors rendered respects following format:

```json
{
  "error": "Error identifier",
  "message": "A text message that describes the error",
  "status": 400
}
```

## Throttling

The number of queries is limited, the maximum is set to 500 queries per minute, otherwise you will hit the limit.

In case you reach the limit the server responds with 429 and the following JSON will be returned:

```json
{
  "error": "rate_limit_exceeded",
  "message": "Rate limit exceeded",
  "status": 429
}
```

