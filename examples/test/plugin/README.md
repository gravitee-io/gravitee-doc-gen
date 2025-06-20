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

| Key |
| ---  |
| API_KEY_MISSING |
| API_KEY_INVALID_KEY |



## Phases
The phases checked below are supported by the `test` policy:

### Compatible with API type:

* `PROXY`
* `MESSAGE`

### Can be used in flow phases:

* Request
* Response
* Publish
* Subscribe

## Compatibility matrix
Strikethrough line are deprecated versions

| Plugin version| APIM| AM| Comment |
| --- | --- | --- | ---  |
|~~1.0~~|~~3.x~~|~~2.1~~|~~-~~ |
|2.x|4.x|4.1|Incompatible with cloud |
|3.x|4.6 and above|4.6 and above|- |



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
| Name <br>`json name`  | Type <br>`constraint`  | Mandatory  | Default  | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:-------------|
| Request body<br>`body`| string|  | | |
| Constraint test bed<br>`constraints`| object|  | | For the sake of testing<br/>See "Constraint test bed" section|
| Error condition<br>`errorCondition`| string|  | `{#calloutResponse.status >= 400 and #calloutResponse.status <= 599}`| The condition which will be verified to end the request (support EL).|
| Error response body<br>`errorContent`| string|  | | The body response of the error if the condition is true (support EL)|
| Error status code<br>`errorStatusCode`| enum (string)|  | `500`| HTTP Status Code send to the consumer if the condition is true<br>Values: `100` `101` `102` `200` `201` `202` `203` `204` `205` `206` `207` `300` `301` `302` `303` `304` `305` `307` `400` `401` `402` `403` `404` `405` `406` `407` `408` `409` `410` `411` `412` `413` `414` `415` `416` `417` `422` `423` `424` `429` `500` `501` `502` `503` `504` `505` `507`|
| Exit on error<br>`exitOnError`| boolean| ✅| | Terminate the request if the error condition is true|
| Fire & forget<br>`fireAndForget`| boolean|  | | Make the http call without expecting any response. When activating this mode, context variables and exit on error are useless.|
| Request Headers<br>`headers`| array|  | | <br/>See "Request Headers" section|
| HTTP Method<br>`method`| enum (string)| ✅| `GET`| HTTP method to invoke the endpoint.<br>Values: `GET` `POST` `PUT` `DELETE` `PATCH` `HEAD` `CONNECT` `OPTIONS` `TRACE`|
| Proxy Options<br>`proxy`| object|  | | <br/>See "Proxy Options" section|
| SSL Options<br>`ssl`| object|  | | <br/>See "SSL Options" section|
| Tags<br>`tags`| array (string)<br>`[1, 3], unique`|  | `[defaulted and again]`| Some tags|
| URL<br>`url`| string<br>`[1, 150]`| ✅| | |
| Context variables<br>`variables`| array|  | | <br/>See "Context variables" section|


#### Constraint test bed (Object)
| Name <br>`json name`  | Type <br>`constraint`  | Mandatory  | Default  | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:-------------|
| Always the same value<br>`always`| string<br>`read-only`|  | `Static`| |
| Description<br>`description`| string<br>`[0, 1000]`|  | | |
| Lower bounds<br>`lowerBounds`| number<br>`(1, 5]`|  | | |
| Middle bounds<br>`middleBounds`| number<br>`[5, 10.333333333333]`|  | | |
| Open lower bound<br>`openLowerBounds`| number<br>`[-Inf, 10]`|  | | |
| Open upper bound<br>`openUpperBounds`| number<br>`[10, +Inf]`|  | | |
| Never returned<br>`password`| string<br>`write-only`|  | | |
| Ratio<br>`ratio`| number<br>`[0, 1]`|  | | |
| Upper bounds<br>`upperBounds`| number<br>`[10.25, 15)`|  | | |


#### Request Headers (Array)
| Name <br>`json name`  | Type <br>`constraint`  | Mandatory  | Description  |
|:----------------------|:-----------------------|:----------:|:-------------|
| Name<br>`name`| string|  | |
| Value<br>`value`| string|  | |


#### Proxy Options (Object)
| Name <br>`json name`  | Type <br>`constraint`  | Mandatory  | Description  |
|:----------------------|:-----------------------|:----------:|:-------------|
| Enabled<br>`enabled`| object| ✅| Enabled of Proxy Options<br>Values: `false` `true`|
| Use System Proxy<br>`useSystemProxy`| object| ✅| Use System Proxy of Proxy Options<br>Values: `false` `true`|


