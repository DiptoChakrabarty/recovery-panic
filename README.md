# recovery-panic

This project contains example of how to deal with panics in golang webapps.

It is built using the Gin framework

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