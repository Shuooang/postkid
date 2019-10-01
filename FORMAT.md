# Postkid file format

Format is plain [Yaml](https://en.wikipedia.org/wiki/YAML), so no surprises here

## Method, Host and Path

The basics are defined in a following way, 

    * `method` can be one of normal HTTP methods
    * `host` is a hostname of remote host including protocol (!FIXME, this should be changed in order to support HTTP/2, HTTP/3, ..)
    * `path` is path part, will be escaped, so it is safe to pass any incompatible string there

```yaml
method: POST
host: https://example.com
path: matomo.php
```

## Query string

Is written like Yaml associative array, which is nice match with query
arguments. What is not supported is an existence of multiple values, because
there is a mismatch between yaml and query string

```yaml
query:
    _cvar: '{"1":["OS","iphone 5.0"],"2":["Matomo Mobile Version","1.6.2"],"3":["Locale","en::en"],"4":["Num Accounts","2"]}'
    action_name: View settings
    url: http://mobileapp.piwik.org/window/settings
    idsite: 8876
```

Query string is properly encoded and passed to created URL

## Headers

Again, usage of more values for the same header is not supported. Note that
values are transfered as-is, so it's better to use the same case as expected.
HTTP is case insensitive, but is common usage to have first characker in upper
case.

```yaml
header:
    Origin: http://example.com
    Content-Type: application/json
```

## Body

Sent only for POST/PUT/PATCH requests

```yaml
body: |
    {
        "some": {
            "nested" : {
                "json": "string"
                }
            }
        }
```
