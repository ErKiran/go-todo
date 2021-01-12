### To Run the Server

`make run`

### User (AUTH)

1. **Create User First**
`/user`

Request Body 

``` 
{
 "email": "test@email.com",
 "password":" TestPassword"
}
```

2. Login With the Created User Detail For Auth 

`/user/login`

Request Body 

``` 
{
 "email": "test@email.com",
 "password":" TestPassword"
}
```
It will return back JWT Token which needs to be sent with every other request.
