# Gravitee APIM Gateway Environment Variables Configuration

For docker-compose, Kubernetes, Helm the environment variables are passed to the container.

JVM Options can be set using the `JAVA_OPTS` environment variable.


### 



####  
| | |
|---:|---|
|ENV| **GRAVITEE_HTTP_ALPN**|
|JVM|`-Dgravitee.http.alpn`|

<hr>


####  
| | |
|---:|---|
|ENV| **GRAVITEE_HTTP_COMPRESSIONSUPPORTED**|
|JVM|`-Dgravitee.http.compressionsupported`|

<hr>


####  
| | |
|---:|---|
|ENV| **GRAVITEE_HTTP_SECURED**|
|JVM|`-Dgravitee.http.secured`|

<hr>


####  
| | |
|---:|---|
|ENV| **GRAVITEE_HTTP_TCPKEEPALIVE**|
|JVM|`-Dgravitee.http.tcpkeepalive`|

<hr>


####  
| | |
|---:|---|
|ENV| **GRAVITEE_HTTP_HAPROXY_PROXYPROTOCOL**|
|JVM|`-Dgravitee.http.haproxy.proxyprotocol`|

<hr>


####  
| | |
|---:|---|
|ENV| **GRAVITEE_HTTP_SSL_OPENSSL**|
|JVM|`-Dgravitee.http.ssl.openssl`|

<hr>


####  
| | |
|---:|---|
|ENV| **GRAVITEE_HTTP_SSL_SNI**|
|JVM|`-Dgravitee.http.ssl.sni`|

<hr>


####  
| | |
|---:|---|
|ENV| **GRAVITEE_HTTP_SSL_KEYSTORE_WATCH**|
|JVM|`-Dgravitee.http.ssl.keystore.watch`|

<hr>


####  
| | |
|---:|---|
|ENV| **GRAVITEE_HTTP_SSL_TRUSTSTORE_WATCH**|
|JVM|`-Dgravitee.http.ssl.truststore.watch`|

<hr>


####  
| | |
|---:|---|
|ENV| **GRAVITEE_HTTP_WEBSOCKET_ENABLED**|
|JVM|`-Dgravitee.http.websocket.enabled`|

<hr>


####  
| | |
|---:|---|
|ENV| **GRAVITEE_HTTP_WEBSOCKET_PERFRAMEWEBSOCKETCOMPRESSIONSUPPORTED**|
|JVM|`-Dgravitee.http.websocket.perframewebsocketcompressionsupported`|

<hr>


####  
| | |
|---:|---|
|ENV| **GRAVITEE_HTTP_WEBSOCKET_PERMESSAGEWEBSOCKETCOMPRESSIONSUPPORTED**|
|JVM|`-Dgravitee.http.websocket.permessagewebsocketcompressionsupported`|

<hr>


####  
| | |
|---:|---|
|ENV| **GRAVITEE_KAFKA_ENABLED**|
|JVM|`-Dgravitee.kafka.enabled`|

<hr>


####  
| | |
|---:|---|
|ENV| **GRAVITEE_KAFKA_SECURED**|
|JVM|`-Dgravitee.kafka.secured`|

<hr>


####  
| | |
|---:|---|
|ENV| **GRAVITEE_KAFKA_TCPKEEPALIVE**|
|JVM|`-Dgravitee.kafka.tcpkeepalive`|

<hr>


####  
| | |
|---:|---|
|ENV| **GRAVITEE_KAFKA_SSL_OPENSSL**|
|JVM|`-Dgravitee.kafka.ssl.openssl`|

<hr>


####  
| | |
|---:|---|
|ENV| **GRAVITEE_KAFKA_SSL_SNI**|
|JVM|`-Dgravitee.kafka.ssl.sni`|

<hr>


####  
| | |
|---:|---|
|ENV| **GRAVITEE_KAFKA_SSL_KEYSTORE_WATCH**|
|JVM|`-Dgravitee.kafka.ssl.keystore.watch`|

<hr>


####  
| | |
|---:|---|
|ENV| **GRAVITEE_KAFKA_SSL_TRUSTSTORE_WATCH**|
|JVM|`-Dgravitee.kafka.ssl.truststore.watch`|

<hr>


####  Enable Kubernetes
| | |
|---:|---|
|ENV| **GRAVITEE_SECRETS_KUBERNETES_ENABLED**|
|JVM|`-Dgravitee.secrets.kubernetes.enabled`|

<hr>


####  Resolution namespace
| | |
|---:|---|
|ENV| **GRAVITEE_SECRETS_KUBERNETES_NAMESPACE**|
|JVM|`-Dgravitee.secrets.kubernetes.namespace`|
|Default| `gravitee`|
Default is the namespace where gravitee is deployed
<hr>


####  Connection timeout in seconds
| | |
|---:|---|
|ENV| **GRAVITEE_SECRETS_VAULT_CONNECTTIMEOUTSEC**|
|JVM|`-Dgravitee.secrets.vault.connecttimeoutsec`|
|Default| `3`|

