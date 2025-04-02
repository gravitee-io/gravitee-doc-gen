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


### Configuration options

## Phases
The phases checked below are supported by the `test` policy:
| v2 Phases| Compatible?| v4 Phases| Compatible? |
| --- | --- | --- | ---  |
| onRequest|  | onRequest|   |
| onResponse|  | onResponse|   |
| onRequestContent| X| onMessageRequest| X |
| onResponseContent| X| onMessageResponse| X |

## Compatibility matrix
Strikethrough line are deprecated versions

| Plugin version| APIM| AM| Cockpit| Comment |
| --- | --- | --- | --- | ---  |
|~~1.0~~|~~3.x~~|~~2.1~~|~~-~~|~~-~~ |
|2.x|4.x|4.1|-|Incompatible with cloud |
|3.x|4.6 and above|4.6 and above|-|- |

