# Main Export Fields

## Threads(`Export::ContentThreads`)

- **Export name**: threads

- **Incremental availability**: Yes

- **Time filtering**: Creation time, Last content time

| Column Name                  | Description                                             | Type     | Example                                            | Comment                                       |
|------------------------------|---------------------------------------------------------|----------|----------------------------------------------------|-----------------------------------------------|
| id                           | ID                                                      | ObjectId | 523ffffb7aa58d1b6700000f                           |                                               |
| foreign_id                   | External ID                                             | String   | 102423                                             |                                               |
| source_id                    | Source ID                                               | ObjectId | 523ffffb7aa58d1b6700000f                           |                                               |
| source_type                  | Source type                                             | String   | Facebook                                           |                                               |
| source_name                  | Source Name                                             | String   | My Facebook Page                                   |                                               |
| created_at                   | Creation date                                           | Datetime | 23/09/2013 18:23                                   |                                               |
| updated_at                   | Update date                                             | Datetime | 23/09/2013 20:45                                   |                                               |
| closed                       | Thread locking status                                   | Boolean  | true                                               |                                               |
| first_categorization_at      | Date of first categorization                            | Datetime | 24/09/2013 19:00                                   |                                               |
| last_content_id              | ID of last message                                      | String   | 5240b1bca90ffbb6c7000006                           |                                               |
| last_content_at              | Date of last content                                    | Datetime | 24/09/2013 19:00                                   |                                               |
| title                        | Title                                                   | String   | How can I unlock my simcard ?                      | Sensitive                                     |
| contents_count               | Messages count                                          | Integer  | 12                                                 |                                               |
| all_categories               | Categories                                              | Array    | Mobile, Adsl, TV                                   |                                               |
| categories                   | Categories                                              | Array    | Mobile, Adsl, TV                                   | Contains only global categories of the thread |
| first_content_id             | ID of first message                                     | String   | 5240b1bca90ffbb6c7000006                           |                                               |
| first_content_author_id      | ID of first message author                              | String   | 5240b1bca90ffbb6c7000006                           |                                               |
| languages                    | Languages                                               | Array    | fr, en                                             |                                               |
| interventions_count          | Interventions count                                     | Integer  | 4                                                  |                                               |
| intervention_user_ids        | Agents IDs                                              | Array    | 5240b1bca90ffbb6c7000006, 4340b1bca90debb6c7003333 |                                               |
| opened_intervention_user_ids | Active agents IDs                                       | Array    | 5240b1bca90ffbb6c7000006, 4340b1bca90debb6c7003333 |                                               |
| ratings                      | Rating of individual messages in the thread, uniquified | Array    | 1, 3 , 4, 5                                        |                                               |

### Optional fields

Other custom data can be present in the thread export depending on the source type, which will result in additional columns in the export.

**Source Chat**

| Column Name            | Description              | Type                                     | Example                      | Comment                                                                                                                            |
|------------------------|--------------------------|------------------------------------------|------------------------------|------------------------------------------------------------------------------------------------------------------------------------|
| trigger_id             | Trigger ID               | ObjectId                                 | 54e1fb7077656269ea110200     | Sensitive                                                                                                                          |
| trigger_name           | Trigger name             | String                                   | Chat Button                  | Sensitive                                                                                                                          |
| close_cause            | Close cause              | String                                   | visitor_closed               | Sensitive Possible values : -client_impossible_delay -disconnection -intervention_closed -intervention_auto_closed -visitor_closed |
| closed_at              | Time of chat closing     | Datetime                                 | 8/5/2016 14:29:45            | Sensitive                                                                                                                          |
| page_title             | Page title               | String                                   | Activity and Sleep Wristband | Sensitive                                                                                                                          |
| page_url               | Page URL                 | String                                   | http://example.com/          | Sensitive                                                                                                                          |
| page_visit_count       | Page visit count         | Integer                                  | 4                            | Sensitive                                                                                                                          |
| page_visit_started_at  | Page visit start time    | Timestamp                                | 1470400259                   | Sensitive                                                                                                                          |
| visit_count            | Website visit count      | Integer                                  | 21                           | Sensitive                                                                                                                          |
| visit _started_at      | Website visit start time | Timestamp                                | 1470399681                   | Sensitive                                                                                                                          |
| {custom_variable_name} | Chat custom variable     | Depends on the chat custom variable type | 12                           | Sensitive Columns corresponding to the chat custom variables will also be available                                                |

## Message(`Export::Contents`)

- **Export name**: messages

- **Incremental availability**: Yes

- **Time filtering**: Creation time

