This should stay here
<!-- GENERATED CODE - DO NOT ALTER THIS OR THE FOLLOWING LINES -->
# Test policy

## Overview
This my almost empty overview

with some *markup* in it



## Usage
Stuff that are complicated




## Errors
You can use the response template feature to override the default response provided by the policy.
These templates are be defined at the API level, in "Entrypoint" section for V4 Apis, or in "Response Templates" for V2 APIs.

The error keys sent by this policy are as follows:

| Key| Parameters |
| --- | ---  |
| API_KEY_MISSING| - |
| API_KEY_INVALID_KEY| - |



## Phases
The phases checked below are supported by the `test` policy:

| v2 Phases| Compatible?| v4 Phases| Compatible? |
| --- | --- | --- | ---  |
| onRequest|  | onRequest| ✅ |
| onResponse|  | onResponse| ✅ |
| onRequestContent| ✅| onMessageRequest|   |
| onResponseContent| ✅| onMessageResponse|   |


## Compatibility matrix
Strikethrough line are deprecated versions

| Plugin version| APIM| AM| Cockpit| Comment |
| --- | --- | --- | --- | ---  |
|~~1.0~~|~~3.x~~|~~2.1~~|~~-~~|~~-~~ |
|2.x|4.x|4.1|-|Incompatible with cloud |
|3.x|4.6 and above|4.6 and above|-|- |



## Configuration
### Gateway configuration
You need to configure Gravitee as follows

gravitee.yml
```YAML
secrets:
  kubernetes:
    enable: true
```
values.yaml
```YAML
gateway:
  secrets:
    kubernetes:
      enable: true
```
Warning: this some heavy doc


### Configuration options


#### 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Request body<br>`body`| string|  | |  |  | |
| Constraint test bed<br>`constraints`| object|  | |  |  | For the sake of testing<br/>See "Constraint test bed" section|
| Error condition<br>`errorCondition`| string|  | `{#calloutResponse.status >= 400 and #calloutResponse.status <= 599}`| ✅|  | The condition which will be verified to end the request (support EL).|
| Error response body<br>`errorContent`| string|  | |  |  | The body response of the error if the condition is true (support EL)|
| Error status code<br>`errorStatusCode`| enum (string)|  | `500`|  |  | HTTP Status Code send to the consumer if the condition is true<br>Values:`100` `101` `102` `200` `201` `202` `203` `204` `205` `206` `207` `300` `301` `302` `303` `304` `305` `307` `400` `401` `402` `403` `404` `405` `406` `407` `408` `409` `410` `411` `412` `413` `414` `415` `416` `417` `422` `423` `424` `429` `500` `501` `502` `503` `504` `505` `507` |
| Exit on error<br>`exitOnError`| boolean| ✅| |  |  | Terminate the request if the error condition is true|
| Fire & forget<br>`fireAndForget`| boolean|  | |  |  | Make the http call without expecting any response. When activating this mode, context variables and exit on error are useless.|
| Request Headers<br>`headers`| array|  | |  |  | <br/>See "Request Headers" section|
| HTTP Method<br>`method`| enum (string)| ✅| `GET`|  |  | HTTP method to invoke the endpoint.<br>Values:`GET` `POST` `PUT` `DELETE` `PATCH` `HEAD` `CONNECT` `OPTIONS` `TRACE` |
| Proxy Options<br>`proxy`| object|  | |  |  | <br/>See "Proxy Options" section|
| Scope<br>`scope`| enum (string)|  | `REQUEST`|  |  | Execute policy on <strong>request</strong> (HEAD) phase, <strong>response</strong> (HEAD) phase, <strong>request_content</strong> (includes payload) phase, <strong>response content</strong> (includes payload) phase.<br>Values:`REQUEST` `RESPONSE` `REQUEST_CONTENT` `RESPONSE_CONTENT` |
| SSL Options<br>`ssl`| object|  | |  |  | <br/>See "SSL Options" section|
| Tags<br>`tags`| array<br>[1, 3], unique|  | `[defaulted]`|  |  | Some tags|
| URL<br>`url`| string<br>[1, 150]| ✅| | ✅| ✅| |
| Context variables<br>`variables`| array|  | |  |  | <br/>See "Context variables" section|


