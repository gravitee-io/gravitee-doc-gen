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
| Request Headers<br>`headers`| array|  | |  |  | <br/>See "Request Headers" section|| Fire & forget<br>`fireAndForget`| boolean|  | `false`|  |  | Make the http call without expecting any response. When activating this mode, context variables and exit on error are useless.|| Error condition<br>`errorCondition`| string|  | `{#calloutResponse.status >= 400 and #calloutResponse.status <= 599}`|  |  | The condition which will be verified to end the request (support EL).|| Error status code<br>`errorStatusCode`| enum (string)|  | `500`|  |  | HTTP Status Code send to the consumer if the condition is true<br>Values:`100` `101` `102` `200` `201` `202` `203` `204` `205` `206` `207` `300` `301` `302` `303` `304` `305` `307` `400` `401` `402` `403` `404` `405` `406` `407` `408` `409` `410` `411` `412` `413` `414` `415` `416` `417` `422` `423` `424` `429` `500` `501` `502` `503` `504` `505` `507` || Scope<br>`scope`| enum (string)|  | `REQUEST`|  |  | Execute policy on <strong>request</strong> (HEAD) phase, <strong>response</strong> (HEAD) phase, <strong>request_content</strong> (includes payload) phase, <strong>response content</strong> (includes payload) phase.<br>Values:`REQUEST` `RESPONSE` `REQUEST_CONTENT` `RESPONSE_CONTENT` || URL<br>`url`| string| ✅| | ✅| ✅| || Error response body<br>`errorContent`| string|  | |  |  | The body response of the error if the condition is true (support EL)|| HTTP Method<br>`method`| enum (string)| ✅| `GET`|  |  | HTTP method to invoke the endpoint.<br>Values:`GET` `POST` `PUT` `DELETE` `PATCH` `HEAD` `CONNECT` `OPTIONS` `TRACE` || Use system proxy<br>`useSystemProxy`| boolean|  | |  |  | Use the system proxy configured by your administrator.|| Request body<br>`body`| string|  | |  |  | || Context variables<br>`variables`| array|  | |  |  | <br/>See "Context variables" section|| Exit on error<br>`exitOnError`| boolean| ✅| `false`|  |  | Terminate the request if the error condition is true|| SSL Options<br>`ssl`| object|  | |  |  | <br/>See "SSL Options" section|| No properties | | | | | | | 

#### Context variables (Array)
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Value<br>`value`| string|  | `{#jsonPath(#calloutResponse.content, '$.field')}`|  |  | || Name<br>`name`| string|  | |  |  | || No properties | | | | | | | 

#### SSL Options (Object)
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Verify Host<br>`hostnameVerifier`| boolean|  | `true`|  |  | Use to enable host name verification|| Trust all<br>`trustAll`| boolean|  | `false`|  |  | Use this with caution (if over Internet). The gateway must trust any origin certificates. The connection will still be encrypted but this mode is vulnerable to 'man in the middle' attacks.|| Truststore<br>`trustStore`| object|  | |  |  | <br/>See "Truststore" section|| Key store<br>`keyStore`| object|  | |  |  | <br/>See "Key store" section|| No properties | | | | | | | 

#### Truststore (OneOf)
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Type<br>`type`| string| ✅| |  |  | Type of Truststore<br>Values:`""` `JKS` `JKS` `PKCS12` `PKCS12` `PEM` `PEM` || No properties | | | | | | | 

#### Truststore: None `type = ""`
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| No properties | | | | | | | 

#### Truststore: JKS with path `type = "JKS"`
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Password<br>`password`| string| ✅| |  |  | Truststore password|| Path to truststore<br>`path`| string| ✅| |  |  | Path to the truststore file|| No properties | | | | | | | 

#### Truststore: JKS with content `type = "JKS"`
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Password<br>`password`| string| ✅| |  |  | Truststore password|| Content<br>`content`| string| ✅| |  |  | Binary content as Base64|| No properties | | | | | | | 

