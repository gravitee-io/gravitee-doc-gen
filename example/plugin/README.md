Hello, I should stay here!!
<!-- generated-start -->
# Test policy

## Overview
This my almost empty overview

with some *markup* in it : 

### Hey!
Hello!

## Foo

Bar!



    ## Errors
    You can use the response template feature to override the default response provided by the policy.
These templates are be defined at the API level, in "Entrypoint" section for V4 Apis, or in "Response Templates" for V2 APIs.

The error keys sent by this policy are as follows:

| Key| Parameters |
| --- | ---  |
| API_KEY_MISSING| - |
| API_KEY_INVALID_KEY| - |


<!-- extended-section-start -->

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
| HTTP Method<br>`method`| enum (string)| ✅| `GET`|  |  | HTTP method to invoke the endpoint.<br>Values:`GET` `POST` `PUT` `DELETE` `PATCH` `HEAD` `CONNECT` `OPTIONS` `TRACE` |
| URL<br>`url`| string| ✅| | ✅| ✅| |
| Fire & forget<br>`fireAndForget`| boolean|  | |  |  | Make the http call without expecting any response. When activating this mode, context variables and exit on error are useless.|
| Proxy Options<br>`proxy`| object|  | |  |  | <br/>See "Proxy Options" section|
| Context variables<br>`variables`| array|  | |  |  | <br/>See "Context variables" section|
| Scope<br>`scope`| enum (string)|  | `REQUEST`|  |  | Execute policy on <strong>request</strong> (HEAD) phase, <strong>response</strong> (HEAD) phase, <strong>request_content</strong> (includes payload) phase, <strong>response content</strong> (includes payload) phase.<br>Values:`REQUEST` `RESPONSE` `REQUEST_CONTENT` `RESPONSE_CONTENT` |
| Request Headers<br>`headers`| array|  | |  |  | <br/>See "Request Headers" section|
| SSL Options<br>`ssl`| object|  | |  |  | <br/>See "SSL Options" section|
| Error condition<br>`errorCondition`| string|  | `{#calloutResponse.status >= 400 and #calloutResponse.status <= 599}`|  |  | The condition which will be verified to end the request (support EL).|
| Error status code<br>`errorStatusCode`| enum (string)|  | `500`|  |  | HTTP Status Code send to the consumer if the condition is true<br>Values:`100` `101` `102` `200` `201` `202` `203` `204` `205` `206` `207` `300` `301` `302` `303` `304` `305` `307` `400` `401` `402` `403` `404` `405` `406` `407` `408` `409` `410` `411` `412` `413` `414` `415` `416` `417` `422` `423` `424` `429` `500` `501` `502` `503` `504` `505` `507` |
| Exit on error<br>`exitOnError`| boolean| ✅| |  |  | Terminate the request if the error condition is true|
| Use system proxy<br>`useSystemProxy`| boolean|  | |  |  | Use the system proxy configured by your administrator.|
| Request body<br>`body`| string|  | |  |  | |
| Error response body<br>`errorContent`| string|  | |  |  | The body response of the error if the condition is true (support EL)|


#### Proxy Options (OneOf)
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Enabled<br>`enabled`| object| ✅| |  |  | <br/>See "Enabled" sectionEnabled of Proxy Options<br>Values:`""` `true` `true` |
| Use System Proxy<br>`useSystemProxy`| object| ✅| |  |  | <br/>See "Use System Proxy" sectionUse System Proxy of Proxy Options<br>Values:`""` `true` `""` |


#### Proxy Options: No proxy `enabled = false` `useSystemProxy = false` 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|


#### Proxy Options: Use proxy configured at system level `enabled = true` `useSystemProxy = true` 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|


#### Proxy Options: Use proxy for client connections `enabled = true` `useSystemProxy = false` 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Proxy port<br>`port`| integer| ✅| |  |  | Proxy port to connect to|
| Proxy username<br>`username`| string|  | |  |  | Optional proxy username|
| Proxy password<br>`password`| string|  | |  |  | Optional proxy password|
| Proxy Type<br>`type`| enum (string)|  | `SOCKS5`|  |  | The type of the proxy<br>Values:`SOCKS4` `SOCKS5` |
| Proxy host<br>`host`| string| ✅| |  |  | Proxy host to connect to|


#### Context variables (Array)
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Name<br>`name`| string|  | |  |  | |
| Value<br>`value`| string|  | `{#jsonPath(#calloutResponse.content, '$.field')}`|  |  | |


#### Request Headers (Array)
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Value<br>`value`| string|  | | ✅| ✅| |
| Name<br>`name`| string|  | |  |  | |


#### SSL Options (Object)
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Trust all<br>`trustAll`| boolean|  | |  |  | Use this with caution (if over Internet). The gateway must trust any origin certificates. The connection will still be encrypted but this mode is vulnerable to 'man in the middle' attacks.|
| Truststore<br>`trustStore`| object|  | |  |  | <br/>See "Truststore" section|
| Key store<br>`keyStore`| object|  | |  |  | <br/>See "Key store" section|
| Verify Host<br>`hostnameVerifier`| boolean|  | `true`|  |  | Use to enable host name verification|


