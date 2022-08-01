## A simple implementation of a expense tracker API in GO. 
---

Application implements a simple weekly/monthly expense tracker. User can add budgets and expenses, as well as see remaining budget by category. 
This is an example project to showcase an understanding of Go.


Runs on localhost:8080 and allows very basic interactions GET, POST, PATCH requests.

Example POST request for an expense
```
curl --header "Content-Type: application/json" \
  --request POST \
  --data @exp1.json \
  http://localhost:8080/expenses?id=3
```
Example POST request for budget

```
curl --header "Content-Type: application/json" \
 --request POST \
 --data @budget1.json \
 http://localhost:8080/budgets

```
