# `berr`: better errors for go

`berr` is an errors package that provides simple functions for creating more descriptive errors.

```
go get github.com/rheisen/berr
```

### Why `berr`

### Error Types

* `berr.ApplicationErrorType`
* `berr.AuthenticationErrorType`
* `berr.AuthorizationErrorType`
* `berr.ValueInvalidErrorType`
* `berr.ValueMissingErrorType`
* `berr.NotFoundErrorType`

### Getting Information from `berr.Error`

* `Error.Type()`
* `Error.Message()`
* `Error.Details()`
* `Error.Map()`
* `Error.Unwrap()`

### Examples

```go
```
