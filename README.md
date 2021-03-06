# golang-RestAPI

## Login 
- Method : GET
- Endpoint : `/login`
- Header :
  - Content-Type: application/json
- Body :
```json 
{
    "username": "admin",
    "password": "admin"
}
```
Response : 
```json 
{
    "Id": 3,
    "name": "admin",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTc0NTMzNzYsInVzZXJfaWQiOjN9.AIUcTkrHipSuuT4OY2mNqEjhN306ZS8YLgfAOdHIAq8"
}
```

Response Error : 
```json 
{
    "error": true,
    "msg": "User Not Valid"
}
```

## GetUsers 
- Method : GET
- Endpoint : `/users`
- Header :
  - Content-Type: application/json
  - Authorization: Bearer token

Response : 
```json 
{
    "code": 200,
    "data": [
        {
            "id": 3,
            "username": "admin",
            "name": "admin",
            "email": "admin@gmail.com",
            "status": true,
            "phone": "08888",
            "password": "",
            "id_role": 1
        }
    ],
    "error": false,
    "msg": null
}
```
Response Error : 
```json 
{
    "error": true,
    "msg": "Missing or malformed JWT"
}
```

## GetUsers By ID
- Method : GET
- Endpoint : `/users/{id}`
- Header :
  - Content-Type: application/json
  - Authorization: Bearer token

Response : 
```json 
{
    "id": 3,
    "username": "admin",
    "name": "admin",
    "email": "admin@gmail.com",
    "status": true,
    "phone": "08888",
    "password": "",
    "id_role": 1
}
```

## Create User
- Method : POST
- Endpoint : `/user`
- Header :
  - Content-Type: application/json
  - Authorization: Bearer token

- Body :
```json 
{
    "username": "admin2",
    "name": "admin2",
    "email": "admin2@gmail.com",
    "status": true,
    "phone": "02038932",
    "password": "admin2",
    "id_role": 1
}
```

Response : 
```json 
{
    "id": 6,
    "username": "admin2",
    "name": "admin2",
    "email": "admin2@gmail.com",
    "status": true,
    "phone": "02038932",
    "password": "admin2",
    "id_role": 1
}
```
Response Error : 
```json 
{
    "error": true,
    "msg": "Missing or malformed JWT"
}

## Update User
- Method : PUT
- Endpoint : `/user/{id}`
- Header :
  - Content-Type: application/json
  - Authorization: Bearer token

- Body :
```json 
{
    "id": 6,
    "username": "updateadmin",
    "name": "aps",
    "email": "updateadmin@gmail.com",
    "status": true,
    "phone": "02038932",
     "id_role": 1
}
```

Response : 
```json 
{
    "id": 6,
    "username": "updateadmin",
    "name": "aps",
    "email": "updateadmin@gmail.com",
    "status": true,
    "phone": "02038932",
    "password": "",
    "id_role": 1
}
```

Response Error : 
```json 
{
    "error": true,
    "msg": "Missing or malformed JWT"
}

## Delete User
- Method : DEL
- Endpoint : `/user/{id}`
- Header :
  - Content-Type: application/json
  - Authorization: Bearer token

Response : 
```json 
{
"Deleted"
}
```

