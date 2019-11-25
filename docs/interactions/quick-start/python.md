no_breadcrumb:true

# Interaction Python Quick Start

Welcome to the Engage Platform.

In this Quick Start, we are going to help you generate a list of threads that agents can engage customers through in just a few minutes. Let's get started.

## Obtain Access Key

The first thing you need to do is obtain an API Access Token if you do not already have one. Access tokens can be created and/or accessed via the "API access token" area in the "Admin" area of your Engage portal.

??? tip "How to generate an Engage API access token"

    1. Login to your Engage portal and click on the "Admin" menu located in the top, horizontal menu.

    2. Select "API access tokens" towards the bottom of the left hand menu.

    3. You should see a list of access tokens if any have been provisioned. Select the token, or click the "+" button to create a new one.

    4. Finally, enter a label/description for your token, and select an Agent on which the token will act on behalf of. Make sure the token is "enabled" and click "Save."

          ![API access token](../../../img/api-token.png)

Make note of the access token generated as you will need it later.

## Create and Edit threads.py

Create a file called `threads.py`. Be sure to edit the variables in &lt;ALL CAPS&gt; with your app's credentials.

```python
import requests

SERVER = "https://<YOUR-DOMAIN>.api.engagement.dimelo.com"
ACCESS_TOKEN = '<API-ACCESS-TOKEN>'
API = "/1.0/content_threads"

try:
    url = SERVER + API
    headers = {
            'Authorization': 'Bearer ' + ACCESS_TOKEN
            }
    try:
        response = requests.get(url, headers=headers)
        if response.status_code == 200:
            print (response._content)
        else:
            print (response.status_code)
    except Exception as e:
        raise ValueError(e)
except Exception as e:
    raise ValueError(e)
```

### Run Your Code

You are almost done. Now run your script.

```bash
$ python threads.py
```

## Need Help?

Having difficulty? Feeling frustrated? Receiving an error you don't understand? Our community is here to help and may already have found an answer. Search our community forums, and if you don't find an answer please ask!

<a target="_new" href="https://forums.developers.ringcentral.com/search.html?c=72&includeChildren=true&f=&type=question+OR+kbentry+OR+topic&redirect=search%2Fsearch&sort=newest&q=interactions">Search the forums &raquo;</a>

## What's Next?

When you have successfully made your first API call, it is time to take your next step towards building a more robust Engage application.

<a class="btn btn-success btn-lg" href="https://developers.ringcentral.com/engage/api-reference/">Explore the REST API &raquo;</a>