| Column Name              | Description                                                   | Type     | Example                                                         | Comment                                                                                                              |
|--------------------------|---------------------------------------------------------------|----------|-----------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------|
| created_at               | Creation date                                                 | Datetime | 20/09/2013 13:15                                                | Creation date on the source                                                                                          |
| source_id                | Source ID                                                     | ObjectId | 193731f8a90ffbd442000b2a                                        |                                                                                                                      |
| source_type              | Source type                                                   | String   | Dimelo Answers                                                  |                                                                                                                      |
| source_name              | Source Name                                                   | String   | My Assistance                                                   |                                                                                                                      |
| content_thread_id        | Thread id                                                     | ObjectId | 512631f8a90ffbd442000037                                        |                                                                                                                      |
| type                     | Type                                                          | String   | Réponse publique                                                |                                                                                                                      |
| id                       | ID                                                            | ObjectId | abc631f8a9b24d442002397                                         |                                                                                                                      |
| private_message          | Private message ?                                             | Boolean  | true / false (or 1 / 0 for BI)                                  |                                                                                                                      |
| created_from             | Created from                                                  | String   | Synchronizer                                                    |                                                                                                                      |
| auto_submitted           | Auto submitted                                                | Boolean  | true / false (or 1 / 0 for BI)                                  | If true, it means the message is an automatic reply: ask a user to follow (twitter), send a survey to a customer,... |
| status                   | Status                                                        | String   | New                                                             | Possible values : - Assigned - Ignored - New - Replied - Agent message - Agent reply                                 |
| ignored_from             | Archived by                                                   | String   | Ice                                                             | Possible values :            - Ignored automatically - Archived by ICE - Manually ignored - Mass ignored             |
| categories               | Categories                                                    | Array    | Mobile, Adsl, TV                                                |                                                                                                                      |
| intervention_id          | Intervention id                                               | ObjectId | 512631f8a90ffbd442000037                                        |                                                                                                                      |
| initial_created_at       | Initial creation date                                         | Datetime | 20/09/2013 13:15                                                | Creation date in Dimelo database                                                                                     |
| creator_id               | Creator id                                                    | ObjectId | 512631f8a90ffbd442000037                                        | ID of the agent that created the message                                                                             |
| creator_name             | Creator name                                                  | String   | Jack Dupont                                                     | Name of the agent that created the message                                                                           |
| author_id                | Author ID                                                     | ObjectId | 512631f8a90ffbd442000037                                        | ID of the identity author of the message                                                                             |
| author_name              | Author Name                                                   | String   | Michelle Michu                                                  | Sensitive Name of the identity author of the message                                                                 |
| anonymized               | Right to be forgotten used                                    | Boolean  | true                                                            | Sensitive                                                                                                            |
| body                     | Message                                                       | Text     | Hi, I need help to unlock my sim card, how can I do it ?        | Sensitive                                                                                                            |
| body_as_text             | Message                                                       | Text     | Hi, I need help to unlock my sim card, how can I do it ?        | Sensitive                                                                                                            |
| body_as_html             | Message in HTML format                                        | Text     | <p>Hi, I need help to unlock my sim card, how can I do it ?</p> | Sensitive                                                                                                            |
| title                    | Title                                                         | String   | How can I unlock my simcard ?                                   | Sensitive                                                                                                            |
| foreign_categories       | External categories                                           | Array    | Orange mobile, Unlock Simcard,                                  | e.g. forum room for Lithium, Dimelo communities categories …                                                         |
| foreign_id               | External ID                                                   | String   | ‘1234’                                                          |                                                                                                                      |
| rating                   | Rating of the message (for reviews)                           | Integer  | 3                                                               |                                                                                                                      |
| published                | Message has been published                                    | Boolean  | true                                                            |                                                                                                                      |
| approval_required        | When user reply, his answers will need approval by supervisor | Boolean  | false                                                           |                                                                                                                      |
| remotely_deleted         | Message has been deleted on the source                        | Boolean  | false                                                           |                                                                                                                      |
| language                 | Detected language                                             | String   | fr                                                              | Format is ISO 369-1                                                                                                  |
| in_reply_to_id           | Source ID                                                     | ObjectId | 523ffffb7aa58d1b6700000f                                        |                                                                                                                      |
| in_reply_to_author_id    | Source ID                                                     | ObjectId | 523ffffb7aa58d1b6700000f                                        |                                                                                                                      |
| attachments_count        | Number of attachments                                         | Integer  | 1                                                               |                                                                                                                      |
| structured_reply_payload | Payload of the structured reply                               | String   | my_payload                                                      | Sensitive                                                                                                            |

## slJournal(`Exports::Events`)

- **Export name**: journal

- **Incremental availability**: Yes

- **Time filtering**: Creation time

| Column Name          | Description                                               | Type     | Example                                            | Comment                                     |
|----------------------|-----------------------------------------------------------|----------|----------------------------------------------------|---------------------------------------------|
| id                   | Event ID                                                  | ObjectId | 512631f8a90ffbd442000037                           |                                             |
| created_at           | Creation date                                             | Datetime | 20/09/2013 13:15                                   |                                             |
| user_id              | Agent id                                                  | ObjectId | 512631f8a90ffbd442000037                           |                                             |
| user_name            | Agent                                                     | String   | Adrien Jarthon                                     |                                             |
| name                 | Name                                                      | String   | content.replied                                    |                                             |
| message              | Message                                                   | Text     | Adrien Jarthon a répondu à Sébastien Luquet        | Sensitive                                   |
| content_thread_id    | Thread ID                                                 | ObjectId | 4f4f3a08a90ffb27ee000583                           |                                             |
| content_source _id   | Source ID                                                 | ObjectId | 4f4f3a08a90ffb27ee000583                           |                                             |
| intervention_id      | Intervention ID                                           | ObjectId | 4f4f3a08a90ffb27ee000583                           |                                             |
| content_id           | Content ID                                                | ObjectId | 4f4f3a08a90ffb27ee000583                           |                                             |
| category_ids         | Categories IDs                                            | Array    | 5240b1bca90ffbb6c7000006, 4340b1bca90debb6c7003333 |                                             |
| task_id              | Task ID                                                   | ObjectId | 4f4f3a08a90ffb27ee000583                           | Sensitive See the full list of events below |
| action               | Optional field describing the action leading to the event | String   | deferred                                           | Sensitive                                   |
| step                 | Task topology step when it expired                        | String   | realtime_default_target                            |                                             |
| rules_engine_rule_id | ID of the Rule which triggered the event                  | ObjectId | 5a4bbf3713047d6d123061d0                           | Sensitive                                   |
| entry_id             | Reply Assistant Entry ID                                  | ObjectId | 5a4bbf3713047d6d123061d0                           | Sensitive                                   |
| version_id           | Reply Assistant Version ID                                | ObjectId | 5a4bbf3713047d6d123061d0                           | Sensitive                                   |

