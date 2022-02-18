# recovery-panic

This project contains example of how to deal with panics in golang webapps.

It is built using the Gin framework

### How to Run 

```sh
- Clone Repository
- Install all go dependencies
- go run main.go
```

The following paths are present

1) / :- home page
2) /panic :- does panic before
3) /paniclater :- splits a message then panics
4) /panicrandom :- panics but handles it by default

### Parts of Code

* middleware
```sh
- This contains the custom middleware developed to support gin webapps
- The RecoverPanic middleware is used to handle panic cases and return helpful error messages in webapp
```

* routes
```sh
- This contains all the routes defined for the gin webapp
- Some routes have a direct panic whereas some have a message then a panic
```

* types
```sh
This contains a gin ResponseWriter for further implementation
```