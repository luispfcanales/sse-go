package domain

import "time"

//User is template struct
type User struct {
	ID         string    `json:"id,omitempty"`
	Email      string    `json:"email,omitempty"`
	Password   string    `json:"password,omitempty"`
	Name       string    `json:"name,omitempty"`
	GivenName  string    `json:"given_name,omitempty"`
	FamilyName string    `json:"family_name,omitempty"`
	Picture    string    `json:"picture,omitempty"`
	Expiry     time.Time `json:"expiry,omitempty"`
}
