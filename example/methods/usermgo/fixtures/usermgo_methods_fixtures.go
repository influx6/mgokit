package fixtures

import (
     "encoding/json"


     "github.com/gokit/mgokit/example/methods"

)


// json fixtures ...
var (
 UserJSON = `{


    "public_id":	"y3w9h93qwhy66bk0parnuf59oisvda",

    "name":	"Frank Hart"

}`
)

// LoadUserJSON returns a new instance of a methods.User.
func LoadUserJSON(content string) (methods.User, error) {
	var elem methods.User

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return methods.User{}, err
	}

	return elem, nil
}

