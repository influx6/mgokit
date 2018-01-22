MgoKit
--------
[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/mgokit)](https://goreportcard.com/report/github.com/gokit/mgokit)

Mgokit implements a code generator which automatically generates go packages for mongodb implementations for annotated structs.

## Install

```
go get -u github.com/gokit/mgokit
```

## Annotations

- `@mongapi`

Generate API implementation structure for a struct.

- `@mongo_method`

Generate package-level functions for CRUD operation with struct.

- `@mongo`

Generate simple package with `Exec` function for interacting with mongodb.


## Examples

See [Examples](./example) for all usage of provided annotations for code generation with associated results.

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

- Generating CRUD methods implementation

You annotate any giving struct with `@mongo_methods` which marks giving struct has a target for code generation. 

*All struct must have a `PublicID` field.*

Sample below:

```go
// User is a type defining the given user related fields for a given.
// @mongo_methods
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

Mgokit will generate a package which contains functions suitable for performing all CRUD operation, it defers from the `@mongoapi` which creates a DB struct which encapsulates all methods to itself necessary to run against a given collection, where as this creates package level functions to perform such operations, that require more input.

- Generating API CRUD implementation

You annotate any giving struct with `@mongoapi` which marks giving struct has a target for code generation. 

*All struct must have a `PublicID` field.*

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

## Interfaces

Sqlkit will generate specific interfaces where each allows the internal functions validate struct validity and are able to get and consume maps containing record details.


```go
type Validate interface{
	Validate() error
}
```

These are required and must be added to target struct as 
they are the means the generated code uses to get data to be saved and sets internal struct's fields:

```go
type UserFields  interface {
	Fields() map[string]interface{}
}

type UserConsumer interface {
	Consume(map[string]interface{}) error
}
```

