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

Most API methods that returns a collection of records can be filtered if they accept a `q` parameters. Be informed that using search in API is much more resource intensive and may be subject to rate limiting.

### Rules & Examples

There are certain search rules:

1. Basic format: `{keyword}:{value}`.
2. Multiple search parameter paris are to be separated by spaces: `{keyword1}:{value1} {keyword2}:{value2}`.
3. If there's any space in {value}, double quote it: `{keyword}:"a value with spaces"`.
4. If you want to apply filter for the same keyword for multiple times, simply separate them: `{keyword}:{value1} {keyword}:{value2}`.
5. `text` is provided as the default keyword, so no need to specify it explicitly: `{keyword1}:{value1} text:{value}` = `{keyword1}:{value1} {value}`.
6. Negate a keyword with `-`: `-{keyword}:true` = `{keyword}:false`.
7. Boolean values can be either `true`(`1`, `on`, `yes`) or `false`(`0`, `off`, `no`).

### Try it with UI

Log in to Engage Digital and go to Agent Inbox. Available search parameters are all in search field. After search, they will be in tab url after `q=`.

<img class="img img-fluid" src="../search-api-ui.png">

### Try it with API calls

Endpoint: [GET] `https://{account_name}.api.{platform_hostname}/1.0/content_threads`

Auth: Bearer `{accessToken}` (refer to [Authentication](../basics/auth.md))

Query parameter: `q={filters}` (eg. `q=assigned_to_me:"true"`)

#### Parameters

| parameter | value | description |
|---|---|---|
| active_and_assigned | true or false | Restrict on assigned and active threads. Active threads have at least one intervention open that is not deferred. |
| active_and_assigned_to | user_id | Restrict on threads active and assigned to specific user id. Active threads have at least one intervention open that is not deferred. |
| active_and_assigned_to_disabled_agents | true | Restrict on threads active and assigned to disabled agents. Active threads have at least one intervention open that is not deferred. |
| active_and_assigned_to_me | true or false | Restrict on threads active and assigned to you. Active threads have at least one intervention open that is not deferred. |
| active_and_assigned_to_team | team_id | Restrict on threads active and assigned to given id. Negation is supported. Active threads have at least one intervention open that is not deferred. |
| assigned | true or false | Restrict on assigned (and not assigned) threads. Assigned threads are threads with at least one intervention (open or closed). |
| assigned_to | user_id | Restrict on thread assigned to specified user id. Assigned threads are threads with at least one intervention (open or closed). |
| assigned_to_me | true or false | Restrict on threads assigned to you (and not assigned to you). Assigned threads are threads with at least one intervention (open or closed). |
| assigned_to_team | team_id | Restrict on thread assigned to given team id. Negation is supported. |
| not_published | true or false | Restrict on threads that includes contents are not published. |
| approval_required | true or false | Restrict on threads that includes contents that require s approval. |
| categorized | true or false | Restrict on threads with at least one category or no category. |
| categorized_in | category_id | Restrict on category id. Many categories can be specified if you want to match at least one of specified categories. |
| categorized_in_mine | true or false | Restrict on your categories. It matches one of your categories. |
| categorized_in_mine_of | category_id | Restrict threads containing categories in at least one of your categories in given category id. Negation is not supported. |
| first_content_after | date | Restrict threads where their first content is created after given date. Negation is not supported. |
| first_content_before | date | Restrict threads where their first content is created before given date. Negation is not supported. |
| foreign | id | Restrict threads with the mentioned foreign_id. Negation is not supported. |
| forwarded_into | category_id | Restrict threads containing forwarded contents in at least one of given category id. Negation is supported. |
| forwarded_into_mine | true or false | Restrict threads containing forwarded contents in at least one of your categories. |
| forwarded_into_mine_of | category_id | Restrict threads containing forwarded categories in at least one of your categories in given category id. Negation is not supported. |
| language |  | Restrict threads by content language. |
| language |  | Restrict threads by content language. |
| last_content_after | date | Restrict threads where their last content is created after given date. Negation is not supported. |
| last_content_before | date | Restrict threads where their last content is created before given date. Negation is not supported. |
| needs_manual_categorization | true | Restrict threads having a content that needs manual categorization (meaning ICE could not classify a mandatory category) |
| not_categorized_in |  | Restrict threads that do not contain one or several specific categories |
| opened_and_assigned | true or false | Restrict on threads under open interventions. |
| opened_and_assigned_to | user_id | Restrict on thread under open intervention assigned to specified user id. |
| opened_and_assigned_to_disabled_agents | true | Restrict on threads under open interventions assigned to disabled agents. |
| opened_and_assigned_to_me | true or false | Restrict on threads under your open interventions. |
| opened_and_assigned_to_team | team_id | Restrict on threads under open interventions of given team id. Negation is supported. |
| status_in | name | Restrict on thread with least one of specified status. Accepted statuses are new (a freshly created), assigned (under intervention), forwarded, replied (replied by an agent), user_reply (a reply made by an agent), user_initiated (initiated by an agent) or ignored. Negation can be specified using a minus before keyword. |
| scheduled | true or false | Restrict on thread with a content scheduled for publication in the future. |
| text |words | Restrict on threads matching specified text (ignore case). Negation is not supported yet. Note that this is the default keyword. |
| source | id | Restrict on threads whose source has the given id. Negation is supported. If you specify many times this criteria, this will match any of specified source id. |
| identity | id | Restrict on threads that contains a message from given identity id. You can specify many ids. Negation is not supported. |
| identity_group | id | Restrict on threads that contains a message from given identity group id. You can specify many ids. Negation is not supported. |

### Additional Info

#### Date keywords

There are some keywords accepting date as value. The value can be given in a specific date format, or it can be a human readable string. When ambiguous, date interpretation will depend on the application locale, as follows:

In `French`: `03/04/2011` will be interpreted as `the third day of the fourth month`
In `English`: `03/04/2011` will be interpreted as `the fourth day of the third month`
But `03/17/2011` will be interpreted as `the 17th day of the third month` in `both locales`

Here are some examples of dates or strings that can be interpreted as dates:

- yesterday
- last month
- 2011-02-03
- 2011-02-03 14:27:33
- 2 hours ago
- 2 weeks ago at 4pm
- last friday at 20:00
- 4th day last week
- dec 25
- january 5 at 4

Business hours are the hours during the day in which business is commonly conducted. Typical business hours vary by domain.

Here are some examples of dates or strings that can be interpreted as business hours dates:

- 6 business hours ago
- 20 business minutes from now
- 1 business day ago

If a date without a time is given, the time will be considered as `00:00:00`.

#### Sorting results

Returned results can be ordered by using `order` keyword. At the moment, ordering only supports `created_at` and `last_content_at` with `asc` or `desc`.

An example: `q={keyword}:{value} order:created_at.asc` (the default order is `last_content_at.desc`)