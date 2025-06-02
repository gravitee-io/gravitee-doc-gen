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
<no value>
```

## Environment variables

<no value>

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
