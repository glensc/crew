##Crew REST API Specification

### Get JWT(JSON Web Token) Token

`POST /v1/token`

Get a *JWT*([JSON Web Token](https://tools.ietf.org/html/draft-ietf-oauth-json-web-token-32)) token.

#### Parameters

- *service* - The name of the service which hosts the resource, as: `containerops.me`, or `v1.containerops.me`.
- *scope* - The resource in question, formatted as one of the space-delimited entries from the `scope` parameters from the `WWW-Authenticate` header shown above. This query parameter should be specified multiple times if there is more than one scope entry from the `WWW-Authenticate` header. The above example would be specified as: `scope=repository:genedna/gennto:push` or `scope=repository:compose:genedna/*:*`.
- *account* - The name of the account which the client is acting as. Optional if page without signin with page grant type. 
- *grant* - The type of client, like web page, docker client or rkt client and so on.

#### Request Example

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

#### Response Example

```
HTTP/1.1 201 
Content-Type: application/json

{"token": "eyJ0eXAiOiJKV1QiLCJhbGciOiJFUzI1NiIsImtpZCI6IlBZWU86VEVXVTpWN0pIOjI2SlY6QVFUWjpMSkMzOlNYVk"}
```

- *token* - A signed token.

#### HTTP Status Code

- *201* -
- *400* -
- *500* -

### Create An User

`POST /v1/user`

Sign up with a user account.

#### Parameters

- *username* - The name of the user, must match the regular expression `[a-z0-9]+(?:[._-][a-z0-9]+)*`
- *email* - the email address of user.
- *passwd* - the password of the user.

#### Request Example

```                                                                                                   
POST /v1/user
HOST: containerops.me                                                                             
Accept: application/json                                                                              
Content-Type: application/json                                                                        
Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJFUzI1NiIsImtpZCI6IlBZWU86VEVXVTpWN0pIOjI2SlY6QVFUWjpMSkMzOlNYVk

{
  "username": "genedna",
  "email": "genedna@gmail.com",
  "passwd": "genednarandompassword"
}
```

#### Response Example

```
HTTP/1.1 201 
Content-Type: application/json

```

#### HTTP Status Code

- *201* -
- *400* -
- *500* -

### User Sign In 

`POST /v1/user/auth`

User sign in the system from web page.

#### Parameters

- *username* - The name of the user, must match the regular expression `[a-z0-9]+(?:[._-][a-z0-9]+)*`
- *email* - the email address of user.

#### Request Example

```
POST /v1/user/auth
HOST: containerops.me                                                                             
Accept: application/json                                                                              
Content-Type: application/json                                                                        
Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJFUzI1NiIsImtpZCI6IlBZWU86VEVXVTpWN0pIOjI2SlY6QVFUWjpMSkMzOlNYVk

{
  "email": "genedna@gmail.com",
  "passwd": "genednarandompassword"
}
```

#### Response Example

```
HTTP/1.1 200
Content-Type: application/json

{"token":"eyJ0eXAiOiJKV1QiLCJhbGciOiJFUzI1NiIsImtpZCI6IlBZWU86VEVXVTpWN0pIOjI2SlY6Q"}
```

When sign in the system successfully, and will return a new token value for the user

#### HTTP Status Code

- *200* -
- *400* -
- *500* -

### Get users list 

`GET /v1/user/list/:count/:page`

Get users of system with count and page parameters.

#### Parameters

- *count* - Amount users from service per request.   
- *page* - The page index which start with 0. 

#### Request Example

```
GET /v1/user/list/30/0
HOST: containerops.me                                                                             
Accept: application/json                                                                              
Content-Type: application/json                                                                        
Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJFUzI1NiIsImtpZCI6IlBZWU86VEVXVTpWN0pIOjI2SlY6QVFUWjpMSkMzOlNYVk
```

#### Response Example

```
HTTP/1.1 200
Content-Type: application/json

```

#### HTTP Status Code

- *200* -
- *400* -
- *500* -

### PUT User Profile

`PUT /v1/user/:user/profile`

Set user's profile data.

#### Parameters

- *user* - User's username

#### Request Example

```
PUT /v1/user/containerme/profile
HOST: containerops.me                                                                             
Accept: application/json                                                                              
Content-Type: application/json                                                                        
Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJFUzI1NiIsImtpZCI6IlBZWU86VEVXVTpWN0pIOjI2SlY6QVFUWjpMSkMzOlNYVk

{
  "username": "containerme",
  "email": "genedna@gmail.com",
  "fullname": "Meaglith Ma",
  "company": "",
  "location": "Beijing, China",
  "mobile": "+86186xxxxxxxx",
  "url": "https://containerops.me",
  "gravatar": "https://containerops.me/profiles/containeropsme/gravatar.png"
}
```

#### Response Example
```
HTTP/1.1 200
Content-Type: application/json

```

#### HTTP Status Code

- *200* -
- *400* -
- *500* -

### Get User Profile

`GET /v1/user/containerme/profile`

Get user's profile data.

#### Parameters

- *user* - User's username

#### Request Example

```
GET /v1/user/containerme/profile
HOST: containerops.me                                                                             
Accept: application/json                                                                              
Content-Type: application/json                                                                        
Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJFUzI1NiIsImtpZCI6IlBZWU86VEVXVTpWN0pIOjI2SlY6QVFUWjpMSkMzOlNYVk
```

#### Response Example

```
HTTP/1.1 200
Content-Type: application/json

{
  "username": "containerme",
  "email": "genedna@gmail.com",
  "fullname": "Meaglith Ma",
  "company": "",
  "location": "Beijing, China",
  "mobile": "+86186xxxxxxxx",
  "url": "https://containerops.me",
  "gravatar": "https://containerops.me/profiles/containeropsme/gravatar.png"
}
```

#### HTTP Status Code

- *200* -
- *400* -
- *500* -