List of events containing a `task_id`:
- Task assigned
- Task completed
- Task moved to queue
- Task timed out
- Task expired in an agent's inbox
- Task priority update from a rule
- Task priority update (only one that didn't have the task_id in it's - event)
- Task resumed by an agent
- Task manually assigned to an agent by a supervisor
- Task accepted
- Task transferred

## Sources(`Export::ContentSources`)

- **Export name**: sources

- **Incremental availability**: No

- **Time filtering**: Creation time

| Column Name                       | Description                       | Type     | Example                    | Comment                                                                                   |
|-----------------------------------|-----------------------------------|----------|----------------------------|-------------------------------------------------------------------------------------------|
| created_at                        | Creation date                     | Datetime | 23/08/2013 12:15           |                                                                                           |
| updated_at                        | Update Date                       | Datetime | 23/08/2013 12:15           |                                                                                           |
| id                                | ID                                | ObjectId | 4f4f3a08a90ffb27ee000583   |                                                                                           |
| active                            | Active                            | Boolean  |                            |                                                                                           |
| name                              | Name                              | String   | Twitter Search Domain Test |                                                                                           |
| status                            | Status                            | String   | ok, error, …               |                                                                                           |
| community_type                    | Community Type                    | String   | Dimelo Users               |                                                                                           |
| community                         | Community                         | String   | Domain Test                |                                                                                           |
| content_archiving                 | Content Archiving                 | Boolean  |                            |                                                                                           |
| auto_detect_content_language      | Auto Detect Content Language      | Boolean  |                            |                                                                                           |
| default_content_language          | Default Content Language          | String   | en                         |                                                                                           |
| content_signature                 | Content Signature                 | String   |                            |                                                                                           |
| hidden_from_stats                 | Hidden From Stats                 | Boolean  |                            |                                                                                           |
| default_task_priority             | Default task priority             | Integer  | 1                          |                                                                                           |
| channel_id                        | Channel ID                        | ObjectId | 55794bda416472d4e8060000   |                                                                                           |
| channel_name                      | Channel Name                      | String   | Chat                       |                                                                                           |
| sla_expired_strategy              | SLA Expired Strategy              | String   | Max                        |                                                                                           |
| sla_response                      | SLA Response                      | Integer  | 0                          |                                                                                           |
| intervention_messages_boost       | Intervention Messages Boost       | Integer  | 9                          |                                                                                           |
| private_messages_supported        | Private Messages Supported        | Boolean  |                            |                                                                                           |
| fb_app_id                         | Facebook App ID                   | String   |                            | Extra (facebook)                                                                          |
| fb_page_id                        | Facebook Page ID                  | String   |                            | Extra (facebook)                                                                          |
| synchronize_messages              | Synchronize Messages              | Boolean  |                            | Extra (facebook)                                                                          |
| synchronize_posts                 | Synchronize Posts                 | Boolean  |                            | Extra (facebook)                                                                          |
| synchronize_dark_posts            | Synchronize Dark Posts            | Boolean  |                            | Extra (facebook)                                                                          |
| page_id                           | Page ID                           | String   |                            | Extra (google plus)                                                                       |
| domain_name                       | Domain Name                       | String   |                            | Extra (ideas, answers)                                                                    |
| facebook_app_id                   | Facebook App ID                   | String   |                            | Extra (answers)                                                                           |
| facebook_page_id                  | Facebook Page ID                  | String   |                            | Extra (answers)                                                                           |
| email                             | Email                             | String   |                            | Extra (email)                                                                             |
| from_email                        | From Email                        | String   |                            | Extra (email)                                                                             |
| open_tracking_enabled             | Open Tracking Enabled             | Boolean  |                            | Extra (email)                                                                             |
| instagram_business_account_id     | Business Account ID (Instagram)   | String   |                            | Extra (instagram)                                                                         |
| abuse_only                        | Abuse Only                        | Boolean  |                            | Extra (lithium)                                                                           |
| url                               | URL                               | String   |                            | Extra (lithium, rightnow, tapatalk)                                                       |
| application                       | Application                       | String   |                            | Extra (mobile messaging)                                                                  |
| notification_generic_message      | Notification Generic Message      | String   |                            | Extra (mobile messaging)                                                                  |
| disable_notification_body_display | Disable Notification Body Display | Boolean  |                            | Extra (mobile messaging)                                                                  |
| from_name                         | Sender Alias                      | String   |                            | Extra (nexmo, chat, mobile messaging). Originally called inbound_number for Nexmo sources |
| phone_number                      | Sender Phone Number               | String   |                            | Extra (RingCentral SMS OTT)                                                               |
| ott_type                          | OTT Type                          | String   |                            | Extra (nexmo OTT)                                                                         |
| ott_account_id                    | OTT Account ID                    | String   |                            | Extra (nexmo OTT)                                                                         |
| base_uri                          | Base URI                          | String   |                            | Extra (SDK)                                                                               |
| readable_from_metas               | Readable From Metas               | Array    |                            | Extra (tapatalk)                                                                          |
| readable_from_ids                 | Readable From IDs                 | Array    |                            | Extra (tapatalk)                                                                          |
| twitter_id                        | Twitter ID                        | String   |                            | Extra (twitter)                                                                           |
| twitter_uuid                      | Twitter UUID                      | String   |                            | Extra (twitter)                                                                           |
| main_source_id                    | Main Source ID                    | ObjectId |                            | Extra (twitter search)                                                                    |
| searches                          | Searches                          | Array    |                            | Extra (twitter search)                                                                    |

