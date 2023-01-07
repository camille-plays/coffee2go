# coffee2go

Coffee credit tracker

Using GIN web framework and GORM
Connects to sqlite3

On the root directory, ``` docker-compose up``` to start the postgres DB, then ``` go run . ``` to start the server

To get a list of users:
```bash
curl http://localhost:8080/users
``` 

To find a user by ID:
```bash
curl http://localhost:8080/user/{UUID}
```

To add a new person:
```bash
curl http://localhost:8080/user --include --header "Content-Type: application/json" --request "POST" --data '{"name": "Foo","email": "foo@transferwise.com"}'
```

To delete a user:
```bash
curl -X DELETE -d '{"id": "ba8ad36c-0e1f-400f-afc0-9839209cebe2"}' http://localhost:8080/user
```

To get a list of transactions:
```bash
curl http://localhost:8080/transactions
```

To get a transaction by ID:
```bash
curl http://localhost:8080/transaction/{UUID}
```

To create a new transaction:
```bash
curl -X POST -d '{"owner": {UUID}, "recipients": [{UUID1}, {UUID2}]}' http://localhost:8080/transaction
```
