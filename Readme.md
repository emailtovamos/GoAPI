# Quick Run
- `kubectl apply -f database/.`

- `kubectl apply -f devops/.`

- Open Postman

- `localhost:30900/api/user/new` -> `POST` Request

- Body: `{"email": "satya3@gmail.com", "password": "mypassword"}`

- Headers: `Content-type` - `application/json`

- Execute it and get the token

- Use this token to do the `GET` request to get the roles:

    - `localhost:8001/api/roles`

    - Authorization: Bearer Token - <token you copied before>

    - Body: `{"subject": "system"}` 

    - Execute it to get the output. 


# Architecture

Client -> Authentication/Authorisation -> REST API -> Kubeclient GO 

# Points regarding the codebase

 - **Authentication & Authorisation** -> Done using GORM and JWT

 - **Database** -> Postgresql to store user's data

 - **Actual work** -> Using kube-client Go to interact with Kubernetes API to be able to fetch the Roles etc.

 - **Containerisation** -> Docker

 - **Security** is taken care of by the above methods of Authentication and Authorisation. 

 - **Backwards compatibility** 
    - It can be generally ensured easily if the server undergoes non-breaking changes.
    - In case of breaking changes, the client needs to be updated. 
    - Regular `go mod vendor` for client
    - Go build, test and integration tests before full deployment of a client to ensure compatibility.
    - One can use gRPC to keep track of changes and make it easier for backwards compatibility by using versions.

# To run the program (Locally without Kubernetes)

Before running the main program locally: 

 - Open command line

 - `brew services start postgresql`

 - `psql postgres`

 - `CREATE DATABASE goapi;`

 - To stop: `brew services stop postgresql`

Now run `main` and create user using Postman by doing `POST` to `localhost:8001/api/user/new` with raw Body set to be whatever email and password you want to create your account. e.g. `{"email": "satya@gmail.com", "password": "mypassword"}`

Headers: `Content-type` - `application/json`

Execute it and get the token

Use this token to do the `GET` request to get the roles:

`localhost:8001/api/roles`

Authorization: Bearer Token - <token you copied before>

Body: `{"subject": "system"}` 

Execute it to get the output. 




