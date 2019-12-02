# Uploading Files

The Attachments API allows you to upload file in order to use it later on in another API call (e.g. to create a content). In order to upload a file to our API you need to pass the file parameter as multipart form data.

## Using Curl

You can upload file via Curl by using the -F option with the path to your file, here’s an example:

`curl -X POST https://[YOUR DOMAIN].api.engagement.local.dimelo.info/1.0/attachments?access_token=ACCESS_TOKEN -F 'file=@path/to/your/file'`

## Using Postman

Postman also allows you to upload files by adding a form-data parameter named `file`, and by choosing the file you want to upload as shown in screenshot below.

1. Select POST method
2. Select Body tab
3. Select 'form-data' option
4. Add key called "file"
5. Choose file off of file-system
6. Click "Send"

<img class="img img-fluid" src="../postman.png">

