## Overview
This my almost empty overview

with some *markup* in it



## Usage
Stuff that are complicated




## Errors
You can use the response template feature to override the default response provided by the policy.
These templates are be defined at the API level, in "Entrypoint" section for V4 Apis, or in "Response Templates" for V2 APIs.

The error keys sent by this policy are as follows:

| Key |
| ---  |
| API_KEY_MISSING |
| API_KEY_INVALID_KEY |




## As Yaml with comments
```yaml
# Error condition (string)
# The condition which will be verified to end the request (support EL).
errorCondition: "{#calloutResponse.status >= 400 and #calloutResponse.status <= 599}"
# Error status code (enum (string))
# HTTP Status Code send to the consumer if the condition is true
errorStatusCode: 500 # Possible values: "100" "101" "102" "200" "201" "202" "203" "204" "205" "206" "207" "300" "301" "302" "303" "304" "305" "307" "400" "401" "402" "403" "404" "405" "406" "407" "408" "409" "410" "411" "412" "413" "414" "415" "416" "417" "422" "423" "424" "429" "500" "501" "502" "503" "504" "505" "507" 
# Exit on error (boolean)
# Terminate the request if the error condition is true
exitOnError: 
# Fire & forget (boolean)
# Make the http call without expecting any response. When activating this mode, context variables and exit on error are useless.
fireAndForget: 
# HTTP Method (enum (string))
# HTTP method to invoke the endpoint.
method: GET # Possible values: "GET" "POST" "PUT" "DELETE" "PATCH" "HEAD" "CONNECT" "OPTIONS" "TRACE" 
# 
proxy: 
   # 
   # When enabled = false and useSystemProxy = false
   enabled:  # Possible values: false true 
   # 
   # When enabled = false and useSystemProxy = false
   useSystemProxy:  # Possible values: false true 
   # Proxy host (string)
   # Proxy host to connect to
   # When enabled = true and useSystemProxy = false
   host: proxy.acme.com
   # Proxy password (string)
   # Optional proxy password
   # When enabled = true and useSystemProxy = false
   password: "[redacted]"
   # Proxy port (integer)
   # Proxy port to connect to
   # When enabled = true and useSystemProxy = false
   port: 3524
   # Proxy Type (enum (string))
   # The type of the proxy
   # When enabled = true and useSystemProxy = false
   type: SOCKS5
   # Proxy username (string)
   # Optional proxy username
   # When enabled = true and useSystemProxy = false
   username: admin
# Scope (enum (string))
# Execute policy on <strong>request</strong> (HEAD) phase, <strong>response</strong> (HEAD) phase, <strong>request_content</strong> (includes payload) phase, <strong>response content</strong> (includes payload) phase.
scope: REQUEST # Possible values: "REQUEST" "RESPONSE" "REQUEST_CONTENT" "RESPONSE_CONTENT" 
# 
ssl: 
   # Verify Host (boolean)
   # Use to enable host name verification
   hostnameVerifier: true
   # 
   keyStore: 
      # 
      # When type = ''
      type:  # Possible values: "" "JKS" "PKCS12" "PEM" 
   # Trust all (boolean)
   # Use this with caution (if over Internet). The gateway must trust any origin certificates. The connection will still be encrypted but this mode is vulnerable to 'man in the middle' attacks.
   trustAll: 
   # 
   trustStore: 
      # 
      # When type = ''
      type:  # Possible values: "" "JKS" "PKCS12" "PEM" 
      # Password (string)
      # Truststore password
      # When type = 'JKS'
      password: "[redacted]"
      # Content (string)
      # Binary content as Base64
      # When type = 'JKS'
      content: |-
          --- BEGIN CERTIFICATE ---
      
          --- END CERTIFICATE ---
# Tags
# Some tags
tags: 
  # 
  - defaulted
  # 
  - and again
# URL (string)
url: http://localhost:8080/api
# Context variables
variables: 
   # Name (string)
   - name: field
     # Value (string)
     value: "{#jsonPath(#calloutResponse.content, '$.field')}"

```

## Environment variables


### 



####  Error condition
| | |
|---:|---|
|ENV| **GRAVITEE_ERRORCONDITION**|
|JVM|`-Dgravitee.errorcondition`|
|Default| `{#calloutResponse.status >= 400 and #calloutResponse.status <= 599}`|
The condition which will be verified to end the request (support EL).
<hr>


