package web

import "mading/model/domain"

// TODO: All request and response should be here
type UsersResponse struct {
	Email    string               `json:"email"`
	Username string               `json:"username"`
	Fullname string               `json:"fullname"`
	Trxs     []domain.Transaction `json:"transactions"`
}

type UserRequest struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Time     int64  `json:"time" binding:"required"`
}
