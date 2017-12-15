MgoKit
--------
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

It also respects `@associates` annotation which gives extra information to the generator for the following data:

1. What struct to be used as representing a new struct type.
2. What struct contain information for representing the updated struct type.

Sample below:

```go
// User is a type defining the given user related fields for a given.
// @mongoapi
//@associates(@mongoapi, New, NewUser)
//@associates(@mongoapi, Update, UpdatedUser)
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

Sample Interface for `User` struct:

```go
type UserFields  interface {
	Fields() map[string]interface{}
}

type UserBSON interface {
	BSON() bson.M
}

type UserBSONConsumer interface {
	BSONConsume(bson.M) error
}

type UserConsumer interface {
	Consume(map[string]interface{}) error
}
```