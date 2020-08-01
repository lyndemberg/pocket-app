# pocket-app
A simple web server using Golang for training my skills in language.

## What is pocket-app?
Pocket-app is a simple domain model that aims to control users' spending / finances.

## Future purposes
Create a client application developed in react.js, to also develop my skills in this other technology

### Dependencies used
```
- github.com/gorilla/mux
- github.com/go-sql-driver/mysql
- github.com/joho/godotenv
- github.com/dgrijalva/jwt-go
```

### Endpoints
Taking into account that pocket-app runs on port 8080, as described in [`main.go`](main.go). So the base path is: `http://localhost:8080`.

#### Users endpoint
```
GET /api/users
```
```
POST /api/users
Content-Type: application/json
{
    "name": "John Doe",
    "email": "johndoe@gmail.com",
    "username": "johndoe",
    "password": "anypasswordhere"
}
```
```
PUT /api/users
Content-Type: application/json
{
    "id":   "453"
    "name": "John Doe update",
    "email": "johndoe71@gmail.com",
    "username": "johndoe",
    "password": "newpassword"
}
```
```
GET /api/users/{id}
```
```
DELETE /api/users/{id}
```

#### Authentication endpoint
```
POST /login 
Content-Type: multipart/form-data
username=johndoe
password=johndoe@gmail.com
```