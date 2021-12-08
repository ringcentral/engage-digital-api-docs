# Events

Here’s a list of the different events available with the embedded UI:

## Pull View

#### `smcc:content:reply`

Triggered when a reply button is clicked

| Key                  | Type      | Example             |
|----------------------|-----------|---------------------|
| inReplyToId          | String    | 59709e1d367b6875d7  |
| inReplyToAuthorId    | String    | 59709e1d367b6875d7  |
| threadId             | String    | 59709e1d367b6875d7  |
| interventionId       | String    | 59709e1d367b6875d7  |
| firstIdentityId      | String    | 59709e1d367b6875d7  |
| sourceId             | String    | 59709e1d367b6875d7  |
| lastUpdatedContentAt | Timestamp | 1525856019.895      |
<br>

#### `smcc:content:initiate_discussion`

Triggered when initiating a new discussion, by clicking in the header menu or directly by adding the `new_content` param to the url.

| Key      | Type   | Example            |
|----------|--------|--------------------|
| body     | String | Hello              |
| sourceId | String | 59709e1d367b6875d7 |
<br>

#### `smcc:content:intervention_opened`

Triggered when clicking on `engage` to open an intervention.

| Key               | Type      | Example            |
|-------------------|-----------|--------------------|
| id                | String    | 59709e1d367b6875d7 |
| authorId          | String    | 59709e1d367b6875d7 |
| threadId          | String    | 59709e1d367b6875d7 |
| inReplyToAuthorId | String    | 59709e1d367b6875d7 |
| inReplyToId       | String    | 59709e1d367b6875d7 |
| interventionId    | String    | 59709e1d367b6875d7 |
| updatedAt         | Timestamp | 1525856019.895     |
| sourceId          | String    | 59709e1d367b6875d7 |
<br>

#### `smcc:content:intervention_closed`

Triggered when clicking on `solve` to close an intervention.

| Key               | Type      | Example            |
|-------------------|-----------|--------------------|
| id                | String    | 59709e1d367b6875d7 |
| authorId          | String    | 59709e1d367b6875d7 |
| threadId          | String    | 59709e1d367b6875d7 |
| inReplyToAuthorId | String    | 59709e1d367b6875d7 |
| inReplyToId       | String    | 59709e1d367b6875d7 |
| interventionId    | String    | 59709e1d367b6875d7 |
| updatedAt         | Timestamp | 1525856019.895     |
| sourceId          | String    | 59709e1d367b6875d7 |
<br>

#### `smcc:content_thread:selected`

Triggered when a content thread is selected

| Key             | Type   | Example            |
|-----------------|--------|--------------------|
| id              | String | 59709e1d367b6875d7 |
| lastContentId   | String | 59709e1d367b6875d7 |
| firstIdentityId | String | 59709e1d367b6875d7 |
<br>

#### `smcc:identity:show`

Triggered when clicking on the name of a client or its avatar.

| Key            | Type   | Example            |
|----------------|--------|--------------------|
| id             | String | 59709e1d367b6875d7 |
| threadId       | String | 59709e1d367b6875d7 |
| interventionId | String | 59709e1d367b6875d7 |
| taskID         | String | 59709e1d367b6875d7 |
<br>

## Task View

#### `smcc:task:selected`

Triggered when a task is selected (in task permalink view, it’ll be triggered once the task is loaded and then every time the iframe gain the focus)

| Key            | Type                        | Example            |
|----------------|-----------------------------|--------------------|
| Id             | String                      | 59709e1d367b6875d7 |
| contentId      | String                      | 59709e1d367b6875d7 |
| identityId     | String                      | 59709e1d367b6875d7 |
| interventionId | String                      | 59709e1d367b6875d7 |
| sourceId       | String                      | 59709e1d367b6875d7 |
| queue          | String [workbin \| history] | workbin            |
<br>

#### `smcc:task:ring`

Triggered when a task rings

| Key            | Type   | Example            |
|----------------|--------|--------------------|
| Id             | String | 59709e1d367b6875d7 |
| contentId      | String | 59709e1d367b6875d7 |
| identityId     | String | 59709e1d367b6875d7 |
| interventionId | String | 59709e1d367b6875d7 |
| sourceId       | String | 59709e1d367b6875d7 |
| queue          | String | global             |
| body           | String | bonjour            |
| author         | String | Martin Martin      |
<br>

#### `smcc:task:missed`

Triggered when a task is missed

| Key            | Type   | Example            |
|----------------|--------|--------------------|
| Id             | String | 59709e1d367b6875d7 |
| contentId      | String | 59709e1d367b6875d7 |
| identityId     | String | 59709e1d367b6875d7 |
| interventionId | String | 59709e1d367b6875d7 |
| sourceId       | String | 59709e1d367b6875d7 |
| queue          | String | global             |
<br>

#### `smcc:task:updated`

Triggered when a task is updated:

- User (agent) reply to Identity (customer) 
- Identity (customer) write to User (agent)
- Task is deferred and then retrieved
- Identity (customer) write to User (agent) but then delete content before agent reply 
- Content is deleted which leads to a different payload with some null values.

| Key                          | Type                                  | Example                                                                          |
|------------------------------|---------------------------------------|----------------------------------------------------------------------------------|
| Id                           | String                                | 59709e1d367b6875d7                                                               |
| channelId                    | String                                | 59709e1d367b6875d7                                                               |
| contentId                    | String (can be null)                  | 59709e1d367b6875d7                                                               |
| contentThreadId              | String                                | 59709e1d367b6875d7                                                               |
| identityId                   | String                                | 59709e1d367b6875d7                                                               |
| interventionId               | String (can be null)                  | 59709e1d367b6875d7                                                               |
| sourceId                     | String                                | 59709e1d367b6875d7                                                               |
| queue                        | String [workbin_<user_id> \| history] | workbin_56d4581f77656276e9000323                                                 |
| segmentIndex                 | Number                                | 1                                                                                |
| lastIdentityContentBody      | String                                | Hello World                                                                      |
| lastIdentityContentCreatedAt | String (can be null)                  | 2020-07-08 11:21:09 +0200                                                        |
| waitingForUser               | Boolean                               | false                                                                            |
| identitySignature            | String (can be null)                  | Anonyme                                                                          |
| identityAvatarUrl            | String (can be null)                  | https://domain-test.engagement.dimelo.com/assets/default_avatar/thumb-412009.png |
<br>

## Global Events

#### `smcc:user:disconnected`

Triggered when the user is disconnected from the application or on failed authentication when the experimental `Disable redirection to login when embedded` feature is enabled.