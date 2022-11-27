# coffee2go

Coffee credit tracker

``` go run . ``` to start the server

```curl http://localhost:8080/users``` to get a list of users

```curl http://localhost:8080/user/{UUID}``` to find a user by ID

```curl http://localhost:8080/user --include --header "Content-Type: application/json" --request "POST" --data '{"name": "Foo","email": "foo@transferwise.com"}'``` to add a new person

```curl http://localhost:8080/transactions``` to get the list of transactions

```curl http://localhost:8080/transaction/{UUID}``` to get a transaction by ID

```curl -X POST -d '{"owner": {UUID}, "recipients": ["{UUID1}", "{UUID2}"]}' http://localhost:8080/transaction``` to create a new transaction