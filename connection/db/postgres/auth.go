package db

import "github.com/i-jonathan/pharmacy-api/model"

func (r *repo) FetchAccountWithPassword(auth model.Auth) (model.Account, error) {
	panic("implement me")
}

func (r *repo) BlacklistToken(hash, token string) (bool, error) {
	panic("implement me")
}
