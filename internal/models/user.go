package models

type User struct {
	UID     int64  `json:"uid,string,omitempty"`
	Name    string `json:"name,string,omitempty"`
	Email   string `json:"email,string,omitempty"`
	Age     int64  `json:"age,string,omitempty"`
	Deleted bool   `json:"-"`
}
