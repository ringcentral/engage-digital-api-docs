# Filters and Paginating Results

## Pagination

All API methods that returns a collection of records are paginated. Requests parameters and responses type always respect same format.

### Request parameters

API methods that supports pagination takes two (optional) parameters:

| Parameter | Description |
|-|-|
| `offset` | The record index to start. Default value is 0. |
| `limit` | The max number of records to return. Default value is 30, max value is 150. | 

### JSON response

```json
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

| Parameter | Description |
|-|-|
| `count` | This is equal to total records count. |
| `offset` | This is equal to given offset parameter. | 
| `limit` | This is equal to given limit parameter. |
| `records` | The array of returned records. |

## Searching & Filtering 

Most API methods that returns a collection of records can be filtered if they accept a q parameters. Be informed that using search in API is much more resource intensive and may be subject to rate limiting.

### Request parameters

API methods that supports search (filtering) can be used in the following way:

* With a `q` parameter containing a query string (equivalent to Engage Digital search interface)<br />Example: fetch all threads where source id equal `c157a` or `b12ec`:<br />
`/1.0/content_threads?q=source:”c157a”%20source:”b12ec”`
* With the search query passed as URL parameters Fetch all threads where source id equal `c157a` or `b12ec`: `/1.0/content_threads?source[]=c157a&source[]=b12ec`

Please refer to Engage Digital search API for all details about available query parameters.