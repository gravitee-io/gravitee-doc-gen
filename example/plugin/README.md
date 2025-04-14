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


## As Yaml with comments
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
errorStatusCode: 500 # Possible values: "100" "101" "102" "200" "201" "202" "203" "204" "205" "206" "207" "300" "301" "302" "303" "304" "305" "307" "400" "401" "402" "403" "404" "405" "406" "407" "408" "409" "410" "411" "412" "413" "414" "415" "416" "417" "422" "423" "424" "429" "500" "501" "502" "503" "504" "505" "507" 
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
method: GET # Possible values: "GET" "POST" "PUT" "DELETE" "PATCH" "HEAD" "CONNECT" "OPTIONS" "TRACE" 
# Proxy Options
proxy: 
  # 
  enabled:  # Possible values: false true 
  # 
  useSystemProxy:  # Possible values: false true 
  # Proxy username
  # When enabled = true and useSystemProxy = false
  username: admin
  # Proxy host
  # When enabled = true and useSystemProxy = false
  host: proxy.acme.com
  # Proxy password
  # When enabled = true and useSystemProxy = false
  password: "[redacted]"
  # Proxy port
  # When enabled = true and useSystemProxy = false
  port: 3524
  # Proxy Type
  # When enabled = true and useSystemProxy = false
  type: SOCKS5 # Possible values: "SOCKS5" "SOCKS4" 
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
    type:  # Possible values: "" "JKS" "PKCS12" "PEM" 
    # Certificate
    # When type = 'PEM'
    certContent: 
    # Alias for the key
    # When type = 'PKCS12' or 'JKS'
    alias: 
    # Password
    # When type = 'JKS' or 'PKCS12'
    password: 
    # Content
    # When type = 'JKS' or 'PKCS12'
    content: 
    # Path to cert file
    # When type = 'PEM'
    certPath: 
    # Path to private key file
    # When type = 'PEM'
    keyPath: 
    # Private key
    # When type = 'PEM'
    keyContent: 
    # Key Password
    # When type = 'JKS' or 'PKCS12'
    keyPassword: 
    # Path to key store
    # When type = 'PKCS12' or 'JKS'
    path: 
  # Trust all (boolean)
  # Use this with caution (if over Internet). The gateway must trust any origin certificates. The connection will still be encrypted but this mode is vulnerable to 'man in the middle' attacks.
  trustAll: 
  # Truststore
  trustStore: 
    # 
    type:  # Possible values: "" "JKS" "PKCS12" "PEM" 
    # Password
    # When type = 'PKCS12' or 'PEM' or 'JKS'
    password: "[redacted]"
    # Path to truststore
    # When type = 'JKS' or 'PKCS12' or 'PEM'
    path: 
    # Content
    # When type = 'PEM' or 'JKS' or 'PKCS12'
    content: |-
        --- BEGIN CERTIFICATE ---
    
        --- END CERTIFICATE ---
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

## Changelog

### [4.0.0](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/3.1.0...4.0.0) (2025-01-09)
 [4.0.0](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/3.1.0...4.0.0) (2025-01-09)


##### chore
 chore

* ack for BC ([bac00f7](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/bac00f7b71b71e51a958c7b1bdf3da1607647cd7))


##### BREAKING CHANGES
 BREAKING CHANGES

* use of secret-api 1.0.0

### [3.1.0](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/3.0.0...3.1.0) (2025-01-09)
 [3.1.0](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/3.0.0...3.1.0) (2025-01-09)


##### Bug Fixes
 Bug Fixes

* reinstate commons pool as a dependency ([b22823c](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/b22823c0df3d992d6f2f667548b67309d9eb783c))


##### Features
 Features

* add EL via annotation processor support and secrets ([a381747](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/a3817476837788e6124599838539ee56b0b9e6c0))
* rework pom management ([0741c9a](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/0741c9a90a721f0f5bf55052691c5915833c73b8))

