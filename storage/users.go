package storage

import (
	"encoding/json"
	"fmt"
	bolt "go.etcd.io/bbolt"
)

type User struct {
	Id             string `json:"id"`
	Email          string `json:"email"`
	Username       string `json:"username"`
	HashedPassword string `json:"hashedPassword"`
}

func (storage *Storage) CreateUser(user *User) error {
	userBytes, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("error marshalling the user: %w", err)
	}

	return storage.db.Update(func(tx *bolt.Tx) error {
		usersBucket := tx.Bucket([]byte(USERS_BUCKET))
		emailToUserIdBucket := tx.Bucket([]byte(EMAIL_TO_USER_ID_BUCKET))
		usernameToUserIdBucket := tx.Bucket([]byte(USERNAME_TO_USER_ID_BUCKET))

		if err := usersBucket.Put([]byte(user.Id), userBytes); err != nil {
			return err
		}

		if err := emailToUserIdBucket.Put([]byte(user.Email), []byte(user.Id)); err != nil {
			return err
		}

		if err := usernameToUserIdBucket.Put([]byte(user.Username), []byte(user.Id)); err != nil {
			return err
		}

		return nil
	})
}

func (storage *Storage) GetUserById(userId string) (*User, error) {
	var userBytes []byte

	if err := storage.db.View(func(tx *bolt.Tx) error {
		usersBucket := tx.Bucket([]byte(USERS_BUCKET))

		userBytes = usersBucket.Get([]byte(userId))

		return nil
	}); err != nil {
		return nil, err
	}

	if userBytes == nil {
		return nil, nil
	}

	user := new(User)
	if err := json.Unmarshal(userBytes, user); err != nil {
		return nil, fmt.Errorf("error unmarshalling the user: %w", err)
	}

	return user, nil
}

func (storage *Storage) GetUserByEmail(email string) (*User, error) {
	var userBytes []byte

	if err := storage.db.View(func(tx *bolt.Tx) error {
		emailToUserIdBucket := tx.Bucket([]byte(EMAIL_TO_USER_ID_BUCKET))

		userId := string(emailToUserIdBucket.Get([]byte(email)))
		if userId == "" {
			return nil
		}

		usersBucket := tx.Bucket([]byte(USERS_BUCKET))

		userBytes = usersBucket.Get([]byte(userId))

		return nil
	}); err != nil {
		return nil, err
	}

	if userBytes == nil {
		return nil, nil
	}

	user := new(User)
	if err := json.Unmarshal(userBytes, user); err != nil {
		return nil, fmt.Errorf("error unmarshalling the user: %w", err)
	}

	return user, nil
}

func (storage *Storage) GetUserByUsername(username string) (*User, error) {
	var userBytes []byte

	if err := storage.db.View(func(tx *bolt.Tx) error {
		usernameToUserIdBucket := tx.Bucket([]byte(USERNAME_TO_USER_ID_BUCKET))

		userId := string(usernameToUserIdBucket.Get([]byte(username)))
		if userId == "" {
			return nil
		}

		usersBucket := tx.Bucket([]byte(USERS_BUCKET))

		userBytes = usersBucket.Get([]byte(userId))

		return nil
	}); err != nil {
		return nil, err
	}

	if userBytes == nil {
		return nil, nil
	}

	user := new(User)
	if err := json.Unmarshal(userBytes, user); err != nil {
		return nil, fmt.Errorf("error unmarshalling the user: %w", err)
	}

	return user, nil
}
