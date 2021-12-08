# Extensions Report Fields

In RingCentral Engage Digital, we can enable extensions.

Each extension may append columns to existing csv or provide new exports.

## Customer Satisfaction(`Export::Responses`)

**Note**: Sometimes customers reply to the same survey several times, we take their last submission.

| Export name              | responses      |
|--------------------------|----------------|
| Incremental availability | Yes            |
| Time filtering           | Responded time |

| Column Name              | Description                                | Type     | Example                  | Comment                                                                                                                                                                                                                                                                                                                                    |
|--------------------------|--------------------------------------------|----------|--------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| responded_at             | Survey Response date                       | Datetime | 23/08/2013 12:15         |                                                                                                                                                                                                                                                                                                                                            |
| satisfaction_response_id | Provider response ID                       | String   | 42                       | Provider is Alchemer (formerly Survey Gizmo)                                                                                                                                                                                                                                                                                               |
| satisfaction_survey_id   | Provider survey ID                         | String   | 987654321                | Provider is Alchemer (formerly Survey Gizmo)                                                                                                                                                                                                                                                                                               |
| content_thread_id        | Content Thread  ID                         | ObjectId | 55cdaade77656222d500061d |                                                                                                                                                                                                                                                                                                                                            |
| intervention_id          | Intervention ID                            | ObjectId | 55dc83d677656254ca000a2e |                                                                                                                                                                                                                                                                                                                                            |
| intervention_categories  | Intervention categories                    | String   | SAV, Important, etc      |                                                                                                                                                                                                                                                                                                                                            |
| user_id                  | Agent ID                                   | ObjectId | 55cda687776562377900083d |                                                                                                                                                                                                                                                                                                                                            |
| user_name                | Agent name                                 | String   | Pierre Dupont            |                                                                                                                                                                                                                                                                                                                                            |
| identity_id              | Respondent ID                              | ObjectId | 55cda7bd7765622b58000735 |                                                                                                                                                                                                                                                                                                                                            |
| identity_name            | Respondent name                            | ObjectId | Paul Martin              | Sensitive                                                                                                                                                                                                                                                                                                                                  |
| main_indicator           | Selected answer value of the main_question | Integer  | 4                        |                                                                                                                                                                                                                                                                                                                                            |
| question_x               | Question x                                 | String   | Are you satisfied ?      | X is 1 to max number of survey question Be careful in case of automation there will be as may question_x answer_x than there are question in the survey(s)  This mean that this may change if the survey change                                                                                                                            |
| answer_x                 | Answer x                                   | String   | Yes, it’s awesome        | X is 1 to max number of survey question Be careful in case of automation there will be as many question_x answer_x than there are questions in the survey(s)  This means that this may change if the survey change                                                                                                                         |
| source_id                | Source ID                                  | ObjectId | 55cda687776562377900083d |                                                                                                                                                                                                                                                                                                                                            |
| source_name              | Source Name                                | String   | My Facebook Page         |                                                                                                                                                                                                                                                                                                                                            |
| source_type              | Source type                                | String   | Facebook                 |                                                                                                                                                                                                                                                                                                                                            |
| id                       | ED internal response ID                    | ObjectId | 61796662713368469afc0c18 |                                                                                                                                                                                                                                                                                                                                            |
| survey_id                | ED internal survey id                      | ObjectId | 61785f135bd54419e99ba270 |                                                                                                                                                                                                                                                                                                                                            |
| survey_type              | Survey provider type                       | String   | alchemer                 | Currently only Alchemer (formerly Survey Gizmo)                                                                                                                                                                                                                                                                                            |
| question_x_foreign_id    | Question x ID from provider                | String   | 14                       | X is 1 to max number of survey question Be careful in case of automation there will be as many question_x answer_x than there are questions in the survey(s)  This means that this may change if the survey change                                                                                                                         |
| answer_x_value           | Answer x internal value                    | String   | 5                        | Referred to as “reporting value” in Alchemer. Identical to answer_x for questions that allow custom text from customer. X is 1 to max number of survey question Be careful in case of automation there will be as many question_x answer_x than there are questions in the survey(s)  This means that this may change if the survey change |

