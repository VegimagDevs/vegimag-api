package storage

import (
	bolt "go.etcd.io/bbolt"
)

func (storage *Storage) HasRefreshToken(refreshToken string) (bool, error) {
	hasRefreshToken := false

	if err := storage.db.View(func(tx *bolt.Tx) error {
		refreshTokensBucket := tx.Bucket([]byte(REFRESH_TOKENS_BUCKET))

		if refreshTokensBucket.Get([]byte(refreshToken)) != nil {
			hasRefreshToken = true
		}

		return nil
	}); err != nil {
		return false, err
	}

	return hasRefreshToken, nil
}
