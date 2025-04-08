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
| Request body<br>`body`| string|  | |  |  | |
| Description<br>`description`| string<br>[0, 1000]|  | |  |  | |
| Error condition<br>`errorCondition`| string|  | `{#calloutResponse.status >= 400 and #calloutResponse.status <= 599}`|  |  | The condition which will be verified to end the request (support EL).|
| Error response body<br>`errorContent`| string|  | |  |  | The body response of the error if the condition is true (support EL)|
| Error status code<br>`errorStatusCode`| enum (string)|  | `500`|  |  | HTTP Status Code send to the consumer if the condition is true<br>Values:`100` `101` `102` `200` `201` `202` `203` `204` `205` `206` `207` `300` `301` `302` `303` `304` `305` `307` `400` `401` `402` `403` `404` `405` `406` `407` `408` `409` `410` `411` `412` `413` `414` `415` `416` `417` `422` `423` `424` `429` `500` `501` `502` `503` `504` `505` `507` |
| Exit on error<br>`exitOnError`| boolean| ✅| |  |  | Terminate the request if the error condition is true|
| Fire & forget<br>`fireAndForget`| boolean|  | |  |  | Make the http call without expecting any response. When activating this mode, context variables and exit on error are useless.|
| Request Headers<br>`headers`| array|  | |  |  | <br/>See "Request Headers" section|
| Lower bounds<br>`lowerBounds`| number<br>(1, 5]|  | |  |  | |
| HTTP Method<br>`method`| enum (string)| ✅| `GET`|  |  | HTTP method to invoke the endpoint.<br>Values:`GET` `POST` `PUT` `DELETE` `PATCH` `HEAD` `CONNECT` `OPTIONS` `TRACE` |
| Middle bounds<br>`middleBounds`| number<br>[5, 10.333333333333]|  | |  |  | |
| Open lower bound<br>`openLowerBounds`| number<br>[-Inf, 10]|  | |  |  | |
| Open upper bound<br>`openUpperBounds`| number<br>[10, +Inf]|  | |  |  | |
| Proxy Options<br>`proxy`| object|  | |  |  | <br/>See "Proxy Options" section|
| Ratio<br>`ratio`| number<br>[0, 1]|  | |  |  | |
| Scope<br>`scope`| enum (string)|  | `REQUEST`|  |  | Execute policy on <strong>request</strong> (HEAD) phase, <strong>response</strong> (HEAD) phase, <strong>request_content</strong> (includes payload) phase, <strong>response content</strong> (includes payload) phase.<br>Values:`REQUEST` `RESPONSE` `REQUEST_CONTENT` `RESPONSE_CONTENT` |
| SSL Options<br>`ssl`| object|  | |  |  | <br/>See "SSL Options" section|
| Tags<br>`tags`| array<br>[1, 3], unique|  | `[defaulted]`|  |  | Some tags|
| Upper bounds<br>`upperBounds`| number<br>[10.25, 15)|  | |  |  | |
| URL<br>`url`| string<br>[1, 150]| ✅| | ✅| ✅| |
| Context variables<br>`variables`| array|  | |  |  | <br/>See "Context variables" section|


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


#### Proxy Options: Use proxy configured at system level `enabled = true` `useSystemProxy = true` 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|


#### Proxy Options: Use proxy for client connections `enabled = true` `useSystemProxy = false` 
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Proxy host<br>`host`| string| ✅| |  |  | Proxy host to connect to|
| Proxy password<br>`password`| string|  | |  |  | Optional proxy password|
| Proxy port<br>`port`| integer| ✅| |  |  | Proxy port to connect to|
| Proxy Type<br>`type`| enum (string)|  | `SOCKS5`|  |  | The type of the proxy<br>Values:`SOCKS4` `SOCKS5` |
| Proxy username<br>`username`| string|  | |  |  | Optional proxy username|


#### Context variables (Array)
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Name<br>`name`| string|  | |  |  | |
| Value<br>`value`| string|  | `{#jsonPath(#calloutResponse.content, '$.field')}`|  |  | |




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
                "tags": [
                  "defaulted"
                ],
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
                "tags": [
                  "some",
                  "many"
                ],
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
              tags:
                - c
                - d
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
              tags:
                - defaulted
              variables:
                - value: '{#jsonPath(#calloutResponse.content, ''$.field'')}'

```


