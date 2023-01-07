# coffee2go

Coffee credit tracker

Using GIN web framework and GORM
Connects to sqlite3

On the root directory, ``` docker-compose up``` to start the postgres DB, then ``` go run . ``` to start the server

```curl http://localhost:8080/users``` to get a list of users

```curl http://localhost:8080/user/{UUID}``` to find a user by ID

To add a new person:

```bash
curl http://localhost:8080/user --include --header "Content-Type: application/json" --request "POST" --data '{"name": "Foo","email": "foo@transferwise.com"}'
```

To delete a user:

```bash
curl -X DELETE -d '{"id": "ba8ad36c-0e1f-400f-afc0-9839209cebe2"}' http://localhost:8080/user
```






```curl http://localhost:8080/transactions``` to get the list of transactions

```curl http://localhost:8080/transaction/{UUID}``` to get a transaction by ID

```curl -X POST -d '{"owner": {UUID}, "recipients": [{UUID1}, {UUID2}]}' http://localhost:8080/transaction``` to create a new transaction
