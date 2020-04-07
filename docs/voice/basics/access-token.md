# Obtain an Access Token

There are a few ways to access Engage Voice APIs.s

1. RingCentral OAuth
1. Legacy OAuth
1. Legacy Perrmanent Token

## RingCentral OAuth

For customers with RingCentral Office, the approach to access RingCentral Engage APIs is to request a RingCentral Office accesss token and theen exchange it for an Exchange token.

### Retrieve RingCentral Office Access Token

First retrieve a RingCentral Officee access token using the following instructions:

https://developers.ringcentral.com/guide/authentication

### Retrive RingCentral Exchange Access Token

The call the following Exchange API to receive an Exchange Bearer access token:

```
POST 'https://engage.ringcentral.com/api/auth/login/rc/accesstoken
Content-Type: application/x-www-form-urlencoded

rcAccessToken=<rcOfficeAccessToken>&rcTokenType=Bearer
```

The response will contain the `accessToken` property that can be used in an API call. Take note of the `accountId` property as that will be used to make future API calls.

```json
{
  "refreshToken":null,
  "accessToken":"<rcEngageOfficeToken>",
  "tokenType":"Bearer",
  "agentDetails":[
    {
      "agentId":111111,
      "firstName":"John",
      "lastName":"Wang",
      "email":"john.wang@example.com",
      "username":"john.wang@example.com",
      "agentType":"AGENT",
      "rcUserId":222222,
      "accountId":"333333",
      "accountName":"RingForce",
      "agentGroupId":null,
      "allowLoginControl":true,
      "allowLoginUpdates":true
    }
  ],
  "adminId":1111,
  "adminUrl":"/voice/admin/",
  "agentUrl":"/voice/agent/",
  "ssoLogin":true
}
```

### Example Engage Voice API Call

The following is an example Enggage Voice API Call using a RingCentral Engage Access Token.

Of note, all Engage Voice APIs will be rooted at:

`https://engage.ringcentral.com/voice/api/`

```
GET https://engage.ringcentral.com/voice/api/v1/admin/users
Authorization: Bearer <rcEngageAccessToken>
```