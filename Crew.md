##Crew REST API Specification

### Get JWT(JSON Web Token) Token

`POST /v1/token`

Get a *JWT*([JSON Web Token](https://tools.ietf.org/html/draft-ietf-oauth-json-web-token-32)) token.

#### Parameters

- *service* - The name of the service which hosts the resource, as: `containerops.me`, or `v1.containerops.me`.
- *scope* - The resource in question, formatted as one of the space-delimited entries from the `scope` parameters from the `WWW-Authenticate` header shown above. This query parameter should be specified multiple times if there is more than one scope entry from the `WWW-Authenticate` header. The above example would be specified as: `scope=repository:genedna/gennto:push` or `scope=repository:compose:genedna/*:*`.
- *account* - The name of the account which the client is acting as. Optional if page without signin with page grant type. 
- *grant* - The type of client, like web page, docker client or rkt client and so on.

#### Request 

```                                                                                                   
POST /v1/token                                                                                         
HOST: containerops.me                                                                             
Accept: application/json                                                                              
Content-Type: application/json                                                                        

{
  "service": "containerops.me",
  "scope": "scope=web:/v1/user:post;web:/v1/user/auth:post",
  "account": "",
  "grant": "web"
}
```

#### Response

```
HTTP/1.1 201 
Content-Type: application/json

{"token": "eyJ0eXAiOiJKV1QiLCJhbGciOiJFUzI1NiIsImtpZCI6IlBZWU86VEVXVTpWN0pIOjI2SlY6QVFUWjpMSkMzOlNYVko6WEdIQTozNEYyOjJMQVE6WlJNSzpaN1E2In0.eyJpc3MiOiJhdXRoLmRvY2tlci5jb20iLCJzdWIiOiJqbGhhd24iLCJhdWQiOiJyZWdpc3RyeS5kb2NrZXIuY29tIiwiZXhwIjoxNDE1Mzg3MzE1LCJuYmYiOjE0MTUzODcwMTUsImlhdCI6MTQxNTM4NzAxNSwianRpIjoidFlKQ08xYzZjbnl5N2tBbjBjN3JLUGdiVjFIMWJGd3MiLCJhY2Nlc3MiOlt7InR5cGUiOiJyZXBvc2l0b3J5IiwibmFtZSI6InNhbWFsYmEvbXktYXBwIiwiYWN0aW9ucyI6WyJwdXNoIl19XX0.QhflHPfbd6eVF4lM9bwYpFZIV0PfikbyXuLx959ykRTBpe3CYnzs6YBK8FToVb5R47920PVLrh8zuLzdCr9t3w"}
```

- *token* - A signed token.
