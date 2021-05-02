Enter this at the terminal to test application

```
curl -X POST -H "Content-Type: application/json" -d '{"name":"Miss Moneypenny","gender":"female","age":27}' http://localhost:8080/user
```

```
curl http://localhost:8080/user/<enter-user-id-here>

```


```
curl -X DELETE http://localhost:8080/user/<enter-user-id-here>
```