require 'faraday'

SERVER = "https://<YOUR-DOMAIN>.api.digital.ringcentral.com"
ACCESS_TOKEN = '<API-ACCESS-TOKEN>'
API = "/1.0/content_threads"

headers = {
  headers: { 'Authorization' => 'Bearer ' + ACCESS_TOKEN }
}
res = Faraday.new(SERVER + API, headers).get

puts res.body
