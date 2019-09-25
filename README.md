# Postkid

Small and lean commandline oriented HTTP request builder

## Problem

As a developer I want to call HTTP API via `curl`, however manually writing
curl command is cumbersome. One need to properly escape things for shell as
well as for web.

## Postkid

Specify your HTTP request in nice plain Yaml file and let computers to do the work

Here is totally fabricated example to showcase it, see `examples/` directory for more

```yaml
method: POST
host: https://example.com
path: matomo.php
query:
    _cvar: '{"1":["OS","iphone 5.0"],"2":["Matomo Mobile Version","1.6.2"],"3":["Locale","en::en"],"4":["Num Accounts","2"]}'
    action_name: View settings
    url: http://mobileapp.piwik.org/window/settings
    idsite: 8876
header:
    Origin: http://example.com
    Content-Type: application/json
body: |
    {
        "some": {
            "nested" : {
                "json": "string"
                }
            }
        }
```

```sh
$ go get github.com/vyskocilm/postkid
$ postkid example.yaml
curl -XPOST -H 'Content-Type: application/json' -H 'Origin: http://example.com' --data '{
    "some": {
        "nested" : {
            "json": "string"
            }
        }
    }' 'https://example.com/matomo.php?idsite=8876&_cvar=%7B%221%22%3A%5B%22OS%22%2C%22iphone+5.0%22%5D%2C%222%22%3A%5B%22Matomo+Mobile+Version%22%2C%221.6.2%22%5D%2C%223%22%3A%5B%22Locale%22%2C%22en%3A%3Aen%22%5D%2C%224%22%3A%5B%22Num+Accounts%22%2C%222%22%5D%7D&action_name=View+settings&url=http%3A%2F%2Fmobileapp.piwik.org%2Fwindow%2Fsettings' 
```

## Run curl

The `run-curl` switch will execute `curl` command and print the standard ouptut

```sh
postkid -run-curl examples/httpbin.yml
{
  "args": {
    "a+b": "a&b", 
    "foo": "bar"
  }, 
  "headers": {
    "Accept": "*/*", 
    "Content-Type": "application/json", 
    "Host": "httpbin.org", 
    "User-Agent": "curl/7.66.0"
  }, 
  "origin": "217.30.65.6, 217.30.65.6", 
  "url": "https://httpbin.org/get?foo=bar&a%2Bb=a%26b"
}
```

## License

MIT

## Naming

Well, not so surpsising

* Postman - https://www.getpostman.com/ impressive and well designed tool, however closed source and requires registration
* Postwoman.io - https://postwoman.io nice online tool, can generate XHR, Fetch and curl, however I do preffer CLI tools - at also don't escape parameters

So in between Postman and Postwoman, there is a place for Postkid - small simple easy to work with command line tool without Electron based bells & whistles.
