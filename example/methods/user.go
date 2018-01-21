package methods

// User contains user data.
// @mongo_methods
type User struct {
	PublicID string `json:"public_id"`
	Name     string `json:"name"`
}