#### Proxy Options: No proxy `enabled = false` `useSystemProxy = false` 
| Name <br>`json name`  | Type <br>`constraint`  | Mandatory  | Default  | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:-------------|
| No properties | | | | | | | 

#### Proxy Options: Use proxy configured at system level `enabled = true` `useSystemProxy = true` 
| Name <br>`json name`  | Type <br>`constraint`  | Mandatory  | Default  | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:-------------|
| No properties | | | | | | | 

#### Proxy Options: Use proxy for client connections `enabled = true` `useSystemProxy = false` 
| Name <br>`json name`  | Type <br>`constraint`  | Mandatory  | Default  | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:-------------|
| Proxy host<br>`host`| string| ✅| | Proxy host to connect to|
| Proxy password<br>`password`| string|  | | Optional proxy password|
| Proxy port<br>`port`| integer| ✅| | Proxy port to connect to|
| Proxy Type<br>`type`| enum (string)|  | `SOCKS5`| The type of the proxy<br>Values: `SOCKS4` `SOCKS5`|
| Proxy username<br>`username`| string|  | | Optional proxy username|


#### SSL Options (Object)
| Name <br>`json name`  | Type <br>`constraint`  | Mandatory  | Default  | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:-------------|
| Verify Host<br>`hostnameVerifier`| boolean|  | `true`| Use to enable host name verification|
| Key store<br>`keyStore`| object|  | | <br/>See "Key store" section|
| Trust all<br>`trustAll`| boolean|  | | Use this with caution (if over Internet). The gateway must trust any origin certificates. The connection will still be encrypted but this mode is vulnerable to 'man in the middle' attacks.|
| Truststore<br>`trustStore`| object|  | | <br/>See "Truststore" section|


#### Key store (Object)
| Name <br>`json name`  | Type <br>`constraint`  | Mandatory  | Description  |
|:----------------------|:-----------------------|:----------:|:-------------|
| Type<br>`type`| object| ✅| Type of Key store<br>Values: `""` `JKS` `PKCS12` `PEM`|


#### Key store: None `type = ""` 
| Name <br>`json name`  | Type <br>`constraint`  | Mandatory  | Default  | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:-------------|
| No properties | | | | | | | 

#### Key store: JKS with path `type = "JKS"` 
| Name <br>`json name`  | Type <br>`constraint`  | Mandatory  | Default  | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:-------------|
| Alias for the key<br>`alias`| string|  | | Alias of the key to use in case the key store contains more than one key|
| Key Password<br>`keyPassword`| string|  | | Password to use to access the key when protected by password|
| Password<br>`password`| string| ✅| | Password to use to open the key store|
| Path to key store<br>`path`| string| ✅| | Path to the key store file|


#### Key store: JKS with content `type = "JKS"` 
| Name <br>`json name`  | Type <br>`constraint`  | Mandatory  | Default  | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:-------------|
| Alias for the key<br>`alias`| string|  | | Alias of the key to use in case the key store contains more than one key|
| Content<br>`content`| string| ✅| | Binary content as Base64|
| Key Password<br>`keyPassword`| string|  | | Password to use to access the key when protected by password|
| Password<br>`password`| string| ✅| | Password to use to open the key store|


#### Key store: PKCS#12 / PFX with path `type = "PKCS12"` 
| Name <br>`json name`  | Type <br>`constraint`  | Mandatory  | Default  | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:-------------|
| Alias for the key<br>`alias`| string|  | | Alias of the key to use in case the key store contains more than one key|
| Key Password<br>`keyPassword`| string|  | | Password to use to access the key when protected by password|
| Password<br>`password`| string| ✅| | Password to use to open the key store|
| Path to key store<br>`path`| string| ✅| | Path to the key store file|


#### Key store: PKCS#12 / PFX with content `type = "PKCS12"` 
| Name <br>`json name`  | Type <br>`constraint`  | Mandatory  | Default  | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:-------------|
| Alias for the key<br>`alias`| string|  | | Alias of the key to use in case the key store contains more than one key|
| Content<br>`content`| string| ✅| | Binary content as Base64|
| Key Password<br>`keyPassword`| string|  | | Password to use to access the key when protected by password|
| Password<br>`password`| string| ✅| | Password to use to open the key store|


#### Key store: PEM with path `type = "PEM"` 
| Name <br>`json name`  | Type <br>`constraint`  | Mandatory  | Default  | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:-------------|
| Path to cert file<br>`certPath`| string| ✅| | Path to cert file (.PEM)|
| Path to private key file<br>`keyPath`| string| ✅| | Path to private key file (.PEM)|


