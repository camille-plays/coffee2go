# coffee2go
Backend logic for the application that tracks coffee credit across the team


to run:

``` go run . ``` to start the server


```curl http://localhost:8080/people``` to get a list of people


```curl http://localhost:8080/people --include --header "Content-Type: application/json" --request "POST" --data '{"id": "7","name": "Yagiz","email": "yagiz@transferwise.com","credit": 0}'``` to add a new person

```curl http://localhost:8080/people/2``` to find a person by ID