## Identities(`Export::Identities`)

- **Export name**: identities

- **Incremental availability**: Yes

- **Time filtering**: Creation time, Update time

| Column Name                      | Description                   | Type     | Example                              | Comment                                                                 |
|----------------------------------|-------------------------------|----------|--------------------------------------|-------------------------------------------------------------------------|
| created_at                       | Creation date                 | Datetime | 23/08/2013 12:15                     |                                                                         |
| updated_at                       | Last update date              | Datetime | 23/08/2013 12:15                     |                                                                         |
| community_type                   | Community  type               | String   | Tapatalk                             |                                                                         |
| community                        | Community                     | String   | Forum LesMobiles                     |                                                                         |
| puppet                           | Controlled                    | Boolean  | true                                 |                                                                         |
| id                               | ID                            | ObjectId | 4f4f3a08a90ffb27ee000583             |                                                                         |
| uuid                             | UUID                          | String   | ‘1234’                               | Often same value as foreign_id                                          |
| foreign_id                       | External ID                   | String   | ‘1234’                               |                                                                         |
| screenname                       | Screen name                   | String   | Pierre Dupont                        | Sensitive                                                               |
| firstname                        | Firstname                     | String   | Pierre                               | Sensitive                                                               |
| lastname                         | Lastname                      | String   | Dupont                               | Sensitive                                                               |
| email                            | Email                         | String   | pierredupont@gmail.com               | Sensitive                                                               |
| home_phone                       | Home phone                    | String   | 0136656565                           | Sensitive                                                               |
| mobile_phone                     | Mobile phone                  | String   | 0636656565                           | Sensitive                                                               |
| address                          | Address                       | String   | 32 rue de Trévise                    | Sensitive                                                               |
| city                             | City                          | String   | Paris                                | Sensitive                                                               |
| anonymized                       | Right to be forgotten used    | Boolean  | true                                 | Sensitive                                                               |
| tags                             | Tags                          | Array    | VIP, Important                       |                                                                         |
| identity_group_id                | ID                            | ObjectId | 571648194379720b53000048             |                                                                         |
| company                          | Company name                  | String   | Apple                                | Sensitive From Identity Group                                           |
| emails                           | Emails                        | Array    | [steve.jobs@apple.com, steve@me.com] | Sensitive From Identity Group                                           |
| home_phones                      | Home Phones                   | Array    | [415-323-1232]                       | Sensitive From Identity Group                                           |
| mobile_phones                    | Mobile Phones                 | Array    | [415-321-2321]                       | Sensitive From Identity Group                                           |
| tw_followers_count               | People following the identity | Integer  | 543                                  | Extra (twitter) updated only when there is a new tweet from the user    |
| tw_following_count               | People the identity follows   | Integer  | 865                                  | Extra (twitter) updated only when there is a new tweet from the user    |
| tw_statuses_count                | Number of tweets              | Integer  | 1023                                 | Extra (twitter) updated only when there is a new tweet from the user    |
| tw_location                      | User's location               | String   | Paris                                | Extra (twitter) updated only when there is a new tweet from the user    |
| dimelo_type                      | DC Account type               | String   | facebook                             | Extra (DC)                                                              |
| dimelo_custom_field_1 (up to 10) | DC custom fields              | String   | 01 23 45 67 89                       | Extra (DC)                                                              |
| fb_bio                           | Facebook bio                  | String   | Painter living in Paris              | Extra (facebook) updated only when there is a new message from the user |
| fb_category                      | Facebook category             | String   | Computers/Internet                   | Extra (facebook) updated only when there is a new message from the user |
| fb_locale                        | Facebook locale               | String   | fr_FR                                | Extra (facebook) updated only when there is a new message from the user |
| mobile_device_info               | DM device info                | String   | motorola Nexus 6; 6.0                | Extra (DM) updated when the DM isinitialized by the host app            |
| mobile_authenticated             | DM is the user authenticated? | Boolean  | true                                 | Extra (DM)  updated when the DM is initialized by the host app          |
| ott_type                         | OTT provider                  | String   | wechat                               | Extra (wechat, line...)                                                 |

## IdentityGroup(`Export:IdentityGroup`)

- **Export name**: identity_groups

- **Incremental availability**: Yes

- **Time filtering**: Creation time, Update time

| Column Name   | Description      | Type     | Example                              | Comment                                           |
|---------------|------------------|----------|--------------------------------------|---------------------------------------------------|
| created_at    | Creation date    | Datetime | 23/08/2013 12:15                     |                                                   |
| updated_at    | Last update date | Datetime | 23/08/2013 12:15                     |                                                   |
| id            | ID               | ObjectId | 4f4f3a08a90ffb27ee000583             |                                                   |
| firstname     | Firstname        | String   | Steve                                | Sensitive                                         |
| lastname      | Lastname         | String   | Jobs                                 | Sensitive                                         |
| company       | Company name     |          | Apple                                | Sensitive                                         |
| gender        |                  |          | man                                  | Sensitive Possible values: - ‘man’ - ‘woman’ - ‘’ |
| emails        | Emails           | Array    | [steve.jobs@apple.com, steve@me.com] | Sensitive                                         |
| home_phones   | Home phones      | Array    | [415-323-1232]                       | Sensitive                                         |
| mobile_phones | Mobile phones    | Array    | [415-321-2321]                       | Sensitive                                         |
| notes         | Notes            | Text     | Iphone 6 IMEI                        | Sensitive                                         |
| tag_ids       | Tag Ids          | Array    | ["4ff17af67aa58d052300376d"]         | Array of tag ids associated                       |
| identity_ids  | Identity ids     | Array    | ["4fc636d30f4ca1670400000c"]         | Array of identity ids associated                  |