## Survey

### Interventions Satisfactions(`Export::Interventions`)

The following columns will be added to Interventions csv.

This export can be incremental.

| Column Name              | Description              | Type     | Example / Comment |
|--------------------------|--------------------------|----------|-------------------|
| satisfaction_grade       | Satisfaction grade       | Integer  | 4                 |
| satisfaction_sent_at     | Satisfaction sent date   | Datetime | 07/10/2013 11:52  |
| satisfaction_answered_at | Satisfaction answer date | Datetime | 07/10/2013 11:52  |
| satisfaction_response_id | Gizmo Response ID        | String   | 729               |
| satisfaction_survey_id   | Gizmo Survey ID          | String   | 1284489           |

## Reply Assistant

### Knowledge Base(`Export::ReplyAssistantKnowledgeBaseVersions`)

| Export name              | reply_assistant-knowledge_base |
|--------------------------|--------------------------------|
| Incremental availability | No                             |
| Time filtering           | No                             |

Export of all knowledge base entries for reply assistant

| Column Name | Description      | Type     | Example / Comment                               |
|-------------|------------------|----------|-------------------------------------------------|
| foreign_id  | Entry identifier | String   | 4                                               |
| label       | Label            | String   | Contact Agent by private message                |
| number      | Version number   | Integer  | 1                                               |
| body        | Body             | Text     | Hi, please reach us by private message, Thanks. |
| language    | Language         | String   | en                                              |
| categories  | Categories       | Array    | Mobile, Simcard                                 |
| sources     | Sources          | Array    | My Facebook Page, My Twitter                    |
| group       | Group            | String   | Default Group                                   |
| id          | Version ID       | ObjectId | ID of the given object                          |
| entry_id    | Entry ID         | ObjectId | ID of the entry containing this version         |

### Entry Groups(`Export::ReplyAssistantEntryGroups`)

| Export name              | reply_assistant-entry_groups |
|--------------------------|------------------------------|
| Incremental availability | No                           |
| Time filtering           | No                           |

Export of all reply assistant entry groups

| Column Name  | Description                            | Type     | Example / Comment                                  |
|--------------|----------------------------------------|----------|----------------------------------------------------|
| created_at   | Creation Date                          | Datetime | 23/08/2013 12:15                                   |
| id           | ID                                     | ObjectId | 4f4f3a08a90ffb27ee000583                           |
| name         | Name of the group entry                | String   | Commercial help                                    |
| autocomplete | Enable autocomplete in reply assistant | Boolean  | true                                               |
| entry_count  | Entry count                            | Integer  | 12                                                 |
| entry_ids    | Entries IDs                            | Array    | 5240b1bca90ffbb6c7000006, 4340b1bca90debb6c7003333 |

### Sentence(`Export::ReplyAssistantSentenceVersions`)

| Export name              | reply_assistant-sentences |
|--------------------------|---------------------------|
| Incremental availability | No                        |
| Time filtering           | No                        |

Export of all sentences for reply assistant

| Column Name | Description    | Type    | Example / Comment            |
|-------------|----------------|---------|------------------------------|
| body        | Body           | Text    |                              |
| label       | Label          | String  |                              |
| language    | Language       | String  | en                           |
| number      | Version number | Integer | 1                            |
| sources     | Sources        | Array   | My Facebook Page, My Twitter |

## Sentiment Analysis

### Messages sentiment(`Export::Contents`)

The following columns will be added to [Messages csv](#messages-sentimentexportcontents).

This export can be incremental.

| Column Name    | Description         | Type | Example / Comment             |
|----------------|---------------------|------|-------------------------------|
| sentiment_text | Sentiment (as text) | Text | negative, neutral or positive |
