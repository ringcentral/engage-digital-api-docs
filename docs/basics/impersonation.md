# User Impersonation

All API requests accepts a `impersonated_user_id` parameter. This parameter allows you to overwrite the access token user by given user id.

Then all methods that use token’s user (interventions creation, contents creation, etc.) will use given impersonated user as user.

Note that if token’s user is unable to impersonate given user, an error will be rendered (400 HTTP code).

### Example

You have an access token with value `60576643bec4b6bd903232416ce5efad` associated to user « John Doe ». If you create an new intervention comment, it will be created as « John Doe » author.

`POST https://[YOUR DOMAIN].api.engagement.dimelo.com/1.0/intervention_comments?body=test&intervention_i d=c157a79031e1c40f85931829bc5fc552`

Then, if you want to create intervention comment as « Bill Murray », with id d3b07384d113edec49eaa6238ad5ff00, you just have to add impersonated_user_id parameter:

`POST https://[YOUR DOMAIN].api.engagement.dimelo.com/1.0/intervention_comments?body=test&intervention_id=c157a79031e1c40f85931829bc5fc552&impersonated_user_id=d3b07384d113edec49eaa6238ad5ff00`