## Interventions(`Export::Interventions`)

- **Export name**: interventions

- **Incremental availability**: Yes

- **Time filtering**: Creation time, Update time

| Column Name                     | Description                                                                                | Type     | Example                  | Comment                                                                                                               |
|---------------------------------|--------------------------------------------------------------------------------------------|----------|--------------------------|-----------------------------------------------------------------------------------------------------------------------|
| created_at                      | Creation date                                                                              | Datetime | 23/08/2013 12:15         |                                                                                                                       |
| updated_at                      | Last update date                                                                           | Datetime | 23/08/2013 12:15         |                                                                                                                       |
| closed_at                       | Close Date                                                                                 | Datetime | 30/04/2016 15:10         | Date of last agent reply (=last_user_reply_at), and not when the intervention was closed by clicking the Solve button |
| closed_automatically            | Closed Automatically                                                                       | String   | chat_disconnected        |                                                                                                                       |
| source_id                       | Source ID                                                                                  | ObjectId | 4f4f3a08a90ffb27ee000583 |                                                                                                                       |
| source_type                     | Source  type                                                                               | String   | Tapatalk                 |                                                                                                                       |
| source_name                     | Source name                                                                                | String   | Forum LesMobiles         |                                                                                                                       |
| content_thread_id               | Thread id                                                                                  | ObjectId | 4f4f3a08a90ffb27ee000583 |                                                                                                                       |
| id                              | ID                                                                                         | ObjectId | 4f4f3a08a90ffb27ee000583 |                                                                                                                       |
| status                          | Status                                                                                     | String   | Opened                   | Possible values : - Opened - Closed                                                                                   |
| deferred_at                     | Deferred Datetime                                                                          | Datetime | 27/08/2013 08:32         | Date which intervention has been postponed to                                                                         |
| user_id                         | Agent ID                                                                                   | ObjectId | 4f4f3a08a90ffb27ee000583 |                                                                                                                       |
| user_name                       | Agent name                                                                                 | String   | Pierre Dupont            | Firstname Lastname                                                                                                    |
| user_replies_count              | Replies count                                                                              | Integer  | 3                        |                                                                                                                       |
| user_private_replies_count      | Private replies count                                                                      | Integer  | 1                        |                                                                                                                       |
| user_public_replies_count       | Public replies count                                                                       | Integer  | 2                        |                                                                                                                       |
| first_identity_content_id       | First Identity Content ID                                                                  | ObjectId | 2c1c3c08a90ffb27ee000113 |                                                                                                                       |
| first_user_reply_id             | First Agent Reply ID                                                                       | ObjectId | cc223b08a90ffb27ee001252 |                                                                                                                       |
| first_user_reply_at             | First reply date                                                                           | String   | 24/08/2013 00:08         |                                                                                                                       |
| last_user_reply_at              | Last reply date                                                                            | String   | 27/08/2013 08:32         |                                                                                                                       |
| last_user_reply_in              | Intervention duration (minutes or seconds according to configuration)                      | Integer  | 1400                     |                                                                                                                       |
| last_user_reply_in_bh           | Intervention duration in business hours (minutes or seconds according to configuration)    | Integer  | 1400                     |                                                                                                                       |
| first_user_reply_in             | First reply time (minutes or seconds according to configuration)                           | Integer  | 4                        |                                                                                                                       |
| first_user_reply_in_bh          | First reply time in business hours (minutes or seconds according to configuration)         | Integer  | 4                        |                                                                                                                       |
| handling_time                   | Total time spent in push on the task when active (seconds)                                 | Integer  | 87                       |                                                                                                                       |
| title                           | Message title                                                                              | String   | Where is my order ?      |                                                                                                                       |
| identity_id                     | Identity ID                                                                                | ObjectId | 4f4f3a08a90ffb27ee000583 |                                                                                                                       |
| identity_name                   | Identity name                                                                              | String   | pdupont                  |                                                                                                                       |
| identity_contents_count         | Identity messages count                                                                    | Integer  | 200                      |                                                                                                                       |
| identity_private_contents_count | Identity private messages count                                                            | Integer  | 30                       |                                                                                                                       |
| identity_public_contents_count  | Identity public messages count                                                             | Integer  | 170                      |                                                                                                                       |
| categories                      | Categories                                                                                 | Array    | simcard, mobile          |                                                                                                                       |
| comments_count                  | Comments count                                                                             | Integer  | 30                       |                                                                                                                       |
| user_reply_in_average           | Average agent reply time (minutes or seconds according to configuration)                   | Integer  | 3                        |                                                                                                                       |
| user_reply_in_average_bh        | Average agent reply time in business hours (minutes or seconds according to configuration) | Integer  | 2                        |                                                                                                                       |
| user_reply_in_average_count     | Number of exchanges between Identity -> User (used to weight user_reply)                   | Integer  | 3                        |                                                                                                                       |

## TaskSelectionActivity(`Export::TaskSelectionActivities`)

- **Export name**: task_selection_activities

- **Incremental availability**: Yes

- **Time filtering**: Begin date, End date

