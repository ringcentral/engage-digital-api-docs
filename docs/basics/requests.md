# RingCX API Requests

RingCX Digital provides a REST JSON API to retrieve, create data and manipulate data from third party applications.

!!! tip "Test in your browser"
    The RingCX API can easily be tested from any web browser or command line terminal.

!!! tip "Using this Guide"
    Throughout this Developer Guide you will see sample endpoint URLs. These URLs all contain a placeholder string of `[YOUR DOMAIN]`. When developing, be sure to replace this string with your assigned domain.

## Building an HTTP Request

### Scheme and hostname

Request must be done with https scheme. Hostname is determined from your application name. If your application name is "example", then the API hostname will be: example.api.engagement.dimelo.com.

### HTTP method

As it is specified in API methods list, HTTP method can be GET, POST, PUT or DELETE.

### URL path

All API paths is prefixed by `/1.0`. This is the version of RingCX Digital Rest API.

#### Multiple parameters

Some API methods described below take extra parameters. Some of them are multiple (example: category_ids, tags_ids or some custom fields). You must add double brackets [], after the parameters name.

Examples:

* `?firstname=john&category_ids[]=4242&category_ids[]=2854`
* `tag_ids[]=1&tag_ids[]=2`
* `custom_field_values[multiple_custom_field_key][]=value1&custom_field_values[multiple_custom_field_key][]=value2&custom_field_values[multiple_custom_field_key][]=value3`

## Authentication

See [Authenticating to the RingCX API](../auth/)
