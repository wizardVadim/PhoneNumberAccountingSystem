package models

type User struct {
	Id       int64  `json:"id"`
	Login    string `json:"login"`
	Password string `json:"-"`
	RoleId   int64  `json:"role_id"`
	RoleName string  `json:"role_name"`
}