### [3.1.0-alpha.3](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/3.1.0-alpha.2...3.1.0-alpha.3) (2025-01-07)
 [3.1.0-alpha.3](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/3.1.0-alpha.2...3.1.0-alpha.3) (2025-01-07)


##### Bug Fixes
 Bug Fixes

* reinstate commons pool as a dependency ([b22823c](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/b22823c0df3d992d6f2f667548b67309d9eb783c))

### [3.1.0-alpha.2](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/3.1.0-alpha.1...3.1.0-alpha.2) (2025-01-07)
 [3.1.0-alpha.2](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/3.1.0-alpha.1...3.1.0-alpha.2) (2025-01-07)


##### Features
 Features

* rework pom management ([0741c9a](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/0741c9a90a721f0f5bf55052691c5915833c73b8))

### [3.1.0-alpha.1](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/3.0.0...3.1.0-alpha.1) (2025-01-07)
 [3.1.0-alpha.1](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/3.0.0...3.1.0-alpha.1) (2025-01-07)


##### Features
 Features

* add EL via annotation processor support and secrets ([a381747](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/a3817476837788e6124599838539ee56b0b9e6c0))

### [3.0.0](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/2.0.0...3.0.0) (2024-12-30)
 [3.0.0](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/2.0.0...3.0.0) (2024-12-30)


##### Bug Fixes
 Bug Fixes

* **deps:** bump gravitee-gateway-api ([5dca38b](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/5dca38bb89a51dee1a69603b7dad8f7be3d82831))


##### Features
 Features

* update cache provider api ([1b5cdce](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/1b5cdce98c9cc37323f3853c76c5d4862ed7e787))


##### BREAKING CHANGES
 BREAKING CHANGES

* requires gravitee-gateway-api 3.9.0+ & resource-cache-provider-api 2.0.0+

### [3.0.0-alpha.1](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/2.0.0...3.0.0-alpha.1) (2024-12-30)
 [3.0.0-alpha.1](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/2.0.0...3.0.0-alpha.1) (2024-12-30)


##### Bug Fixes
 Bug Fixes

* **deps:** bump gravitee-gateway-api ([5dca38b](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/5dca38bb89a51dee1a69603b7dad8f7be3d82831))


##### Features
 Features

* update cache provider api ([1b5cdce](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/1b5cdce98c9cc37323f3853c76c5d4862ed7e787))


##### BREAKING CHANGES
 BREAKING CHANGES

* requires gravitee-gateway-api 3.9.0+ & resource-cache-provider-api 2.0.0+

### [3.0.0-alpha.1](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/2.0.0...3.0.0-alpha.1) (2024-11-12)
 [3.0.0-alpha.1](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/2.0.0...3.0.0-alpha.1) (2024-11-12)


##### Features
 Features

* update cache provider api ([8022fc3](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/8022fc32fb1266dfa5f2a1c5647feeb0e28e9d99))


##### BREAKING CHANGES
 BREAKING CHANGES

* requires gravitee-gateway-api 3.9.0+ & resource-cache-provider-api 2.0.0+

### [2.0.0](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/1.5.0...2.0.0) (2024-09-27)
 [2.0.0](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/1.5.0...2.0.0) (2024-09-27)


##### Features
 Features

* rework schema-form to use new GioJsonSchema Ui component ([e047513](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/e047513248b76f22c7cc9f113c8ef698d4a29c7f))


##### BREAKING CHANGES
 BREAKING CHANGES

* rework schema-form

### [1.5.0](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/1.4.0...1.5.0) (2024-09-27)
 [1.5.0](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/1.4.0...1.5.0) (2024-09-27)


##### Features
 Features

* **release:** compatibility issue 1.4.0 introduced a breaking change ([0200cb4](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/0200cb4371d89ca4b994a49e29580cb229ae2a9e))

