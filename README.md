
This is simple server written by Golang using Gin and MongoDB.
# Description
This is just the first step of the project. Still very simple, we could build more features in the future. Some features we could build in near future:
- Call to blockchain
- Authentication and Authorization
- Logger system and Notification to handle errors

# Document
- https://drive.google.com/drive/folders/1Saa3Tj4IjB0Q1kiHqvCeA5Anuu09z8F6?usp=drive_link

# Infrastructure
- Golang : Gin framework
- DB: Mongodb

# Why MongoDB?
-> Because we did not have the final business logic, so that the data structure will be changed in the future. So that we need a database which could help us easy to modify the structure.

# The structure
```
├───cmd
├───config
├───internal
│   ├───api
│   │   ├───handler
│   │   └───presenter
│   ├───common
│   │   ├───cmentity
│   │   ├───constant
│   │   └───util
│   ├───docs
│   ├───entity
│   │   └───request
│   ├───infrastructure
│   │   ├───repository
│   │   └───router
│   ├───registry
│   └───usecase
│       ├───attendance
│       ├───employeemanagement
│       ├───userdoaction
│       └───usergetinfo
├───pkg
│   ├───blockchain
│   │   └───chain
│   └───mongo
└───resources
```

# How to run
Please make sure that you have Golang in your PC or laptop and also install go swagger. Run the following command:

```
swag init -g internal/infrastructure/router/router.go -o internal/docs && go mod tidy && go mod download && go run ./cmd
```
