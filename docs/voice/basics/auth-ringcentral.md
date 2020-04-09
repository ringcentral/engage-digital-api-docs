# Obtaining an Access Token using RingCentral Office

For customers with RingCentral Office, the approach to access RingCentral Engage APIs is to request a RingCentral Office accesss token and then exchange it for an Engage token to be used with Engage APIs.

> Note: Engage Voice APIs for Office customers are rooted at:
> 
> `https://engage.ringcentral.com/voice/api/`

## Retrieve RingCentral Office Access Token

First retrieve a RingCentral Office access token using the following instructions:

https://developers.ringcentral.com/guide/authentication

## Retrieve RingCentral Engage Access Token

Once you have a RingCentral Office Access Token, call the following Engage API to receive an Engage Bearer access token:

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

## Example Engage Voice API Call

The following is an example Engage Voice API Call using a RingCentral Engage Access Token.

```
GET https://engage.ringcentral.com/voice/api/v1/admin/users
Authorization: Bearer <rcEngageAccessToken>