#### Constraint test bed (Object)
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Description<br>`description`| string<br>[0, 1000]|  | |  |  | |
| Lower bounds<br>`lowerBounds`| number<br>(1, 5]|  | |  |  | |
| Middle bounds<br>`middleBounds`| number<br>[5, 10.333333333333]|  | |  |  | |
| Open lower bound<br>`openLowerBounds`| number<br>[-Inf, 10]|  | |  |  | |
| Open upper bound<br>`openUpperBounds`| number<br>[10, +Inf]|  | |  |  | |
| Ratio<br>`ratio`| number<br>[0, 1]|  | |  |  | |
| Upper bounds<br>`upperBounds`| number<br>[10.25, 15)|  | |  |  | |


#### Request Headers (Array)
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Name<br>`name`| string|  | |  |  | |
| Value<br>`value`| string|  | | ✅| ✅| |


#### Proxy Options (OneOf)
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Enabled<br>`enabled`| object| ✅| |  |  | <br/>See "Enabled" sectionEnabled of Proxy Options<br>Values:`""` `true` `true` |
| Use System Proxy<br>`useSystemProxy`| object| ✅| |  |  | <br/>See "Use System Proxy" sectionUse System Proxy of Proxy Options<br>Values:`""` `true` `""` |


#### Proxy Options: No proxy `enabled = false` `useSystemProxy = false` 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| No properties | | | | | | | 

#### Proxy Options: Use proxy configured at system level `enabled = true` `useSystemProxy = true` 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| No properties | | | | | | | 

#### Proxy Options: Use proxy for client connections `enabled = true` `useSystemProxy = false` 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Proxy host<br>`host`| string| ✅| |  |  | Proxy host to connect to|
| Proxy password<br>`password`| string|  | |  |  | Optional proxy password|
| Proxy port<br>`port`| integer| ✅| |  |  | Proxy port to connect to|
| Proxy Type<br>`type`| enum (string)|  | `SOCKS5`|  |  | The type of the proxy<br>Values:`SOCKS4` `SOCKS5` |
| Proxy username<br>`username`| string|  | |  |  | Optional proxy username|


#### SSL Options (Object)
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Verify Host<br>`hostnameVerifier`| boolean|  | `true`|  |  | Use to enable host name verification|
| Key store<br>`keyStore`| object|  | |  |  | <br/>See "Key store" section|
| Trust all<br>`trustAll`| boolean|  | |  |  | Use this with caution (if over Internet). The gateway must trust any origin certificates. The connection will still be encrypted but this mode is vulnerable to 'man in the middle' attacks.|
| Truststore<br>`trustStore`| object|  | |  |  | <br/>See "Truststore" section|


#### Key store (OneOf)
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Type<br>`type`| object| ✅| |  |  | <br/>See "Type" sectionType of Key store<br>Values:`""` `JKS` `JKS` `PKCS12` `PKCS12` `PEM` `PEM` |


#### Key store: None `type = ""` 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| No properties | | | | | | | 

#### Key store: JKS with path `type = "JKS"` 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Alias for the key<br>`alias`| string|  | |  |  | Alias of the key to use in case the key store contains more than one key|
| Key Password<br>`keyPassword`| string|  | |  |  | Password to use to access the key when protected by password|
| Password<br>`password`| string| ✅| |  |  | Password to use to open the key store|
| Path to key store<br>`path`| string| ✅| |  |  | Path to the key store file|


#### Key store: JKS with content `type = "JKS"` 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Alias for the key<br>`alias`| string|  | |  |  | Alias of the key to use in case the key store contains more than one key|
| Content<br>`content`| string| ✅| |  |  | Binary content as Base64|
| Key Password<br>`keyPassword`| string|  | |  |  | Password to use to access the key when protected by password|
| Password<br>`password`| string| ✅| |  |  | Password to use to open the key store|


#### Key store: PKCS#12 / PFX with path `type = "PKCS12"` 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Alias for the key<br>`alias`| string|  | |  |  | Alias of the key to use in case the key store contains more than one key|
| Key Password<br>`keyPassword`| string|  | |  |  | Password to use to access the key when protected by password|
| Password<br>`password`| string| ✅| |  |  | Password to use to open the key store|
| Path to key store<br>`path`| string| ✅| |  |  | Path to the key store file|


#### Key store: PKCS#12 / PFX with content `type = "PKCS12"` 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Alias for the key<br>`alias`| string|  | |  |  | Alias of the key to use in case the key store contains more than one key|
| Content<br>`content`| string| ✅| |  |  | Binary content as Base64|
| Key Password<br>`keyPassword`| string|  | |  |  | Password to use to access the key when protected by password|
| Password<br>`password`| string| ✅| |  |  | Password to use to open the key store|


#### Key store: PEM with path `type = "PEM"` 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Path to cert file<br>`certPath`| string| ✅| |  |  | Path to cert file (.PEM)|
| Path to private key file<br>`keyPath`| string| ✅| |  |  | Path to private key file (.PEM)|


