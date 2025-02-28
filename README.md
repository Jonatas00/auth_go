# Go authentication API

Go basic API with JWT authentication

## Technologies Used

- Golang
- Gin Gonic
- GORM
- JWT (JSON Web Tokens)
- Docker
- PostgreSQL

## Getting Started

First create a ".env" file in the root directory with these variables

```bash
# Default port for application
APP_PORT=8080
# Token for JWT signature 
JWT_TOKEN=tokenhere

# Default database variables to open connection
DB_HOST=database
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=postgres
DB_PORT=5432
```

You need docker installed to run this project

```bash
docker compose up
```

This command will run the app on default route "localhost:8080"

## Usage

- Using an API Client to request API:
  
  - You can use a client like [postman](https://www.postman.com) or [Insomnia](https://insomnia.rest/download)

- Signup route:
  - Method: POST
  - This route register the user in database
  - ```localhost:8080/signup```

```json
  // Params:
  {
    "name": "",
    "email": "",
    "password": ""
  }
```

This route return a json like a this:

```json
  {
    "message": "successfully signed up",
    "user_id": 1
  }
```

- Login route:
  - Method: POST
  - This route validate user and make cookies with the authorization token
  - ```localhost:8080/login```

```json
  // Params:
  {
    "email": "",
    "password": ""
  }
```

This route return a json like a this:

```json
  {
    "message": "Login Successful",
    "token": "JWT Token"
    // This token is added to authorization cookie
  }
```

- Get Users route:
  - Method: GET
  - This route verify the JWT on cookies and return all users on database
  - ```localhost:8080/getusers```

This route return a json like a this:

```json
  [
    {
      "ID": 1,
      "CreatedAt": "2024-06-01T21:05:04.735348Z",
      "UpdatedAt": "2024-06-01T21:05:04.735348Z",
      "DeletedAt": null,
      "name": "Jonatas",
      "email": "djonatas31@gmail.com"
    },
    {
      "ID": 3,
      "CreatedAt": "2024-06-01T21:13:03.905145Z",
      "UpdatedAt": "2024-06-01T21:13:03.905145Z",
      "DeletedAt": null,
      "name": "Eduardo",
      "email": "Edu_test@outlook.com.br"
    }
  ]
```