### [1.4.0](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/1.3.0...1.4.0) (2024-07-12)
 [1.4.0](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/1.3.0...1.4.0) (2024-07-12)


##### Features
 Features

* rework schema-form to use new GioJsonSchema Ui component ([5f08b0c](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/5f08b0c6daafe89304863cac3ecd40110a0b1edf))

### [1.3.0](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/1.2.0...1.3.0) (2023-03-17)
 [1.3.0](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/1.2.0...1.3.0) (2023-03-17)


##### Bug Fixes
 Bug Fixes

* **deps:** bump dependencies ([7a18ca5](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/7a18ca58294bb95577986a1319422e8e1dc694a5))


##### Features
 Features

* rename 'jupiter' package in 'reactive' ([39e045c](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/39e045c85d43af4b3f10305d5dd24752f3da9e05))

### [1.3.0-alpha.1](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/1.2.0...1.3.0-alpha.1) (2023-03-13)
 [1.3.0-alpha.1](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/1.2.0...1.3.0-alpha.1) (2023-03-13)


##### Features
 Features

* rename 'jupiter' package in 'reactive' ([4fb6401](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/4fb6401959a84a025f2e0d5423a19ce2102060dd))

### [1.2.0](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/1.1.0...1.2.0) (2022-09-02)
 [1.2.0](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/1.1.0...1.2.0) (2022-09-02)


##### Features
 Features