#### Key store: PEM with content `type = "PEM"` 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Certificate<br>`certContent`| string| ✅| |  |  | |
| Private key<br>`keyContent`| string| ✅| |  |  | |


#### Truststore (OneOf)
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Type<br>`type`| object| ✅| |  |  | <br/>See "Type" sectionType of Truststore<br>Values:`""` `JKS` `JKS` `PKCS12` `PKCS12` `PEM` `PEM` |


#### Truststore: None `type = ""` 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| No properties | | | | | | | 

#### Truststore: JKS with path `type = "JKS"` 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Password<br>`password`| string| ✅| |  |  | Truststore password|
| Path to truststore<br>`path`| string| ✅| |  |  | Path to the truststore file|


#### Truststore: JKS with content `type = "JKS"` 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Content<br>`content`| string| ✅| |  |  | Binary content as Base64|
| Password<br>`password`| string| ✅| |  |  | Truststore password|


#### Truststore: PKCS#12 / PFX with path `type = "PKCS12"` 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Password<br>`password`| string| ✅| |  |  | Truststore password|
| Path to truststore<br>`path`| string| ✅| |  |  | Path to the truststore file|


#### Truststore: PKCS#12 / PFX with content `type = "PKCS12"` 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Content<br>`content`| string| ✅| |  |  | Binary content as Base64|
| Password<br>`password`| string| ✅| |  |  | Truststore password|


#### Truststore: PEM with path `type = "PEM"` 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Password<br>`password`| string| ✅| |  |  | Truststore password|
| Path to truststore<br>`path`| string| ✅| |  |  | Path to the truststore file|


#### Truststore: PEM with content `type = "PEM"` 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Content<br>`content`| string| ✅| |  |  | Binary content as Base64|
| Password<br>`password`| string| ✅| |  |  | Truststore password|


