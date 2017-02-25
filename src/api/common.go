package api

import "github.com/seccom/kpass/src/schema"

// AuthResult ...
type AuthResult struct {
	Token  string             `json:"access_token" swaggo:"true,access_token,tokenxxxxxxxx...."`
	Type   string             `json:"token_type" swaggo:"true,will always be \"Bearer\",Bearer"`
	Expire float64            `json:"expires_in" swaggo:"true,expires time duration in seconds,3600"`
	User   *schema.UserResult `json:"user,omitempty" swaggo:"false,user info"`
}
