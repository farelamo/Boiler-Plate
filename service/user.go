package service

import (
	"errors"
	"fmt"
	"mading/model/web"

	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Service) GetUserWithTrx(email, username string) (web.UsersResponse, error) {

	user, err := s.repository.TrxFindUser(email, username)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return web.UsersResponse{}, fmt.Errorf("user not found")
		}
		s.log.Error(err.Error())
		return web.UsersResponse{}, fmt.Errorf("error while search on database")
	}

	trxs, err := s.repository.FindTrxByUserId(user.Id)
	if err != nil {
		s.log.Error(err.Error())
		return web.UsersResponse{}, fmt.Errorf("error while search on database")
	}

	res := web.UsersResponse{
		Email:    user.Email,
		Username: user.Username,
		Fullname: user.FirstName + user.LastName,
		Trxs:     trxs,
	}

	return res, nil
}
