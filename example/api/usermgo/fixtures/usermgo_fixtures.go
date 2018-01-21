package fixtures

import (
     "encoding/json"


     "github.com/gokit/mgokit/example/api"

)


// json fixtures ...
var (
 UserJSON = `{


    "public_id":	"7agfuiblb2t3fkaei77gmg0mucw9az",

    "name":	"Diane Clark"

}`
)

// LoadUserJSON returns a new instance of a api.User.
func LoadUserJSON(content string) (api.User, error) {
	var elem api.User

	if err := json.Unmarshal([]byte(content), &elem); err != nil {
		return api.User{}, err
	}

	return elem, nil
}

