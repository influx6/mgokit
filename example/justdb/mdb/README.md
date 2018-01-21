MongoDB Implementation
===============================

MongoDB Implementation is a code generation package that provides a basic implementation
for interacting with a underine mongodb database.

Following methods are implemented:

## WithIndex

```go
AddIndex(ctx context.Context, db MongoDB, col string, indexes ...mgo.Index) error) error
```

## Count

```go
Count(ctx context.Context, db MongoDB, col string) (int, error)
```

## Exec

```go
Exec(ctx context.Context, db MongoDB, col string, isread bool, fx func(col *mgo.Collection) error) error
```