####  Error status code
| | |
|---:|---|
|ENV| **GRAVITEE_ERRORSTATUSCODE**|
|JVM|`-Dgravitee.errorstatuscode`|
|Default| `500`|
|Values| `100` `101` `102` `200` `201` `202` `203` `204` `205` `206` `207` `300` `301` `302` `303` `304` `305` `307` `400` `401` `402` `403` `404` `405` `406` `407` `408` `409` `410` `411` `412` `413` `414` `415` `416` `417` `422` `423` `424` `429` `500` `501` `502` `503` `504` `505` `507` |
HTTP Status Code send to the consumer if the condition is true
<hr>


####  Exit on error
| | |
|---:|---|
|ENV| **GRAVITEE_EXITONERROR**|
|JVM|`-Dgravitee.exitonerror`|
Terminate the request if the error condition is true
<hr>


####  Fire & forget
| | |
|---:|---|
|ENV| **GRAVITEE_FIREANDFORGET**|
|JVM|`-Dgravitee.fireandforget`|
Make the http call without expecting any response. When activating this mode, context variables and exit on error are useless.
<hr>


####  HTTP Method
| | |
|---:|---|
|ENV| **GRAVITEE_METHOD**|
|JVM|`-Dgravitee.method`|
|Default| `GET`|
|Values| `GET` `POST` `PUT` `DELETE` `PATCH` `HEAD` `CONNECT` `OPTIONS` `TRACE` |
HTTP method to invoke the endpoint.
<hr>


####  Scope
| | |
|---:|---|
|ENV| **GRAVITEE_SCOPE**|
|JVM|`-Dgravitee.scope`|
|Default| `REQUEST`|
|Values| `REQUEST` `RESPONSE` `REQUEST_CONTENT` `RESPONSE_CONTENT` |
Execute policy on <strong>request</strong> (HEAD) phase, <strong>response</strong> (HEAD) phase, <strong>request_content</strong> (includes payload) phase, <strong>response content</strong> (includes payload) phase.
<hr>


####  URL
| | |
|---:|---|
|ENV| **GRAVITEE_URL**|
|JVM|`-Dgravitee.url`|
|Default| `http://localhost:8080/api`|

<hr>


####  
| | |
|---:|---|
|ENV| **GRAVITEE_PROXY_ENABLED**|
|JVM|`-Dgravitee.proxy.enabled`|
|Values| `false` `true` |
|When| `enabled = false`  and `useSystemProxy = false` |

<hr>


####  
| | |
|---:|---|
|ENV| **GRAVITEE_PROXY_USESYSTEMPROXY**|
|JVM|`-Dgravitee.proxy.usesystemproxy`|
|Values| `false` `true` |
|When| `enabled = false`  and `useSystemProxy = false` |

<hr>


####  Proxy host
| | |
|---:|---|
|ENV| **GRAVITEE_PROXY_HOST**|
|JVM|`-Dgravitee.proxy.host`|
|Default| `proxy.acme.com`|
|When| `enabled = true`  and `useSystemProxy = false` |
Proxy host to connect to
<hr>


####  Proxy password
| | |
|---:|---|
|ENV| **GRAVITEE_PROXY_PASSWORD**|
|JVM|`-Dgravitee.proxy.password`|
|Default| `[redacted]`|
|When| `enabled = true`  and `useSystemProxy = false` |
Optional proxy password
<hr>


####  Proxy port
| | |
|---:|---|
|ENV| **GRAVITEE_PROXY_PORT**|
|JVM|`-Dgravitee.proxy.port`|
|Default| `3524`|
|When| `enabled = true`  and `useSystemProxy = false` |
Proxy port to connect to
<hr>


####  Proxy Type
| | |
|---:|---|
|ENV| **GRAVITEE_PROXY_TYPE**|
|JVM|`-Dgravitee.proxy.type`|
|Default| `SOCKS5`|
|When| `enabled = true`  and `useSystemProxy = false` |
The type of the proxy
<hr>


####  Proxy username
| | |
|---:|---|
|ENV| **GRAVITEE_PROXY_USERNAME**|
|JVM|`-Dgravitee.proxy.username`|
|Default| `admin`|
|When| `enabled = true`  and `useSystemProxy = false` |
Optional proxy username
<hr>


####  Verify Host
| | |
|---:|---|
|ENV| **GRAVITEE_SSL_HOSTNAMEVERIFIER**|
|JVM|`-Dgravitee.ssl.hostnameverifier`|
|Default| `true`|
Use to enable host name verification
<hr>


####  Trust all
| | |
|---:|---|
|ENV| **GRAVITEE_SSL_TRUSTALL**|
|JVM|`-Dgravitee.ssl.trustall`|
Use this with caution (if over Internet). The gateway must trust any origin certificates. The connection will still be encrypted but this mode is vulnerable to 'man in the middle' attacks.
<hr>