| Column Name     | Description              | Type     | Example                                            | Comment |
|-----------------|--------------------------|----------|----------------------------------------------------|---------|
| id              | TaskSelectionActivity ID | ObjectId | 4f4f3a08a90ffb27ee000583                           |         |
| created_at      | Creation date            | Datetime | 23/08/2013 12:15                                   |         |
| start           | Datetime                 | Datetime | 23/08/2013 12:15                                   |         |
| end             | Datetime                 | Datetime | 23/08/2013 12:15                                   |         |
| duration        | Duration                 | Integer  | 12                                                 |         |
| user_name       | Agent name               | String   | Pierre Dupont                                      |         |
| user_id         | Agent ID                 | ObjectId | 4f4f3a08a90ffb27ee000583                           |         |
| task_id         | Task ID                  | ObjectId | 4f4f3a08a90ffb27ee000583                           |         |
| category_ids    | Category IDs             | Array    | 5240b1bca90ffbb6c7000006, 4340b1bca90debb6c7003333 |         |
| intervention_id | Intervention ID          | ObjectId | 4340b1bca90debb6c7003333                           |         |
| source_id       | Source ID                | ObjectId | 4340b1bca90debb6c7003333                           |         |

## Presence Time

- **Export name**: presence_time

- **Incremental availability**: Yes

- **Time filtering**: Creation time

| Column Name               | Description                                                 | Type             | Example                  | Comment                                                                             |
|---------------------------|-------------------------------------------------------------|------------------|--------------------------|-------------------------------------------------------------------------------------|
| date                      | Date                                                        | Date             | 27/08/2013               |                                                                                     |
| user_id                   | Agent id                                                    | ObjectId         | 4f4f3a08a90ffb27ee000583 |                                                                                     |
| user_name                 | Agent                                                       | String           | Pierre Dupont            |                                                                                     |
| activity                  | Activity time                                               | Integer \| Float | 6,67                     | In seconds (Integer)  or hours (Float) depending on the “Export in seconds” setting |
| presence                  | Presence time                                               | Integer \| Float | 10,47                    | In seconds (Integer)  or hours (Float) depending on the “Export in seconds” setting |
| {channel}_available       | Time spent available                                        | Integer \| Float | 0                        | In seconds (Integer)  or hours (Float) depending on the “Export in seconds” setting |
| {channel}_away            | Time spent away                                             | Integer \| Float | 0,63                     | In seconds (Integer)  or hours (Float) depending on the “Export in seconds” setting |
| {channel}_busy            | Time spent busy                                             | Integer \| Float | 6,39                     | In seconds (Integer)  or hours (Float) depending on the “Export in seconds” setting |
| {channel}_full            | Time spent full                                             | Integer \| Float | 3,66                     | In seconds (Integer)  or hours (Float) depending on the “Export in seconds” setting |
| {channel}_unoccupied      | Time spent unoccupied                                       | Integer \| Float | 0,04                     | In seconds (Integer)  or hours (Float) depending on the “Export in seconds” setting |
| {channel}_active_total    | Time spent active (available + busy + full)                 | Integer \| Float | 10,05                    | In seconds (Integer)  or hours (Float) depending on the “Export in seconds” setting |
| {channel}_available_total | Time spent available (available + busy + full + unoccupied) | Integer \| Float | 10,09                    | In seconds (Integer)  or hours (Float) depending on the “Export in seconds” setting |
| {channel}_away_total      | Time spent away (away + custom away statuses)               | Integer \| Float | 1,72                     | In seconds (Integer)  or hours (Float) depending on the “Export in seconds” setting |
| away_{custom_status_id}   | Time spent under a custom away status                       | Integer \| Float | 1,10                     | In seconds (Integer)  or hours (Float) depending on the “Export in seconds” setting |

## Roles(`Export::Roles`)

- **Export name**: roles

- **Incremental availability**: No

- **Time filtering**: No

