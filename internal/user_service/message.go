package userservice

type UserCreateInfo struct {
	UserId int64  `json:"user_id"`
	Name   string `json:"name"`
	Age    int64  `json:"age"`
	Email  string `json:"email"`
	Time   int64  `json:"time"`
}