<hr>


####  Enable Hashicorp Vault
| | |
|---:|---|
|ENV| **GRAVITEE_SECRETS_VAULT_ENABLED**|
|JVM|`-Dgravitee.secrets.vault.enabled`|

<hr>


####  Host (IP or name) of the Vault instance
| | |
|---:|---|
|ENV| **GRAVITEE_SECRETS_VAULT_HOST**|
|JVM|`-Dgravitee.secrets.vault.host`|
|Default| `127.0.0.1`|

<hr>


####  Key-Value engine, no mixing supported
| | |
|---:|---|
|ENV| **GRAVITEE_SECRETS_VAULT_KVENGINE**|
|JVM|`-Dgravitee.secrets.vault.kvengine`|
|Default| `V2`|
|Values| `V1` `V2` |

<hr>


####  Vault namespace
| | |
|---:|---|
|ENV| **GRAVITEE_SECRETS_VAULT_NAMESPACE**|
|JVM|`-Dgravitee.secrets.vault.namespace`|
|Default| `default`|

<hr>


####  Port of the Vault instance
| | |
|---:|---|
|ENV| **GRAVITEE_SECRETS_VAULT_PORT**|
|JVM|`-Dgravitee.secrets.vault.port`|
|Default| `8082`|

<hr>


####  Read timeout in seconds
| | |
|---:|---|
|ENV| **GRAVITEE_SECRETS_VAULT_READTIMEOUTSEC**|
|JVM|`-Dgravitee.secrets.vault.readtimeoutsec`|
|Default| `2`|

<hr>


####  Enable secured connection to Vault
| | |
|---:|---|
|ENV| **GRAVITEE_SECRETS_VAULT_SSL_ENABLED**|
|JVM|`-Dgravitee.secrets.vault.ssl.enabled`|

<hr>


####  Format of the client certificate or CA if not using a public one
| | |
|---:|---|
|ENV| **GRAVITEE_SECRETS_VAULT_SSL_FORMAT**|
|JVM|`-Dgravitee.secrets.vault.ssl.format`|
|Default| `pem`|
|Values| `pem` `pemfile` `truststore` |
|When| `format = 'pem'` |

<hr>


####  Content of the PEM file with headers
| | |
|---:|---|
|ENV| **GRAVITEE_SECRETS_VAULT_SSL_PEM**|
|JVM|`-Dgravitee.secrets.vault.ssl.pem`|
|Default| `--- BEGIN CERTIFICATE ---
MIIFxjCCA64CCQD9kAnHVVL02TANBgkqhkiG...
--- END CERTIFICATE ---
`|
|When| `format = 'pem'` |

<hr>


####  Location of the file
| | |
|---:|---|
|ENV| **GRAVITEE_SECRETS_VAULT_SSL_FILE**|
|JVM|`-Dgravitee.secrets.vault.ssl.file`|
|Default| `ssl/cert.pem`|
|When| `format = 'pemfile'` |

<hr>


####  
| | |
|---:|---|
|ENV| **GRAVITEE_SECRETS_VAULT_WATCH_ENABLED**|
|JVM|`-Dgravitee.secrets.vault.watch.enabled`|

<hr>


####  
| | |
|---:|---|
|ENV| **GRAVITEE_TCP_ENABLED**|
|JVM|`-Dgravitee.tcp.enabled`|

<hr>


####  
| | |
|---:|---|
|ENV| **GRAVITEE_TCP_SECURED**|
|JVM|`-Dgravitee.tcp.secured`|

<hr>


####  
| | |
|---:|---|
|ENV| **GRAVITEE_TCP_TCPKEEPALIVE**|
|JVM|`-Dgravitee.tcp.tcpkeepalive`|

<hr>


####  
| | |
|---:|---|
|ENV| **GRAVITEE_TCP_HAPROXY_PROXYPROTOCOL**|
|JVM|`-Dgravitee.tcp.haproxy.proxyprotocol`|

<hr>


####  
| | |
|---:|---|
|ENV| **GRAVITEE_TCP_SSL_OPENSSL**|
|JVM|`-Dgravitee.tcp.ssl.openssl`|

<hr>


####  
| | |
|---:|---|
|ENV| **GRAVITEE_TCP_SSL_SNI**|
|JVM|`-Dgravitee.tcp.ssl.sni`|

<hr>


####  
| | |
|---:|---|
|ENV| **GRAVITEE_TCP_SSL_KEYSTORE_WATCH**|
|JVM|`-Dgravitee.tcp.ssl.keystore.watch`|

<hr>


####  
| | |
|---:|---|
|ENV| **GRAVITEE_TCP_SSL_TRUSTSTORE_WATCH**|
|JVM|`-Dgravitee.tcp.ssl.truststore.watch`|

<hr>


