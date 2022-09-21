package models

type UserListRequest struct {
	ListRequest
	Search *string
}