#### Key store: PEM with content `type = "PEM"` 
| Name <br>`json name`  | Type <br>`constraint`  | Mandatory  | Default  | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:-------------|
| Certificate<br>`certContent`| string| ✅| | |
| Private key<br>`keyContent`| string| ✅| | |


#### Truststore (Object)
| Name <br>`json name`  | Type <br>`constraint`  | Mandatory  | Description  |
|:----------------------|:-----------------------|:----------:|:-------------|
| Type<br>`type`| object| ✅| Type of Truststore<br>Values: `""` `JKS` `PKCS12` `PEM`|


#### Truststore: None `type = ""` 
| Name <br>`json name`  | Type <br>`constraint`  | Mandatory  | Default  | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:-------------|
| No properties | | | | | | | 

#### Truststore: JKS with path `type = "JKS"` 
| Name <br>`json name`  | Type <br>`constraint`  | Mandatory  | Default  | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:-------------|
| Password<br>`password`| string| ✅| | Truststore password|
| Path to truststore<br>`path`| string| ✅| | Path to the truststore file|


#### Truststore: JKS with content `type = "JKS"` 
| Name <br>`json name`  | Type <br>`constraint`  | Mandatory  | Default  | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:-------------|
| Content<br>`content`| string| ✅| | Binary content as Base64|
| Password<br>`password`| string| ✅| | Truststore password|


#### Truststore: PKCS#12 / PFX with path `type = "PKCS12"` 
| Name <br>`json name`  | Type <br>`constraint`  | Mandatory  | Default  | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:-------------|
| Password<br>`password`| string| ✅| | Truststore password|
| Path to truststore<br>`path`| string| ✅| | Path to the truststore file|


#### Truststore: PKCS#12 / PFX with content `type = "PKCS12"` 
| Name <br>`json name`  | Type <br>`constraint`  | Mandatory  | Default  | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:-------------|
| Content<br>`content`| string| ✅| | Binary content as Base64|
| Password<br>`password`| string| ✅| | Truststore password|


#### Truststore: PEM with path `type = "PEM"` 
| Name <br>`json name`  | Type <br>`constraint`  | Mandatory  | Default  | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:-------------|
| Password<br>`password`| string| ✅| | Truststore password|
| Path to truststore<br>`path`| string| ✅| | Path to the truststore file|


#### Truststore: PEM with content `type = "PEM"` 
| Name <br>`json name`  | Type <br>`constraint`  | Mandatory  | Default  | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:-------------|
| Content<br>`content`| string| ✅| | Binary content as Base64|
| Password<br>`password`| string| ✅| | Truststore password|


