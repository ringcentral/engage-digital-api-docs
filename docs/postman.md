# Using Postman to Test RingCX Digital APIs

For easy testing using [Postman](https://www.getpostman.com/), RingCentral provides a Postman 2.0 Collection for RingCX Digital. It is based on the RingCentral RingCX Digital OpenAPI 3.0 Specificaion. While Postman can import an OpenAPI 3.0 Specification directly, RingCentral recommends using the Collection as it provides better authorization handling using Postman variables and environments as recommended by Postman.

The files are available here:

* [Postman 2.0 Collection](https://raw.githubusercontent.com/ringcentral/engage-digital-api-docs/master/specs/engage-digital_postman2.json)
* [OpenAPI 3.0 Specification](https://raw.githubusercontent.com/ringcentral/engage-digital-api-docs/master/specs/engage-digital_openapi3.yaml)

This document describes how to install and use the Postman 2.0 Collection.

## Pre-Requisites

Before you can use Postman, it is important to note the URL for your account and generate an access token. You will configure both of these in Postman.

### Your API URL

RingCX Digital accounts use a customer subdomain that is important to note when accessing both the web application and API. In the following URL, replace the `{mycompany}` variable in the web URL with your own subdomain when accessing your account.

`https://{mycompany}.digital.ringcentral.com`

Use your customer subdomain to create the following API base URL for your company:

`https://{mycompany}.api.digital.ringcentral.com`

### Creating an Access Token

See the following instructions for how to create an access token in the RingCX Digital administration web console.

* [Obtain an Access Token](https://developers.ringcentral.com/engage/guide/basics/access-token)

## Using Postman

Using Postman once you have your pre-requisites consists of a few steps:

1. Creating your Postman environment
2. Importing the Postman collection
3. Making an API call

### Configuring Your Postman Environment

Use the following steps to configure your Postman environment:

1. In Postman, create an environment by clicking the Gear icon for "Management Environments" in the upper right corner. This will bring up a list of existing environments.
2. Click "Add" to create a new environment.
3. Choose a name of your choice.
4. Enter your server URL with "Variable": `ENGAGE_DIGITAL_SERVER_URL` and "Initial Value" set to your API URL. This would be: `https://{mycompany}.api.digital.ringcentral.com` uising your company's subdomain replacing `{mycompany}`.
4. Enter your static access token with "Variable" set to `ENGAGE_DIGITAL_ACCESS_TOKEN` and "Initial Value" set to your access token. This will set the proper request header.
5. Click the "Add" button to finish adding this environment.

### Importing the Postman collection

Use the following steps to import the RingCX Digital Postman collection.

1. In the upper left corner of the Postman application click the "Import" button.
2. Click the "Import from Link" tab.
3. Paste in the following URL where it says "Enter a URL and press import": [`https://raw.githubusercontent.com/ringcentral/engage-digital-api-docs/master/specs/engage-digital_postman2.json`](https://raw.githubusercontent.com/ringcentral/engage-digital-api-docs/master/specs/engage-digital_postman2.json)
4. Click the "Import" button

### Making an API call

To test the Postman collection, let's call the "Get All Users" API.

1. In the Environments pick list in the upper right corner, select the environment you just created.
1. In the left hand navigation menu, select "Users" > "Getting all users"
1. Ensure all the optional query string parameters are unselected
1. Click the "Send" button

## Feedback

If you have any feedback on using the Postman collection, please post to our GitHub repo [here](https://github.com/ringcentral/engage-api-docs).
