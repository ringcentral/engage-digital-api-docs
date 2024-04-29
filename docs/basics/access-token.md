# Obtain an Access Token

Welcome to the RingCX Developer Platform where you can create amazing, automated, scalable customer experiences. Before you begin building your RingCX application, you will need to obtain an Access Token.

## What is an Access Token used for?

A RingCX Access Token is used to access the RingCX REST API. Developers transmit their access token with every request as a means of authentication.

The permissions bound to an access token are inherited from the associated user.

## How to obtain an Access Token?

A RingCX Access Token is provisioned through the RingCX user interface. Each token is assigned a description/label to help administrators keep better track of the tokens they have provisioned. In addition, each token is associated with a user within your RingCX account. When the token is used, all actions performed will be attributed to the associated user.

To obtain a token, follow these steps:

1. Login to your RingCX portal and click on the "Admin" menu located in the top, horizontal menu.
    
2. Select "API access tokens" towards the bottom of the left hand menu.
    
3. You should see a list of access tokens if any have been provisioned. Select the token, or click the "+" button to create a new one.
    
4. Finally, enter a label/description for your token, and select an Agent on which the token will act on behalf of. Make sure the token is "enabled" and click "Save."
    
      ![API access token](../img/api-token.png)

## How do I transmit the token to the RingCX REST API?

To learn more about how to authenticate to the RingCX REST API, read the article on [Authentication](../auth/).
