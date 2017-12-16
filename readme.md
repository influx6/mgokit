MgoKit
--------
[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/mgokit)](https://goreportcard.com/report/github.com/gokit/mgokit)

Mgokit implements a code generator which automatically generates go packages for mongodb implementations for annotated structs.

## Install

```
go get -u github.com/gokit/mgokit
```

## Usage

Running the following commands instantly generates all necessary files and packages for giving code gen.

```go
> mgokit generate
```

## How It works

### Package Annotation

You annotate a package with the `@mongo` annotation which generates a basic package to interact with an underline mongodb database.

Sample below:

```go
// Package box does ....
//@mongo
package box

```

### Struct Annotation

You annotate any giving struct with `@mongoapi` which marks giving struct has a target for code generation. 

Sample below:

```go
// User is a type defining the given user related fields for a given.
// @mongoapi
type User struct {
	Username      string    `json:"username"`
	PublicID      string    `json:"public_id"`
	PrivateID     string    `json:"private_id"`
	Hash          string    `json:"hash"`
	TwoFactorAuth bool      `json:"two_factor_auth"`
	Created       time.Time `json:"created_at"`
	Updated       time.Time `json:"updated_at"`
}
```

Mgokit expects structs to match specific interfaces, which allows it to function as seperate from the declared struct has much has possible, because the interfaces allow to perform serialization and deserialization.

Each interface is included with all generated files.

Sample interfaces to be implemented for `User` struct (only one set is needed):

```go
type UserFields  interface {
	Fields() map[string]interface{}
}

type UserConsumer interface {
	Consume(map[string]interface{}) error
}
```

```go
type UserBSONConsumer interface {
	BSONConsume(bson.M) error
}

type UserBSON interface {
	BSON() bson.M
}
```
