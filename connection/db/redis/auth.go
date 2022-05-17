package redis

import (
	"context"
	"time"
)

func (r *repo) BlacklistToken(hash, token string) (bool, error) {
	err := r.Conn.Set(context.Background(), hash, token, 16*time.Hour).Err()
	if err != nil {
		return false, err
	}

	return true, nil
}
