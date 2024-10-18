# Contributing Guidelines

1. Please fork this repo (acmpesuecc/Onyx) to your own account and work on your fork of this repo.
2. Create a PR from your fork to this repo (remember to reference the correct issue in your PR)

# Testing
- There exists a `main_test.go` file that has unit tests for every endpoint.
- **For every endpoint that you add/edit, you are expected to add/modify unit tests accordingly**
- You may use `go test` to run the entire test suite or `go test -run <unit_test_fn_name>` to run a specific unit test

# Architecture
This project uses 
- [gin](https://github.com/gin-gonic/gin) as its HTTP library
- [sqlc](https://sqlc.dev/) generated code for making DB calls
- sqlite as the DB. 

# API Spec
The API specification can be found in the `openapi.yaml` file which can be visualized using the [Swagger Editor](https://editor.swagger.io/)

You are expected to modify the `openapi.yaml` file to **document the endpoints you have edited/created**
# Database Schema

<img width="771" alt="image" src="https://github.com/user-attachments/assets/6042db8b-7675-472c-bf6c-71c98bfbd967">

Currently only the Restaurant table is non empty and it is populated with ~51k restaurants ingested from a zomato dataset
# Using `sqlc`

If you want to add or modify a DB operation, make the necesarry changes in `db/query.sql` and run the following command in the root directory of the project
```shell
sqlc generate
```
For more info on using `sqlc` refer to the [docs](https://docs.sqlc.dev/en/stable/tutorials/getting-started-sqlite.html)
