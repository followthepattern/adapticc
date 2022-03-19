package models

type UserListRequest struct {
	ListRequest
	Name  *string
	Email *string
}