#### Truststore: PKCS#12 / PFX with path `type = "PKCS12"`
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Password<br>`password`| string| ✅| |  |  | Truststore password|| Path to truststore<br>`path`| string| ✅| |  |  | Path to the truststore file|| No properties | | | | | | | 

#### Truststore: PKCS#12 / PFX with content `type = "PKCS12"`
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Content<br>`content`| string| ✅| |  |  | Binary content as Base64|| Password<br>`password`| string| ✅| |  |  | Truststore password|| No properties | | | | | | | 

#### Truststore: PEM with path `type = "PEM"`
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Path to truststore<br>`path`| string| ✅| |  |  | Path to the truststore file|| Password<br>`password`| string| ✅| |  |  | Truststore password|| No properties | | | | | | | 

#### Truststore: PEM with content `type = "PEM"`
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Password<br>`password`| string| ✅| |  |  | Truststore password|| Content<br>`content`| string| ✅| |  |  | Binary content as Base64|| No properties | | | | | | | 

#### Key store (OneOf)
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Type<br>`type`| string| ✅| |  |  | Type of Key store<br>Values:`""` `JKS` `JKS` `PKCS12` `PKCS12` `PEM` `PEM` || No properties | | | | | | | 

#### Key store: None `type = ""`
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| No properties | | | | | | | 

#### Key store: JKS with path `type = "JKS"`
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Path to key store<br>`path`| string| ✅| |  |  | Path to the key store file|| Password<br>`password`| string| ✅| |  |  | Password to use to open the key store|| Alias for the key<br>`alias`| string|  | |  |  | Alias of the key to use in case the key store contains more than one key|| Key Password<br>`keyPassword`| string|  | |  |  | Password to use to access the key when protected by password|| No properties | | | | | | | 

#### Key store: JKS with content `type = "JKS"`
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Key Password<br>`keyPassword`| string|  | |  |  | Password to use to access the key when protected by password|| Content<br>`content`| string| ✅| |  |  | Binary content as Base64|| Password<br>`password`| string| ✅| |  |  | Password to use to open the key store|| Alias for the key<br>`alias`| string|  | |  |  | Alias of the key to use in case the key store contains more than one key|| No properties | | | | | | | 

#### Key store: PKCS#12 / PFX with path `type = "PKCS12"`
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Password<br>`password`| string| ✅| |  |  | Password to use to open the key store|| Alias for the key<br>`alias`| string|  | |  |  | Alias of the key to use in case the key store contains more than one key|| Key Password<br>`keyPassword`| string|  | |  |  | Password to use to access the key when protected by password|| Path to key store<br>`path`| string| ✅| |  |  | Path to the key store file|| No properties | | | | | | | 

#### Key store: PKCS#12 / PFX with content `type = "PKCS12"`
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Content<br>`content`| string| ✅| |  |  | Binary content as Base64|| Password<br>`password`| string| ✅| |  |  | Password to use to open the key store|| Alias for the key<br>`alias`| string|  | |  |  | Alias of the key to use in case the key store contains more than one key|| Key Password<br>`keyPassword`| string|  | |  |  | Password to use to access the key when protected by password|| No properties | | | | | | | 

#### Key store: PEM with path `type = "PEM"`
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Path to cert file<br>`certPath`| string| ✅| |  |  | Path to cert file (.PEM)|| Path to private key file<br>`keyPath`| string| ✅| |  |  | Path to private key file (.PEM)|| No properties | | | | | | | 

#### Key store: PEM with content `type = "PEM"`
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Certificate<br>`certContent`| string| ✅| |  |  | || Private key<br>`keyContent`| string| ✅| |  |  | || No properties | | | | | | | 

#### Request Headers (Array)
| Name <br>`json name`  | Type <br>(constraint)  | Mandatory  | Default  | Supports <br>EL  | Supports <br>Secrets | Description  |
|:----------------------|:-----------------------|:----------:|:---------|:----------------:|:--------------------:|:-------------|
| Name<br>`name`| string|  | |  |  | || Value<br>`value`| string|  | | ✅| ✅| || No properties | | | | | | | 


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

