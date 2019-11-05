# Closing Threads

One option for handling a customer inquiry is to bring it to close. A developer does this by closing the associated thread, which will block the agent from replying again on the thread in question. One instance in which it makes sense to close the thread is when the user disconnects in the middle of the conversation, thus negating the agent's need to respond. To close a thread, following these steps. 

To close a thread, a developer will perform an HTTP PUT against the close resource. 

### Query Parameters

| Parameter | Description |
|-|-|
| `id` | The ID of the thread to close |
| `access_token` | API identification token | 

**Sample Request**

`curl -X PUT "https://[DOMAIN].api.engagement.dimelo.com/1.0/content_threads/9c9903dc3d559a6801ec544 1/close?access_token=yyyyyyyyyyy"`

**Sample Response**

```json
{
  "id":"9c9903dc3d559a6801ec5441",
  "source_id":"d19c81948c137d86dac77216",
  "title":"ADSL modem iss1ue",
  "interventions_count":1,
  "contents_count":4,
  "closed":false,
  "category_ids":[
    "4d0fb475b242228a32cbdf6d",
    "59248c4dae276a041cb296d2"
  ],
  "thread_category_ids":[
    "4d0fb475b242228a32cbdf6d"
  ],
  "extra_data":{
    "custom_my_number":123456,
    "trigger_id":"foo"
  },
  "foreign_id":"ab-2031",
  "created_at":"2012-05-18T14:24:44Z",
  "updated_at":"2012-05-21T18:10:12Z"
}
```

!!! note "Notes"
    * If token’s user does not have "read" permission on thread’s source, then a 404 HTTP response will be returned.
    * If the thread cannot be closed or if the user does not have the permission to close a thread, then a 403 HTTP response will be returned.
    * Closing a thread is an asynchronous process. Therefore, upon initiating the close thread process, the response from the server may reflect the thread having a `closed` status of false.
    * The `extra_data` field depends on the source type. For more informations, please refer to the exports documentation.




