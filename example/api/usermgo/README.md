User MongoDB API
===================================
[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/mgokit/example/api/usermgo)](https://goreportcard.com/report/github.com/gokit/mgokit/example/api/usermgo)

User MongoDB API is a auto-generated CRUD implementation for the `User` in package `github.com/gokit/mgokit/example/api`.

The following method exists for custom operations:

## Exec

```go
Exec(ctx context.Context, fx func(col *mgo.Collection) error) error
```

The following methods exists in the generated API as pertaining to CRUD:

## Count

```go
Count(ctx context.Context) (int, error)
```

## Create

```go
Create(ctx context.Context, elem api.User) error
```

## Get

```go
Get(ctx context.Context, publicID string) (api.User, error)
```

## Get All

```go
GetAll(ctx context.Context) ([]api.User, error)
```

## Update

```go
Update(ctx context.Context, publicID string, elem api.User) error
```

## Delete

```go
Delete(ctx context.Context, publicID string) error
```