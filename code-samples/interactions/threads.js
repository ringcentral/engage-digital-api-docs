var https = require('https')

const SERVER = '<YOUR-DOMAIN>.api.engagement.dimelo.com'
const ACCESS_TOKEN = '<API-ACCESS-TOKEN>'
const API = "/1.0/content_threads"

var headers = {
        'Authorization': "Bearer " + ACCESS_TOKEN
    }

var options = {host: SERVER, path: API, method: 'GET', headers: headers};

var get_req = https.get(options, function(res) {
      var response = ""
      res.on('data', function (chunk) {
          response += chunk
        }).on("end", function(){
          if (res.statusCode == 200)
            console.log(response)
          else
            console.log(res.statusCode)
        });
    }).on('error', function(e) {
          console.log(e.message)
    });
