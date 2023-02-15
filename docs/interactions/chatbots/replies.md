# Replying to Messages

The common action to take after receiving a new message notification from a customer is to post a reply to that message. This document will help you understand how use the Engage REST API to post a reply to an incoming message. 

Replying to a message is as simple as composing an HTTP POST to the `contents` resource, with the following query and request parameters. 

### Query Parameters

| Parameter | Description |
|-|-|
| `access_token` | API identification token | 

### Request Parameters

| Parameter | Description |
|-|-|
| `author_id` | The identity id of content. This parameter is not mandatory, by default it use the token’s user first identity on source. |
| `body` | The content’s body. This parameter is mandatory. | 
| `in_reply_to_id` | The content’s id you want to reply to. If omitted, a new discussion will be created. If source does not support to initiate discussion this parameter is mandatory. |
| `private` | Created contents are public by default, set this parameter to `true` in order to create a private reply. |
| `source_id` | The source to create content to. If you specify `in_reply_to_id` parameter, source will be determined from. Otherwise, this parameter is mandatory. |

### Authorization

Only users that can reply or initiate a discussion on given source can call this API. The system will return an error if `in_reply_to` isn't synchronized or if `in_reply_to` content is not under an open intervention.

**Sample Request**

`curl -X POST -d "source_id=58ff349313047d2cf6f&in_reply_to_id=58ff349313047d2cf6f&&body=test de contenu" "https://[DOMAIN].api.engagement.dimelo.com/1.0/contents?access_token=yyyyyyyyyyy"`

**Sample Response**

```json
{
  "id":"73f1cb2938229d7fa222d096",
  "source_id":"d19c81948c137d86dac77216",
  "source_url":"http://domain-test.answers.dimelo.com/questions/42",
  "source_type":"answers",
  "thread_id":"26c56bc5b71c5193b6f8c656",
  "author_id":"4f0aa52d656a3d75867f784c",
  "creator_id":"ac24dc966bc7ecb74017c0cd",
  "foreign_id":"7789",
  "type":"question",
  "created_from":"synchronizer",
  "private_message":false,
  "status":"replied",
  "intervention_id":"7f946431b6eebffafae642cc",
  "language":"fr",
  "body":"Hello,\n\nHow to unlock my nokia 3210?\n\nThanks!",
  "body_formatted":{
    "text":"Hello,\n\nHow to unlock my nokia 3210?\n\nThanks!",
    "html":"<p>Hello,</p>\n\n<p>How to unlock my nokia 3210?</p>\n\n<p>Thanks!</p>"
  },
  "body_input_format":"text",
  "title":"Nokia 3210 unlocking",
  "attachments_count":1,
  "attachments":[
    {
      "id":"5464b5c04d61639684110000",
      "created_at":"2011-05-05T22:00:00Z",
      "updated_at":"2011-05-05T22:00:00Z",
      "content_type":"application/pdf",
      "size":174784,
      "filename":"sso.pdf",
      "foreign_id":"123",
      "url":"http://domain-test.engagement.dimelo.dev/attachments/5464b5c04d61639684110000"
    }
  ],
  "synchronization_status":"success",
  "category_ids":[
    "4d0fb475b242228033cbdf6d",
    "60944e5702bdafb74ac96141"
  ],
  "created_at":"2012-05-24T04:00:44Z",
  "updated_at":"2012-05-24T04:00:44Z",
  "approval_required":false,
  "remotely_deleted":false,
  "published":true
}
```

!!! note "Text Format Requirements"
    To this day Engage only supports messages in plain text format, Engage does not support enriched messages or HTML within the integration of a multi-channel bot. Engage is currently studying an extension to support enriched messages.
