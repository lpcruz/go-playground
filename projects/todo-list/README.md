# Todo List (HTTP Server)

This project exercises the following:

* Writing a basic HTTP Server in GO with the `net/http` package
* Processing requests (POST, GET, DELETE) to endpoints


# Usage

```bash
$ go run main.go
```

A server will boot up at http://localhost:8080

# Example Requests & Responses

```bash
$ curl -X POST -H 'Content-type: application/json' --data '{ "name": "go through emails" }' http://localhost:8080/todo-items

{
  "id":"40da9294-e0b3-495d-8fe8-6e1960817ac1",
  "name":"go through emails"}
}
```

```bash
$ curl -X POST -H 'Content-type: application/json' --data '{ "name": "eat breakfast" }' http://localhost:8080/todo-items

{
  "id":"3b1824bf-355c-43cc-97ed-811af8326fe2",
  "name":"eat breakfast"}
}
```

```bash
$ curl -X GET -H 'Content-type: application/json' http://localhost:8080/todo-items

[
  {"id":"40da9294-e0b3-495d-8fe8-6e1960817ac1","name":"go through emails"},
  {"id":"3b1824bf-355c-43cc-97ed-811af8326fe2","name":"eat breakfast"}
]
```

# Resources

* [`net/http`](https://golang.org/pkg/net/http/)
* [Your First http server in Go (YouTube)](https://www.youtube.com/watch?v=5BIylxkudaE)