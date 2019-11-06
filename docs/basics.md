# Using the Engage REST API

Engage Digital provides a REST JSON API to retrieve, create data and manipulate data from third party applications. This document describes what can be done, under which conditions and the expected input and output formats and communication mechanism.

Note that API can easily be tested from any web browser or command line terminal.

## Building HTTP request

### Scheme and hostname

Request must be done with https scheme. Hostname is determined from your application name. If your application name is "example," then the API hostname will be: example.api.engagement.dimelo.com.

!!! tip "Using this Guide"
    Throughout this Developer Guide you will see sample endpoint URLs. These URLs all contain a placeholder string of `[YOUR DOMAIN]`. When developing, be sure to replace this string with your assigned domain.

### HTTP method

As it is specified in API methods list, HTTP method can be GET, POST, PUT or DELETE.

### URL path

All API paths is prefixed by `/1.0`. This is the version of Engage Digital Rest API.

### Access token

You **must** provide an access token on each of your API calls. Note that this access token must stay private. Don't publish it in a public website code. The access token is **unencrypted**.

You need to contact your project manager to have your own access token. Each access token is associated to an existing agent, all contents, interventions you make will be published as associated agent.

Note that some API methods requires authorization. It depends from the token’s user permissions. Authorization is described on all API methods below.

#### As parameter

Access token can be specified in request parameter named `access_token`.

#### As request header

In order to be compliant with OAuth 2.0 standards access token can also be specified with an Authorization request header where value respects following format:

`Authorization: Bearer <access_token_value>`

### Example

We suppose you have an Engage Digital instance accessible from https://[YOUR DOMAIN].engagement.dimelo.com and an access token with value abc42.

#### With access token as parameter

To get all interventions on the source accessible by the token’s users, URL will looks like: https://[YOUR DOMAIN].api.engagement.dimelo.com/1.0/interventions?access_token=abc42

#### With access token as request header

To get all interventions on the source accessible by the token’s users, you’ll need to build your request with Authorization request header with proper value. HTTP request will looks like:

```
GET /1.0/interventions HTTP/1.1
Host: test.api.engagement.dimelo.com
Authorization: Bearer abc42
```
### Multiple parameters

Some API methods described below take extra parameters. Some of them are m ultiple (example: category_ids, tags_ids or some custom fields). You must add double brackets [], after the parameters name.

Examples:

* `?firstname=john&category_ids[]=4242&category_ids[]=2854`
* `tag_ids[]=1&tag_ids[]=2`
* `custom_field_values[multiple_custom_field_key][]=value1&custom_field_values[multiple_custom_field_key][]=value2&custom_field_values[multiple_custom_field_key][]=value3`

### Uploading a file

The Attachments API allows you to upload file in order to use it later on in another API call (e.g. to create a content).
In order to upload a file to our API you need to pass the file parameter as multipart form data.

#### Using Curl

You can upload file via Curl by using the -F option with the path to your file, here’s an example:

`curl -X POST https://[YOUR DOMAIN].api.engagement.local.dimelo.info/1.0/attachments?access_token=ACCESS_TOKEN -F 'file=@path/to/your/file'`

#### Using Postman

Postman also allows you to upload files by doing adding a form-data parameter named file, and choose the file you want to upload:

<<POSTMAN_IMAGE>>

## Response

### JSON

All responses are formatted in J SON (except some errors). Here is an example:

```
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

### Errors

In case of a fatal error, a response is sent in JSON (application/json; charset=utf-8 content type) with an HTTP status code different than 200.

All errors rendered respects following format:

```
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

## User impersonation

All API requests accepts a `impersonated_user_id` parameter. This parameter allows you to overwrite the access token user by given user id.

Then all methods that use token’s user (interventions creation, contents creation, etc.) will use given impersonated user as user.

Note that if token’s user is unable to impersonate given user, an error will be rendered (400 HTTP code).

Example:

You have an access token with value `60576643bec4b6bd903232416ce5efad` associated to user « John Doe ». If you create an new intervention comment, it will be created as « John Doe » author.

`POST https://[YOUR DOMAIN].api.engagement.dimelo.com/1.0/intervention_comments?body=test&intervention_i d=c157a79031e1c40f85931829bc5fc552`

Then, if you want to create intervention comment as « Bill Murray », with id d3b07384d113edec49eaa6238ad5ff00, you just have to add impersonated_user_id parameter:

`POST https://[YOUR DOMAIN].api.engagement.dimelo.com/1.0/intervention_comments?body=test&intervention_id=c157a79031e1c40f85931829bc5fc552&impersonated_user_id=d3b07384d113edec49eaa6238ad5ff00`

## Pagination

All API methods that returns a collection of records are paginated. Requests parameters and responses type always respect same format.

## Request parameters

API methods that supports pagination takes two (optional) parameters:

* offset: The record index to start. Default value is 0.
* limit: The max number of records to return. Default value is 30, max value is 150.

## JSON response

```
{
  "count": 4320,
  "offset": 90,
  "limit": 30,
  "records": [
    {},
    {}
  ]
}
```

● count: This is equal to total records count.
● offset: This is equal to given offset parameter.
● limit: This is equal to given limit parameter.
● records: The array of returned records.

## Search & Filtering parameters

Most API methods that returns a collection of records can be filtered if they accept a q parameters. Be informed that using search in API is much more resource intensive and may be subject to rate limiting.

## Request parameters

API methods that supports search (filtering) can be used in the following way:

* With a `q` parameter containing a query string (equivalent to Engage Digital search interface)<br />Example: fetch all threads where source id equal `c157a` or `b12ec`:<br />
`/1.0/content_threads?q=source:”c157a”%20source:”b12ec”`
* With the search query passed as URL parameters Fetch all threads where source id equal `c157a` or `b12ec`: `/1.0/content_threads?source[]=c157a&source[]=b12ec`

Please refer to Engage Digital search API for all details about available query parameters