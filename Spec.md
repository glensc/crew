### Crew Rest API Specification

#### User Signup

`POST /w1/user`

#### User Signin

`POST /w1/user/auth`

#### User List

`GET /w1/user/list`

#### Put User Profile

`PUT /w1/user/<user>/profile`

#### Get User Profile

`Get /w1/user/<user>/profile`

#### Post User Gravatar

`POST /w1/user/<user>/gravatar`

#### Put User Passwd

`PUT /w1/user/<user>/passwd`

#### Get User Organizations

`GET /w1/user/<user>/organizations`

#### Get User Teams

`GET /w1/user/<user>/teams`

#### Create Organization

`POST /w1/org`

#### Put Organization

`PUT /w1/org/:org`

#### Get Organization

`GET /w1/org/:org`

#### Delete Organization

`DELETE /w1/org/:org`

#### Create Team

`POST /w1/org/:org/team`

#### Get Team List

`GET /w1/org/:org/team/list`

#### Put Team 

`PUT /w1/org/:org/team/:team`

#### GET Team

`GET /w1/org/:org/team/:team`

#### Delete Team

`DELETE /w1/org/:org/team/:team`

#### Get Team User

`GET /w1/org/:org/team/:team/user/list`

#### Add User To Team

`PUT /w1/org/:org/team/:team/user/:user`

#### Delete User From A Team

`PUT /w1/org/:org/team/:team/user/:user`
