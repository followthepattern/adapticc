package models

type UserListResponse struct {
	ListResponse
	Data []User `json:"data"`
}
