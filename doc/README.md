# How to deploy use Docker

Add the following in Redis database to serve goshorty :

| Key (Path)                                                | Value                 |
|-----------------------------------------------------------|-----------------------|
| `traefik/http/routers/goshorty/rule`                      | `Host(`example.com`)` |
| `traefik/http/routers/goshorty/tls`                       | `true`                |
| `traefik/http/routers/goshorty/tls/certresolver`          | `myresolver`          |
| `traefik/http/services/goshorty/loadbalancer/server/port` | `3000`                |

