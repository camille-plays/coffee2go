# coffee2go
Backend logic for the application that tracks coffee credit across the team


to run:

``` go run . ```


```curl http://localhost:8080/people``` gets a list of people


```curl http://localhost:8080/people --include --header "Content-Type: application/json" --request "POST" --data '{"id": "7","name": "Yagiz","email": "yagiz@transferwise.com","credit": 0}'``` to add a new person