| Column Name               | Description                  | Type     | Example / Comment                                                                                                                                                                                                                                  |
|---------------------------|------------------------------|----------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| id                        | ID                           | ObjectId |                                                                                                                                                                                                                                                    |
| label                     | Label                        | String   | Administrateur, CM, ….                                                                                                                                                                                                                             |
| admin_stamp_answer        | Admin or author stamp answer | Boolean  |                                                                                                                                                                                                                                                    |
| access_advanced_analytics | View advanced analytics      | Boolean  |                                                                                                                                                                                                                                                    |
| assign_intervention       | Assign intervention          | Boolean  |                                                                                                                                                                                                                                                    |
| author_block_content      | Block users                  | Boolean  |                                                                                                                                                                                                                                                    |
| create_community          | Create a community           | Boolean  |                                                                                                                                                                                                                                                    |
| create_content_source     | Create a source              | Boolean  |                                                                                                                                                                                                                                                    |
| create_user               | Create an agent              | Boolean  |                                                                                                                                                                                                                                                    |
| delay_export_content      | Defer content publication    | Boolean  |                                                                                                                                                                                                                                                    |
| delete_content_thread     | Delete a thread              | Boolean  |                                                                                                                                                                                                                                                    |
| impersonate_user          | Impersonate agent            | Boolean  |                                                                                                                                                                                                                                                    |
| invite_user               | Invite an agent              | Boolean  |                                                                                                                                                                                                                                                    |
| manage_advanced_analytics | Manage advanced analytics    | Boolean  |                                                                                                                                                                                                                                                    |
| manage_categories         | Manage categories            | Boolean  |                                                                                                                                                                                                                                                    |
| manage_folders            | Manage messages folders      | Boolean  |                                                                                                                                                                                                                                                    |
| manage_identities         | Manage identities            | Boolean  |                                                                                                                                                                                                                                                    |
| manage_roles              | Manage roles                 | Boolean  |                                                                                                                                                                                                                                                    |
| manage_tags               | Manage identity tags         | Boolean  |                                                                                                                                                                                                                                                    |
| manage_teams              | Manage teams                 | Boolean  |                                                                                                                                                                                                                                                    |
| manage_users_of_my_teams  | Manage users of my team only | Boolean  | When false, the agent related permissions (create_user, impersonate_user, invite_user, read_user, update_user) are granted for all agents. When true, they’re granted only for agents belonging to at least one of the teams the agent is leading. |
| publish_content           | Publish contents             | Boolean  |                                                                                                                                                                                                                                                    |
| read_community            | View communities             | Boolean  |                                                                                                                                                                                                                                                    |
| read_content_source       | View sources                 | Boolean  |                                                                                                                                                                                                                                                    |
| read_event                | View journal on interface    | Boolean  |                                                                                                                                                                                                                                                    |
| read_export               | Export data                  | Boolean  |                                                                                                                                                                                                                                                    |
| read_own_stats            | View own stats               | Boolean  |                                                                                                                                                                                                                                                    |
| read_stats                | View all stats               | Boolean  |                                                                                                                                                                                                                                                    |
| read_user                 | View agents                  | Boolean  |                                                                                                                                                                                                                                                    |
| search_event              | View journal on admin        | Boolean  |                                                                                                                                                                                                                                                    |
| update_community          | Edit a community             | Boolean  |                                                                                                                                                                                                                                                    |
| update_content_source     | Edit a source                | Boolean  |                                                                                                                                                                                                                                                    |
| update_intervention       | Edit all interventions       | Boolean  |                                                                                                                                                                                                                                                    |
| update_own_intervention   | Edit own intervention        | Boolean  |                                                                                                                                                                                                                                                    |
| update_settings           | Configure application        | Boolean  |                                                                                                                                                                                                                                                    |
| update_user               | Edit an agent                | Boolean  |                                                                                                                                                                                                                                                    |

## Categories(`Export::Categories`)

- **Export name**: categories

- **Incremental availability**: No

- **Time filtering**: No

Export of all categories for a source or all sources

| Column Name        | Description        | Type     | Example / Comment            |
|--------------------|--------------------|----------|------------------------------|
| created_at         | Creation Date      | Datetime | 23/08/2013 12:15             |
| updated_at         | Update Date        | Datetime | 23/08/2013 12:15             |
| id                 | ID                 | ObjectId | 4f4f3a08a90ffb27ee000583     |
| level              | Level              | Integer  | 1                            |
| mandatory          | Mandatory          | Boolean  |                              |
| post_qualification | Post qualification | Boolean  |                              |
| multiple           | Multiple           | Boolean  |                              |
| name               | Name               | String   | Commercial                   |
| unselectable       | Not selectable     | Boolean  |                              |
| parent_id          | Parent ID          | ObjectId | 4ffe9e4e0f4ca1545c000004     |
| parent_name        | Parent Name        | String   | Categories                   |
| sources            | Sources            | Array    | My Facebook Page, My Twitter |

## Agents(`Export::Users`)

- **Export name**: agents

- **Incremental availability**: Yes

- **Time filtering**: Creation time

| Column Name     | Description     | Type     | Example / Comment        |
|-----------------|-----------------|----------|--------------------------|
| created_at      | Creation Date   | Datetime | 23/08/2013 12:15         |
| id              | ID              | ObjectId | 4f4f3a08a90ffb27ee000583 |
| firstname       | Firstname       | String   | Pierre                   |
| lastname        | Lastname        | String   | Dupont                   |
| nickname        | Nickname        | String   | pierrot                  |
| gender          | Gender          | String   | Man                      |
| email           | Email           | String   | pierredupont@gmail.com   |
| identities      | Identities      | Array    | pdupont, ...             |
| role            | Role            | String   | Administrateur           |
| teams           | Teams           | Array    | RC Paris, RC Toulouse    |
| categories      | Categories      | Array    |                          |
| signatures      | Signatures      | Array    |                          |
| foreign_id      | External ID     | String   | 427210                   |
| foreign_jwt_id  | Foreign JWT ID  | String   |                          |
| foreign_saml_id | Foreign SAML ID | String   |                          |
| enabled         | Enabled         | Boolean  |                          |

## Agents Notifications(`Export::UsersNotifications`)

- **Export name**: agents_notifications

- **Incremental availability**: No

- **Time filtering**: No

| Column Name                                 | Description                           | Type     | Example                  | Comment |
|---------------------------------------------|---------------------------------------|----------|--------------------------|---------|
| user_id                                     | Agent ID                              | ObjectId | 4f4f3a08a90ffb27ee000583 |         |
| user_name                                   | Agent Name                            | String   | Pierre Dupont            |         |
| content_imported_in_identity_tags           | New message from a tagged author      | Boolean  | true                     |         |
| content_imported_in_my_interventions        | New message in your interventions     | Boolean  | true                     |         |
| content_imported_in_my_sources              | New message in your sources           | Boolean  | false                    |         |
| content_requiring_approval                  | Message requiring approval            | Boolean  | true                     |         |
| content_thread_categorized_in_my_categories | Thread categorized in your categories | Boolean  | No                       |         |
| intervention_assigned                       | Intervention assigned                 | Boolean  | true                     |         |
| intervention_reactivated                    | Intervention is out of defer          | Boolean  | true                     |         |

## Agents Permissions(`Export::UsersPermissions`)