* improve execution context structure ([1cd894f](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/1cd894ff48ae41adf8569323c77fd981089097e7)), closes [gravitee-io/issues#8386](https://github.com/gravitee-io/issues/issues/8386)

### [1.1.0](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/1.0.1...1.1.0) (2022-06-10)
 [1.1.0](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/1.0.1...1.1.0) (2022-06-10)


##### Features
 Features

* **jupiter:** implement getCache with jupiter ExecutionContext ([ea96ff2](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/ea96ff232f208ccd40202289f94f17fcca07e27b))

#### [1.0.1](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/1.0.0...1.0.1) (2022-02-22)
 [1.0.1](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/1.0.0...1.0.1) (2022-02-22)


##### Bug Fixes
 Bug Fixes

* resolve form configuration ([985be4f](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/985be4f7ce6e6bd026cf375905cd8e10da346c28)), closes [gravitee-io/issues#7172](https://github.com/gravitee-io/issues/issues/7172)


## Environment variable


### 



#### Request body
ENV: **GRAVITEE_BODY** <br>
JVM: `-Dgravitee.body`  

<p>

</p>



#### Constraint test bed
ENV: **GRAVITEE_CONSTRAINTS** <br>
JVM: `-Dgravitee.constraints`  

<p>
For the sake of testing
</p>



#### Error condition
ENV: **GRAVITEE_ERRORCONDITION** <br>
JVM: `-Dgravitee.errorcondition` <br>
Default: `{#calloutResponse.status >= 400 and #calloutResponse.status <= 599}`  

<p>
The condition which will be verified to end the request (support EL).
</p>



#### Error response body
ENV: **GRAVITEE_ERRORCONTENT** <br>
JVM: `-Dgravitee.errorcontent`  

<p>
The body response of the error if the condition is true (support EL)
</p>



#### Error status code
ENV: **GRAVITEE_ERRORSTATUSCODE** <br>
JVM: `-Dgravitee.errorstatuscode` <br>
Default: `500`  <br>
Values: `100` `101` `102` `200` `201` `202` `203` `204` `205` `206` `207` `300` `301` `302` `303` `304` `305` `307` `400` `401` `402` `403` `404` `405` `406` `407` `408` `409` `410` `411` `412` `413` `414` `415` `416` `417` `422` `423` `424` `429` `500` `501` `502` `503` `504` `505` `507` 

<p>
HTTP Status Code send to the consumer if the condition is true
</p>



#### Exit on error
ENV: **GRAVITEE_EXITONERROR** <br>
JVM: `-Dgravitee.exitonerror`  

<p>
Terminate the request if the error condition is true
</p>



#### Fire & forget
ENV: **GRAVITEE_FIREANDFORGET** <br>
JVM: `-Dgravitee.fireandforget`  

<p>
Make the http call without expecting any response. When activating this mode, context variables and exit on error are useless.
</p>



#### Request Headers
ENV: **GRAVITEE_HEADERS** <br>
JVM: `-Dgravitee.headers`  

<p>

</p>



#### HTTP Method
ENV: **GRAVITEE_METHOD** <br>
JVM: `-Dgravitee.method` <br>
Default: `GET`  <br>
Values: `GET` `POST` `PUT` `DELETE` `PATCH` `HEAD` `CONNECT` `OPTIONS` `TRACE` 

<p>
HTTP method to invoke the endpoint.
</p>



#### Proxy Options
ENV: **GRAVITEE_PROXY** <br>
JVM: `-Dgravitee.proxy`  

<p>

</p>



#### Scope
ENV: **GRAVITEE_SCOPE** <br>
JVM: `-Dgravitee.scope` <br>
Default: `REQUEST`  <br>
Values: `REQUEST` `RESPONSE` `REQUEST_CONTENT` `RESPONSE_CONTENT` 

<p>
Execute policy on <strong>request</strong> (HEAD) phase, <strong>response</strong> (HEAD) phase, <strong>request_content</strong> (includes payload) phase, <strong>response content</strong> (includes payload) phase.
</p>



#### SSL Options
ENV: **GRAVITEE_SSL** <br>
JVM: `-Dgravitee.ssl`  

<p>

</p>



#### Tags
ENV: **GRAVITEE_TAGS** <br>
JVM: `-Dgravitee.tags` <br>
Default: `[defaulted]`  

<p>
Some tags
</p>



#### URL
ENV: **GRAVITEE_URL** <br>
JVM: `-Dgravitee.url` <br>
Default: `http://localhost:8080/api`  

<p>

</p>



#### Context variables
ENV: **GRAVITEE_VARIABLES** <br>
JVM: `-Dgravitee.variables`  

<p>

</p>




### Constraint test bed
For the sake of testing


#### Description
ENV: **GRAVITEE_CONSTRAINTS_DESCRIPTION** <br>
JVM: `-Dgravitee.constraints.description`  

<p>

</p>



#### Lower bounds
ENV: **GRAVITEE_CONSTRAINTS_LOWERBOUNDS** <br>
JVM: `-Dgravitee.constraints.lowerbounds`  

<p>

</p>



#### Middle bounds
ENV: **GRAVITEE_CONSTRAINTS_MIDDLEBOUNDS** <br>
JVM: `-Dgravitee.constraints.middlebounds`  

<p>

</p>



#### Open lower bound
ENV: **GRAVITEE_CONSTRAINTS_OPENLOWERBOUNDS** <br>
JVM: `-Dgravitee.constraints.openlowerbounds`  

<p>

</p>



#### Open upper bound
ENV: **GRAVITEE_CONSTRAINTS_OPENUPPERBOUNDS** <br>
JVM: `-Dgravitee.constraints.openupperbounds`  

<p>

</p>



#### Ratio
ENV: **GRAVITEE_CONSTRAINTS_RATIO** <br>
JVM: `-Dgravitee.constraints.ratio`  

<p>

</p>



#### Upper bounds
ENV: **GRAVITEE_CONSTRAINTS_UPPERBOUNDS** <br>
JVM: `-Dgravitee.constraints.upperbounds`  

<p>

</p>




### Request Headers



#### Name
ENV: **GRAVITEE_HEADERS_{index}_HEADERS_NAME** <br>
JVM: `-Dgravitee.headers.[{index}].headers.name`  

<p>

</p>



#### Value
ENV: **GRAVITEE_HEADERS_{index}_HEADERS_VALUE** <br>
JVM: `-Dgravitee.headers.[{index}].headers.value`  

<p>

</p>




### Proxy Options



#### 
ENV: **GRAVITEE_PROXY_ENABLED** <br>
JVM: `-Dgravitee.proxy.enabled`  <br>
Values: `false` `true` 

<p>

</p>



#### 
ENV: **GRAVITEE_PROXY_USESYSTEMPROXY** <br>
JVM: `-Dgravitee.proxy.usesystemproxy`  <br>
Values: `false` `true` 

<p>

</p>



#### Proxy port
ENV: **GRAVITEE_PROXY_PORT** <br>
JVM: `-Dgravitee.proxy.port` <br>
Default: `3524`  <br>
When: <br>
`enabled = true`  and <br>
`useSystemProxy = false` 

<p>
Proxy port to connect to
</p>



#### Proxy Type
ENV: **GRAVITEE_PROXY_TYPE** <br>
JVM: `-Dgravitee.proxy.type` <br>
Default: `SOCKS5`  <br>
Values: `SOCKS4` `SOCKS5` <br>
When: <br>
`enabled = true`  and <br>
`useSystemProxy = false` 

<p>
The type of the proxy
</p>



#### Proxy username
ENV: **GRAVITEE_PROXY_USERNAME** <br>
JVM: `-Dgravitee.proxy.username` <br>
Default: `admin`  <br>
When: <br>
`enabled = true`  and <br>
`useSystemProxy = false` 

<p>
Optional proxy username
</p>



#### Proxy host
ENV: **GRAVITEE_PROXY_HOST** <br>
JVM: `-Dgravitee.proxy.host` <br>
Default: `proxy.acme.com`  <br>
When: <br>
`enabled = true`  and <br>
`useSystemProxy = false` 

<p>
Proxy host to connect to
</p>



#### Proxy password
ENV: **GRAVITEE_PROXY_PASSWORD** <br>
JVM: `-Dgravitee.proxy.password` <br>
Default: `[redacted]`  <br>
When: <br>
`enabled = true`  and <br>
`useSystemProxy = false` 

<p>
Optional proxy password
</p>




### SSL Options



#### Verify Host
ENV: **GRAVITEE_SSL_HOSTNAMEVERIFIER** <br>
JVM: `-Dgravitee.ssl.hostnameverifier` <br>
Default: `true`  

<p>
Use to enable host name verification
</p>



#### Key store
ENV: **GRAVITEE_SSL_KEYSTORE** <br>
JVM: `-Dgravitee.ssl.keystore`  

<p>

</p>



#### Trust all
ENV: **GRAVITEE_SSL_TRUSTALL** <br>
JVM: `-Dgravitee.ssl.trustall`  

<p>
Use this with caution (if over Internet). The gateway must trust any origin certificates. The connection will still be encrypted but this mode is vulnerable to 'man in the middle' attacks.
</p>



#### Truststore
ENV: **GRAVITEE_SSL_TRUSTSTORE** <br>
JVM: `-Dgravitee.ssl.truststore`  

<p>

</p>



#### 
ENV: **GRAVITEE_SSL_KEYSTORE_TYPE** <br>
JVM: `-Dgravitee.ssl.keystore.type`  <br>
Values: `` `JKS` `PKCS12` `PEM` 

<p>

</p>



#### Alias for the key
ENV: **GRAVITEE_SSL_KEYSTORE_ALIAS** <br>
JVM: `-Dgravitee.ssl.keystore.alias`  <br>
When: <br>
`type = PKCS12` or `JKS` 

<p>
Alias of the key to use in case the key store contains more than one key
</p>



#### Key Password
ENV: **GRAVITEE_SSL_KEYSTORE_KEYPASSWORD** <br>
JVM: `-Dgravitee.ssl.keystore.keypassword`  <br>
When: <br>
`type = PKCS12` or `JKS` 

<p>
Password to use to access the key when protected by password
</p>



#### Path to key store
ENV: **GRAVITEE_SSL_KEYSTORE_PATH** <br>
JVM: `-Dgravitee.ssl.keystore.path`  <br>
When: <br>
`type = PKCS12` or `JKS` 

<p>
Path to the key store file
</p>



#### Path to cert file
ENV: **GRAVITEE_SSL_KEYSTORE_CERTPATH** <br>
JVM: `-Dgravitee.ssl.keystore.certpath`  <br>
When: <br>
`type = PEM` 

<p>
Path to cert file (.PEM)
</p>



#### Password
ENV: **GRAVITEE_SSL_KEYSTORE_PASSWORD** <br>
JVM: `-Dgravitee.ssl.keystore.password`  <br>
When: <br>
`type = JKS` or `PKCS12` 

<p>
Password to use to open the key store
</p>



#### Content
ENV: **GRAVITEE_SSL_KEYSTORE_CONTENT** <br>
JVM: `-Dgravitee.ssl.keystore.content`  <br>
When: <br>
`type = JKS` or `PKCS12` 

<p>
Binary content as Base64
</p>



#### Path to private key file
ENV: **GRAVITEE_SSL_KEYSTORE_KEYPATH** <br>
JVM: `-Dgravitee.ssl.keystore.keypath`  <br>
When: <br>
`type = PEM` 

<p>
Path to private key file (.PEM)
</p>



#### Certificate
ENV: **GRAVITEE_SSL_KEYSTORE_CERTCONTENT** <br>
JVM: `-Dgravitee.ssl.keystore.certcontent`  <br>
When: <br>
`type = PEM` 

<p>

</p>



#### Private key
ENV: **GRAVITEE_SSL_KEYSTORE_KEYCONTENT** <br>
JVM: `-Dgravitee.ssl.keystore.keycontent`  <br>
When: <br>
`type = PEM` 

<p>

</p>



#### 
ENV: **GRAVITEE_SSL_TRUSTSTORE_TYPE** <br>
JVM: `-Dgravitee.ssl.truststore.type`  <br>
Values: `` `JKS` `PKCS12` `PEM` 

<p>

</p>



#### Password
ENV: **GRAVITEE_SSL_TRUSTSTORE_PASSWORD** <br>
JVM: `-Dgravitee.ssl.truststore.password` <br>
Default: `[redacted]`  <br>
When: <br>
`type = JKS` or `PKCS12` or `PEM` 

<p>
Truststore password
</p>



#### Path to truststore
ENV: **GRAVITEE_SSL_TRUSTSTORE_PATH** <br>
JVM: `-Dgravitee.ssl.truststore.path`  <br>
When: <br>
`type = PEM` or `JKS` or `PKCS12` 

<p>
Path to the truststore file
</p>



#### Content
ENV: **GRAVITEE_SSL_TRUSTSTORE_CONTENT** <br>
JVM: `-Dgravitee.ssl.truststore.content` <br>
Default: `--- BEGIN CERTIFICATE ---

--- END CERTIFICATE ---`  <br>
When: <br>
`type = JKS` or `PKCS12` or `PEM` 

<p>
Binary content as Base64
</p>



#### Tags
ENV: **GRAVITEE_TAGS** <br>
JVM: `-Dgravitee.tags` <br>
Default: `[defaulted]`  

<p>
Some tags
</p>




### Context variables



#### Name
ENV: **GRAVITEE_VARIABLES_{index}_VARIABLES_NAME** <br>
JVM: `-Dgravitee.variables.[{index}].variables.name` <br>
Default: `field`  

<p>

</p>



#### Value
ENV: **GRAVITEE_VARIABLES_{index}_VARIABLES_VALUE** <br>
JVM: `-Dgravitee.variables.[{index}].variables.value` <br>
Default: `{#jsonPath(#calloutResponse.content, '$.field')}`  

<p>

</p>



