# Using Postman to Test Engage Digital APIs

For easy testing using [Postman](https://www.getpostman.com/), RingCentral provides a Postman 2.0 Collection for Engage Digital. It is based on the RingCentral Engage Digital OpenAPI 3.0 Specificaion. While Postman can import an OpenAPI 3.0 Specification directly, RingCentral recommends using the Collection as it provides better management of access credentials using Postman variables and environments as recommended by Postman.

The files are available here:

* [Postman 2.0 Collection](http://ringcentral.github.io/api-specs/specs_engage_engage-digital_postman-2.0.json)
* [OpenAPI 3.0 Specification](http://ringcentral.github.io/api-specs/specs_engage_engage-digital_openapi-3.0.yaml)

This document describes how to install and use the Postman 2.0 Collection.

## Pre-Requisites

Before you can use Postman, it is important to note the URL for your account and generate an access token. You will configure both of these in Postman.

### Your API URL

Engager Digital accounts use a customer subdomain that is important to note when accessing both the web application and API. In the following URL, note `mycompany` in the web URL when you access your account. 

`https://mycompany.engagement.dimelo.com`

Use your customer subdomain to create the following API URL for your company:

`https://mycompany.api.engagement.dimelo.com`

### Creating an Access Token

See the following instructions for how to create an access token in the Engage Digital administration web console.

## Using Postman

Using Postman once you have your pre-requisites consists of a few steps:

1. Creating your Postman environment
2. Importing the Postman collection
3. Making an API call

### Configuring Your Postman Environment

Use the following steps to configure your Postman environment:

1. In Postman, create an environment by clicking the Gear icon for "Management Environments". This will bring up a list of existing environments.
2. lick "Add" to create a new environment.
3. Choose a name of your choice.
4. Enter your server URL with Variable: `RINGCENTRAL_ENGAGE_SERVER_URL` and Initial Value" set to your API URL, for example: `https://mycompany.api.engagement.dimelo.com`
4. Enter your static access token with the following Variable name `RINGCENTRAL_ENGAGE_ACCESS_TOKEN` and access token value. This will set the proper request header.

specs_engage_engage-digital.postman2.json

### Importing the Postman collection

Use the following steps to import the Engage Digital Postman collection.

1. In the upper left corner of the Postman application click the "Import" button.
2. Click the "Import from Link" tab.
3. Paste in the following URL where it says "Enter a URL and press import": `http://ringcentral.github.io/api-specs/specs_engage_engage-digital_postman-2.0.json`
4. Click the "Import" button

### Making an API call

To test the Postman collection, we will call the "Get All Users" API.

1. In the Environments pick list in the upper right corner, select the environment you just created.
1. In the left hand nav, select "Users" > "Getting all users"
1. Ensure all the optional query string parameters are unselected
1. Click the "Send" button.