####  
| | |
|---:|---|
|ENV| **GRAVITEE_SSL_KEYSTORE_TYPE**|
|JVM|`-Dgravitee.ssl.keystore.type`|
|Values| `` `JKS` `PKCS12` `PEM` |
|When| `type = ''` |

<hr>


####  
| | |
|---:|---|
|ENV| **GRAVITEE_SSL_TRUSTSTORE_TYPE**|
|JVM|`-Dgravitee.ssl.truststore.type`|
|Values| `` `JKS` `PKCS12` `PEM` |
|When| `type = ''` |

<hr>


####  Password
| | |
|---:|---|
|ENV| **GRAVITEE_SSL_TRUSTSTORE_PASSWORD**|
|JVM|`-Dgravitee.ssl.truststore.password`|
|Default| `[redacted]`|
|When| `type = 'JKS'` |
Truststore password
<hr>


####  Content
| | |
|---:|---|
|ENV| **GRAVITEE_SSL_TRUSTSTORE_CONTENT**|
|JVM|`-Dgravitee.ssl.truststore.content`|
|Default| `--- BEGIN CERTIFICATE ---

--- END CERTIFICATE ---`|
|When| `type = 'JKS'` |
Binary content as Base64
<hr>


####  Tags
| | |
|---:|---|
|ENV| **GRAVITEE_TAGS_{index}**|
|JVM|`-Dgravitee.tags[{index}]`|
Some tags
<hr>


####  Tags
| | |
|---:|---|
|ENV| **GRAVITEE_TAGS_{index}**|
|JVM|`-Dgravitee.tags[{index}]`|
Some tags
<hr>


####  Name
| | |
|---:|---|
|ENV| **GRAVITEE_TAGS_{index}_VARIABLES_{index}_NAME**|
|JVM|`-Dgravitee.tags[{index}].variables[{index}].name`|
|Default| `field`|

<hr>


####  Value
| | |
|---:|---|
|ENV| **GRAVITEE_TAGS_{index}_VARIABLES_{index}_VALUE**|
|JVM|`-Dgravitee.tags[{index}].variables[{index}].value`|
|Default| `{#jsonPath(#calloutResponse.content, '$.field')}`|

<hr>



## Examples
*V4 Proxy API With Defaults*
```json
{
  "api": {
    "definitionVersion": "V4",
    "type": "PROXY",
    "name": "Test policy Example v4 API",
    "flows": [
      {
        "name": "Common Flow",
        "enabled": true,
        "selectors": [
          {
            "type": "HTTP",
            "path": "/",
            "pathOperator": "STARTS_WITH"
          }
        ],
        "request": [
          {
            "name": "Test policy",
            "enabled": true,
            "policy": "test",
            "configuration":
              {
                "errorCondition": "{#calloutResponse.status >= 400 and #calloutResponse.status <= 599}",
                "errorStatusCode": "500",
                "exitOnError": false,
                "fireAndForget": false,
                "method": "GET",
                "proxy": {
                  "enabled": false,
                  "useSystemProxy": false
                },
                "scope": "REQUEST",
                "ssl": {
                  "hostnameVerifier": true,
                  "keyStore": {
                    "type": ""
                  },
                  "trustAll": false,
                  "trustStore": {
                    "type": ""
                  }
                },
                "tags": [
                  "defaulted",
                  "and again"
                ],
                "url": "http://localhost:8080/api",
                "variables": [
                  {
                    "name": "field",
                    "value": "{#jsonPath(#calloutResponse.content, '$.field')}"
                  }
                ]
              }
          }
        ]
      }
    ]
  }
}

```
*CRD for V4 Proxy API With Defaults*
```yaml
apiVersion: "gravitee.io/v1alpha1"
kind: "ApiV4Definition"
metadata:
    name: "test-example-v4-gko-api"
spec:
    name: "Test policy Example V4 GKO API"
    type: "PROXY"
    flows:
      - name: "Common Flow"
        enabled: true
        selectors:
          - type: "HTTP"
            path: "/"
            pathOperator: "STARTS_WITH"
        request:
          - name: "Test policy"
            enabled: true
            policy: "test"
            configuration:
              errorCondition: '{#calloutResponse.status >= 400 and #calloutResponse.status <= 599}'
              errorStatusCode: "500"
              exitOnError: false
              fireAndForget: false
              method: GET
              proxy:
                enabled: false
                useSystemProxy: false
              scope: REQUEST
              ssl:
                hostnameVerifier: true
                keyStore:
                  type: ""
                trustAll: false
                trustStore:
                  type: ""
              tags:
                - defaulted
                - and again
              url: http://localhost:8080/api
              variables:
                - name: field
                  value: '{#jsonPath(#calloutResponse.content, ''$.field'')}'

```