#### Context variables (Array)
| Name <br>`json name`  | Type <br>`constraint`  | Mandatory  | Default  | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:-------------|
| Name<br>`name`| string|  | | |
| Value<br>`value`| string|  | `{#jsonPath(#calloutResponse.content, '$.field')}`| |




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
                "constraints": {
                  "always": "Static"
                },
                "errorCondition": "{#calloutResponse.status >= 400 and #calloutResponse.status <= 599}",
                "errorStatusCode": "500",
                "exitOnError": false,
                "fireAndForget": false,
                "headers": [
                  {
                    "name": "Authorization",
                    "value": "Basic Jfdueh2868d="
                  },
                  {
                    "name": "X-Custom",
                    "value": "Foo"
                  }
                ],
                "method": "GET",
                "proxy": {
                  "enabled": false,
                  "useSystemProxy": false
                },
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
              constraints:
                always: Static
              errorCondition: '{#calloutResponse.status >= 400 and #calloutResponse.status <= 599}'
              errorStatusCode: "500"
              exitOnError: false
              fireAndForget: false
              headers:
                - name: Authorization
                  value: Basic Jfdueh2868d=
                - name: X-Custom
                  value: Foo
              method: GET
              proxy:
                enabled: false
                useSystemProxy: false
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
<no value>
```

## Changelog

## [4.0.0](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/3.1.0...4.0.0) (2025-01-09)


#### chore

* ack for BC ([bac00f7](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/bac00f7b71b71e51a958c7b1bdf3da1607647cd7))


#### BREAKING CHANGES

* use of secret-api 1.0.0

## [3.1.0](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/3.0.0...3.1.0) (2025-01-09)


#### Bug Fixes

* reinstate commons pool as a dependency ([b22823c](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/b22823c0df3d992d6f2f667548b67309d9eb783c))


#### Features

* add EL via annotation processor support and secrets ([a381747](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/a3817476837788e6124599838539ee56b0b9e6c0))
* rework pom management ([0741c9a](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/0741c9a90a721f0f5bf55052691c5915833c73b8))

## [3.1.0-alpha.3](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/3.1.0-alpha.2...3.1.0-alpha.3) (2025-01-07)


#### Bug Fixes

* reinstate commons pool as a dependency ([b22823c](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/b22823c0df3d992d6f2f667548b67309d9eb783c))

## [3.1.0-alpha.2](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/3.1.0-alpha.1...3.1.0-alpha.2) (2025-01-07)


#### Features

* rework pom management ([0741c9a](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/0741c9a90a721f0f5bf55052691c5915833c73b8))

## [3.1.0-alpha.1](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/3.0.0...3.1.0-alpha.1) (2025-01-07)


#### Features

* add EL via annotation processor support and secrets ([a381747](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/a3817476837788e6124599838539ee56b0b9e6c0))

## [3.0.0](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/2.0.0...3.0.0) (2024-12-30)


#### Bug Fixes

* **deps:** bump gravitee-gateway-api ([5dca38b](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/5dca38bb89a51dee1a69603b7dad8f7be3d82831))


#### Features

* update cache provider api ([1b5cdce](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/1b5cdce98c9cc37323f3853c76c5d4862ed7e787))


#### BREAKING CHANGES

* requires gravitee-gateway-api 3.9.0+ & resource-cache-provider-api 2.0.0+

## [3.0.0-alpha.1](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/2.0.0...3.0.0-alpha.1) (2024-12-30)


#### Bug Fixes

* **deps:** bump gravitee-gateway-api ([5dca38b](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/5dca38bb89a51dee1a69603b7dad8f7be3d82831))


#### Features

* update cache provider api ([1b5cdce](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/1b5cdce98c9cc37323f3853c76c5d4862ed7e787))


#### BREAKING CHANGES

* requires gravitee-gateway-api 3.9.0+ & resource-cache-provider-api 2.0.0+

## [3.0.0-alpha.1](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/2.0.0...3.0.0-alpha.1) (2024-11-12)


#### Features

* update cache provider api ([8022fc3](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/8022fc32fb1266dfa5f2a1c5647feeb0e28e9d99))


#### BREAKING CHANGES

* requires gravitee-gateway-api 3.9.0+ & resource-cache-provider-api 2.0.0+

## [2.0.0](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/1.5.0...2.0.0) (2024-09-27)


#### Features

* rework schema-form to use new GioJsonSchema Ui component ([e047513](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/e047513248b76f22c7cc9f113c8ef698d4a29c7f))


#### BREAKING CHANGES

* rework schema-form

## [1.5.0](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/1.4.0...1.5.0) (2024-09-27)


#### Features

* **release:** compatibility issue 1.4.0 introduced a breaking change ([0200cb4](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/0200cb4371d89ca4b994a49e29580cb229ae2a9e))

## [1.4.0](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/1.3.0...1.4.0) (2024-07-12)


#### Features

* rework schema-form to use new GioJsonSchema Ui component ([5f08b0c](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/5f08b0c6daafe89304863cac3ecd40110a0b1edf))

## [1.3.0](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/1.2.0...1.3.0) (2023-03-17)


#### Bug Fixes

* **deps:** bump dependencies ([7a18ca5](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/7a18ca58294bb95577986a1319422e8e1dc694a5))


#### Features

* rename 'jupiter' package in 'reactive' ([39e045c](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/39e045c85d43af4b3f10305d5dd24752f3da9e05))

## [1.3.0-alpha.1](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/1.2.0...1.3.0-alpha.1) (2023-03-13)


#### Features

* rename 'jupiter' package in 'reactive' ([4fb6401](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/4fb6401959a84a025f2e0d5423a19ce2102060dd))

## [1.2.0](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/1.1.0...1.2.0) (2022-09-02)


#### Features

* improve execution context structure ([1cd894f](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/1cd894ff48ae41adf8569323c77fd981089097e7)), closes [gravitee-io/issues#8386](https://github.com/gravitee-io/issues/issues/8386)

## [1.1.0](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/1.0.1...1.1.0) (2022-06-10)


#### Features

* **jupiter:** implement getCache with jupiter ExecutionContext ([ea96ff2](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/ea96ff232f208ccd40202289f94f17fcca07e27b))

### [1.0.1](https://github.com/gravitee-io/gravitee-resource-cache-redis/compare/1.0.0...1.0.1) (2022-02-22)


#### Bug Fixes

* resolve form configuration ([985be4f](https://github.com/gravitee-io/gravitee-resource-cache-redis/commit/985be4f7ce6e6bd026cf375905cd8e10da346c28)), closes [gravitee-io/issues#7172](https://github.com/gravitee-io/issues/issues/7172)


## Environment variable

<no value>