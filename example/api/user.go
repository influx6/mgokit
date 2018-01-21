package api

// User contains user data.
// @mongoapi
type User struct {
	PublicID string `json:"public_id"`
	Name     string `json:"name"`
}