- **Export name**: agents_permissions

- **Incremental availability**: No

- **Time filtering**: No

This export displays user permissions by source.

| Column Name         | Description                                                   | Type     | Example / Comment        |
|---------------------|---------------------------------------------------------------|----------|--------------------------|
| user_id             | ID                                                            | ObjectId | 4f4f3a08a90ffb27ee000583 |
| user_name           | Agent                                                         | String   | Pierre Dupont            |
| source_id           | Source ID                                                     | ObjectId | 4f4f3a08a123af2a2e000583 |
| source_type         | Source Type                                                   | String   | Facebook                 |
| source_name         | Source Name                                                   | String   | My Facebook Page         |
| read                | User can read messages on source                              | Boolean  | true                     |
| reply               | User can reply messages on source                             | Boolean  | true                     |
| reply_with_html     | User can reply messages using html                            | Boolean  | true                     |
| initiate_discussion | User can Initiate discussion on source                        | Boolean  | true                     |
| approval_required   | When user reply, his answers will need approval by supervisor | Boolean  | true                     |
| destroy             | User can delete messages on source                            | Boolean  | true                     |

**Note**
A user may appear multiple times, for example a user may have different permissions on 2 different sources :

| user_id                  | user_name   | source_id                 | source_type    | source_name   | read | reply | reply_with_html | initiate_discussion | approva_required | destroy |
|--------------------------|-------------|---------------------------|----------------|---------------|------|-------|-----------------|---------------------|------------------|---------|
| 4f4f3a08a90ffb27ee000583 | Jack Dupont | 512631f8a90ffbd442000037  | Facebook       | My FB         | true | false | false           | false               | false            | false   |
| 4f4f3a08a90ffb27ee000583 | Jack Dupont | 57def21f8a90ffbd44266663f | Dimelo Answers | My Assistance | true | true  | true            | true                | false            | false   |

## Interventions Comments(`Export::InterventionsComments`)

- **Export name**: interventions_comments

- **Incremental availability**: Yes

- **Time filtering**: Creation time

| Column Name     | Description             | Type     | Example                                          | Comment   |
|-----------------|-------------------------|----------|--------------------------------------------------|-----------|
| created_at      | Creation Date           | Datetime | 23/08/2013 12:15                                 |           |
| body            | Body                    | String   | Jack, I reassigned to you as this is your domain | Sensitive |
| created_from    | Created From            | String   | interface                                        |           |
| intervention_id | Intervention ID         | ObjectId | 55cdaade77656222d500061d                         |           |
| id              | Intervention Comment ID | ObjectId | 55dc83d677656254ca000a2e                         |           |
| identity_id     | Identity ID             | ObjectId | 55cda9e6776562a644000001                         |           |
| identity_name   | Identity Name           | String   | Pierre Dupont                                    | Sensitive |
| source_id       | Source ID               | ObjectId | 55cda687776562377900083d                         |           |
| source_name     | Source Name             | String   | My Facebook Page                                 |           |
| thread_id       | Thread ID               | ObjectId | 55cda7bd7765622b58000735                         |           |
| user_id         | Agent Id                | ObjectId | 51f241de7aa58dd360006706                         |           |
| user_name       | Agent Name              | String   | Jean Dupuis                                      |           |

## Presence Statuses(`Export::PresenceStatus`)

- **Export name**: presence_status

- **Incremental availability**: No

- **Time filtering**: No

| Column Name | Description          | Type     | Example                  | Comment |
|-------------|----------------------|----------|--------------------------|---------|
| created_at  | Creation Date        | Datetime | 23/08/2013 12:15         |         |
| updated_at  | Update Date          | Datetime | 23/08/2013 12:15         |         |
| id          | Presence status  ID  | ObjectId | 55dc83d677656254ca000a2e |         |
| name        | Presence status Name | String   | Meeting                  |         |

## Attachments(`Export::Attachments`)

You can download the attachments by using an API access token with the following URL:
https://{your-domain}.engagement.dimelo.com/attachments/{attachment_id}?access_token={your_access_token}

- **Export name**: attachments

- **Incremental availability**: Yes

- **Time filtering**: Begin date, End date

| Column Name     | Description        | Type     | Example                                                                        | Comment                                        |
|-----------------|--------------------|----------|--------------------------------------------------------------------------------|------------------------------------------------|
| created_at      | Creation Date      | Datetime | 23/08/2013 12:15                                                               |                                                |
| id              | Attachment ID      | ObjectId | 55dc83d677656254ca000a2e                                                       |                                                |
| url             | Attachment Url     | String   | https://domain-test.engagement.dimelo.com/attachments/59a52f3aa5aacd03d93b3151 |                                                |
| content_id      | Content ID         | ObjectId | 55dc83d677656254ca000a2e                                                       |                                                |
| filename        | File Name          | String   | DJwCOLlW4AINC5g.jpg                                                            |                                                |
| embed           | Embed              | Boolean  | true                                                                           |                                                |
| virus_signature | Virus Signature    | String   | this is a virus                                                                |                                                |
| size            | Size               | Integer  | 126199                                                                         |                                                |
| type            | Type               | String   | image                                                                          |                                                |
| foreign_id      | Foreign ID         | ObjectId | 55dc83d677656254ca000a2e                                                       |                                                |
| public          | Attachment public  | Boolean  | true                                                                           |                                                |
| attachable_id   | Id of attachment   | ObjectId | 38dc82c677656254ca000a2e                                                       |                                                |
| attachable_type | type of attachment | string   | reply_assistant                                                                | Can be 'unknown', 'reply_assistant', 'content' |