#### Truststore (OneOf)
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Type<br>`type`| object| ✅| |  |  | <br/>See "Type" sectionType of Truststore<br>Values:`""` `JKS` `JKS` `PKCS12` `PKCS12` `PEM` `PEM` |
| No properties | | | | | | | 

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
| Password<br>`password`| string| ✅| |  |  | Truststore password|
| Content<br>`content`| string| ✅| |  |  | Binary content as Base64|


#### Truststore: PKCS#12 / PFX with path `type = "PKCS12"` 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Password<br>`password`| string| ✅| |  |  | Truststore password|
| Path to truststore<br>`path`| string| ✅| |  |  | Path to the truststore file|


#### Truststore: PKCS#12 / PFX with content `type = "PKCS12"` 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Password<br>`password`| string| ✅| |  |  | Truststore password|
| Content<br>`content`| string| ✅| |  |  | Binary content as Base64|


#### Truststore: PEM with path `type = "PEM"` 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Password<br>`password`| string| ✅| |  |  | Truststore password|
| Path to truststore<br>`path`| string| ✅| |  |  | Path to the truststore file|


#### Truststore: PEM with content `type = "PEM"` 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Password<br>`password`| string| ✅| |  |  | Truststore password|
| Content<br>`content`| string| ✅| |  |  | Binary content as Base64|


#### Key store (OneOf)
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Type<br>`type`| object| ✅| |  |  | <br/>See "Type" sectionType of Key store<br>Values:`""` `JKS` `JKS` `PKCS12` `PKCS12` `PEM` `PEM` |
| No properties | | | | | | | 

#### Key store: None `type = ""` 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| No properties | | | | | | | 

#### Key store: JKS with path `type = "JKS"` 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Password<br>`password`| string| ✅| |  |  | Password to use to open the key store|
| Alias for the key<br>`alias`| string|  | |  |  | Alias of the key to use in case the key store contains more than one key|
| Key Password<br>`keyPassword`| string|  | |  |  | Password to use to access the key when protected by password|
| Path to key store<br>`path`| string| ✅| |  |  | Path to the key store file|


#### Key store: JKS with content `type = "JKS"` 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Password<br>`password`| string| ✅| |  |  | Password to use to open the key store|
| Alias for the key<br>`alias`| string|  | |  |  | Alias of the key to use in case the key store contains more than one key|
| Key Password<br>`keyPassword`| string|  | |  |  | Password to use to access the key when protected by password|
| Content<br>`content`| string| ✅| |  |  | Binary content as Base64|


#### Key store: PKCS#12 / PFX with path `type = "PKCS12"` 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Path to key store<br>`path`| string| ✅| |  |  | Path to the key store file|
| Password<br>`password`| string| ✅| |  |  | Password to use to open the key store|
| Alias for the key<br>`alias`| string|  | |  |  | Alias of the key to use in case the key store contains more than one key|
| Key Password<br>`keyPassword`| string|  | |  |  | Password to use to access the key when protected by password|


#### Key store: PKCS#12 / PFX with content `type = "PKCS12"` 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Password<br>`password`| string| ✅| |  |  | Password to use to open the key store|
| Alias for the key<br>`alias`| string|  | |  |  | Alias of the key to use in case the key store contains more than one key|
| Key Password<br>`keyPassword`| string|  | |  |  | Password to use to access the key when protected by password|
| Content<br>`content`| string| ✅| |  |  | Binary content as Base64|


#### Key store: PEM with path `type = "PEM"` 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Path to cert file<br>`certPath`| string| ✅| |  |  | Path to cert file (.PEM)|
| Path to private key file<br>`keyPath`| string| ✅| |  |  | Path to private key file (.PEM)|


#### Key store: PEM with content `type = "PEM"` 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Private key<br>`keyContent`| string| ✅| |  |  | |
| Certificate<br>`certContent`| string| ✅| |  |  | |




## Examples
*V4 API With default*
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
                "variables": [
                  {
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
                "errorCondition": "{#calloutResponse.status >= 400 and #calloutResponse.status <= 599}",
                "errorStatusCode": "500",
                "exitOnError": false,
                "fireAndForget": false,
                "method": "GET",
                "proxy": {
                  "enabled": true,
                  "useSystemProxy": true
                },
                "scope": "REQUEST",
                "ssl": {
                  "hostnameVerifier": true,
                  "trustAll": false
                },
                "variables": [
                  {
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
*V4 API CRD with headers and custom proxy*
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
                enabled: true
                host: proxy.acme.com
                password: '[redacted]'
                port: 3524
                type: SOCKS5
                useSystemProxy: false
                username: admin
              scope: REQUEST
              ssl:
                hostnameVerifier: true
                trustAll: false
              variables:
                - value: '{#jsonPath(#calloutResponse.content, ''$.field'')}'

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
              errorCondition: '{#calloutResponse.status >= 400 and #calloutResponse.status <= 599}'
              errorStatusCode: "500"
              exitOnError: false
              fireAndForget: false
              method: GET
              scope: REQUEST
              ssl:
                hostnameVerifier: true
                trustAll: false
                trustStore:
                  content: |-
                    --- BEGIN CERTIFICATE ---
              
                    --- END CERTIFICATE ---
                  password: '[redacted]'
                  type: PEM
              variables:
                - value: '{#jsonPath(#calloutResponse.content, ''$.field'')}'

```