#### Context variables (Array)
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Name<br>`name`| string|  | |  |  | |
| Value<br>`value`| string|  | `{#jsonPath(#calloutResponse.content, '$.field')}`|  |  | |




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
                "scope": "REQUEST",
                "ssl": {
                  "hostnameVerifier": true,
                  "trustAll": false
                },
                "tags": [
                  "defaulted"
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
              scope: REQUEST
              ssl:
                hostnameVerifier: true
                trustAll: false
              tags:
                - defaulted
              url: http://localhost:8080/api
              variables:
                - name: field
                  value: '{#jsonPath(#calloutResponse.content, ''$.field'')}'

```

*V4 API with headers and system proxy*
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
                "url":"http://localhost:8080/api",
                "errorStatusCode": "500",
                "exitOnError": false,
                "method": "GET",
                "proxy": {
                  "enabled": true,
                  "useSystemProxy": true
                },
                "ssl": {
                  "hostnameVerifier": true,
                  "trustAll": false
                },
                "tags": [
                  "some",
                  "many"
                ]
              }
          }
        ]
      }
    ]
  }
}

```
*V4 API with headers and custom proxy*
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
                "url":"http://localhost:8080/api",
                "errorCondition": "{#calloutResponse.status >= 400 and #calloutResponse.status <= 599}",
                "errorStatusCode": "500",
                "exitOnError": false,
                "fireAndForget": false,
                "method": "GET",
                "proxy": {
                  "enabled": true,
                  "host": "proxy.acme.com",
                  "password": "[redacted]",
                  "port": 3524,
                  "type": "SOCKS5",
                  "useSystemProxy": false,
                  "username": "admin"
                },
                "scope": "REQUEST",
                "ssl": {
                  "hostnameVerifier": true,
                  "trustAll": false
                },
                "tags": [
                  "c",
                  "d"
                ],
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
*V4 API CRD with ssl pem content*
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
              url: http://localhost:8080/api
              errorStatusCode: '500'
              exitOnError: false
              method: GET
              ssl:
                hostnameVerifier: true
                trustAll: false
                trustStore:
                  content: |-
                    --- BEGIN CERTIFICATE ---
              
                    --- END CERTIFICATE ---
                  password: "[redacted]"
                  type: PEM
              tags:
                - defaulted
              variables:
                - value: "{#jsonPath(#calloutResponse.content, '$.field')}"

```


As Yaml with comments
```yaml
# Request body (string)
body: 
# Constraint test bed
# For the sake of testing
constraints: 
  # Description (string)
  description: 
  # Lower bounds (number)
  lowerBounds: 
  # Middle bounds (number)
  middleBounds: 
  # Open lower bound (number)
  openLowerBounds: 
  # Open upper bound (number)
  openUpperBounds: 
  # Ratio (number)
  ratio: 
  # Upper bounds (number)
  upperBounds: 
# Error condition (string)
# The condition which will be verified to end the request (support EL).
errorCondition: "{#calloutResponse.status >= 400 and #calloutResponse.status <= 599}"

# Error response body (string)
# The body response of the error if the condition is true (support EL)
errorContent: 
# Error status code (enum (string))
# HTTP Status Code send to the consumer if the condition is true
errorStatusCode: 500 # Possible values: "401" "411" "414" "424" "101" "200" "205" "307" "404" "412" "413" "417" "100" "207" "301" "302" "305" "406" "415" "500" "206" "300" "202" "203" "303" "405" "429" "507" "504" "102" "201" "402" "407" "408" "416" "422" "505" "304" "400" "403" "423" "204" "409" "410" "501" "502" "503" 
# Exit on error (boolean)
# Terminate the request if the error condition is true
exitOnError: 
# Fire & forget (boolean)
# Make the http call without expecting any response. When activating this mode, context variables and exit on error are useless.
fireAndForget: 
# Request Headers
headers: 
  # Name (string)
  - name: 
    # Value (string)
    value: 
# HTTP Method (enum (string))
# HTTP method to invoke the endpoint.
method: GET # Possible values: "PUT" "DELETE" "PATCH" "GET" "POST" "HEAD" "CONNECT" "OPTIONS" "TRACE" 
# Proxy Options
proxy: 
  # 
  enabled:  # Possible values: false true 
  # 
  useSystemProxy:  # Possible values: false true 
  # Proxy port
  # When enabled = true and useSystemProxy = false
  port: 3524
  # Proxy Type
  # When enabled = true and useSystemProxy = false
  type: SOCKS5 # Possible values: "SOCKS4" "SOCKS5" 
  # Proxy username
  # When enabled = true and useSystemProxy = false
  username: admin
  # Proxy host
  # When enabled = true and useSystemProxy = false
  host: proxy.acme.com
  # Proxy password
  # When enabled = true and useSystemProxy = false
  password: [redacted]
# Scope (enum (string))
# Execute policy on <strong>request</strong> (HEAD) phase, <strong>response</strong> (HEAD) phase, <strong>request_content</strong> (includes payload) phase, <strong>response content</strong> (includes payload) phase.
scope: REQUEST # Possible values: "REQUEST" "RESPONSE" "REQUEST_CONTENT" "RESPONSE_CONTENT" 
# SSL Options
ssl: 
  # Verify Host (boolean)
  # Use to enable host name verification
  hostnameVerifier: true
  # Key store
  keyStore: 
    # 
    type:  # Possible values: "JKS" "PKCS12" "PEM" "" 
    # Alias for the key
    # When type = 'JKS' or 'PKCS12'
    alias: 
    # Path to key store
    # When type = 'JKS' or 'PKCS12'
    path: 
    # Content
    # When type = 'JKS' or 'PKCS12'
    content: 
    # Path to cert file
    # When type = 'PEM'
    certPath: 
    # Certificate
    # When type = 'PEM'
    certContent: 
    # Private key
    # When type = 'PEM'
    keyContent: 
    # Key Password
    # When type = 'PKCS12' or 'JKS'
    keyPassword: 
    # Password
    # When type = 'JKS' or 'PKCS12'
    password: 
    # Path to private key file
    # When type = 'PEM'
    keyPath: 
  # Trust all (boolean)
  # Use this with caution (if over Internet). The gateway must trust any origin certificates. The connection will still be encrypted but this mode is vulnerable to 'man in the middle' attacks.
  trustAll: 
  # Truststore
  trustStore: 
    # 
    type:  # Possible values: "" "JKS" "PKCS12" "PEM" 
    # Path to truststore
    # When type = 'PEM' or 'JKS' or 'PKCS12'
    path: 
    # Content
    # When type = 'PEM' or 'JKS' or 'PKCS12'
    content: --- BEGIN CERTIFICATE ---

--- END CERTIFICATE ---
    # Password
    # When type = 'JKS' or 'PKCS12' or 'PEM'
    password: [redacted]
# Tags
# Some tags
tags: 
  # 
  - defaulted
# URL (string)
url: http://localhost:8080/api

# Context variables
variables: 
  # Name (string)
  - name: field

    # Value (string)
    value: "{#jsonPath(#calloutResponse.content, '$.field')}"


```
