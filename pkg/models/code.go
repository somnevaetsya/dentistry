package models

type Code struct {
	UserId uint64 `json:"user_id"`
	Code   string `json:"code"`
}
