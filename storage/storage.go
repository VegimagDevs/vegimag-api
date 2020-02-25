package storage

import (
	"fmt"
	bolt "go.etcd.io/bbolt"
)

const (
	USERS_BUCKET               = "users"
	EMAIL_TO_USER_ID_BUCKET    = "emailToUserId"
	USERNAME_TO_USER_ID_BUCKET = "usernameToUserId"
	REFRESH_TOKENS_BUCKET      = "refreshTokens"
)

type Config struct {
	Path string
}

type Storage struct {
	config *Config

	db *bolt.DB
}

func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

func (storage *Storage) Open() error {
	var err error
	storage.db, err = bolt.Open(storage.config.Path, 0666, nil)
	if err != nil {
		return fmt.Errorf("can't open the database: %w", err)
	}

	if err := storage.initBuckets(); err != nil {
		return fmt.Errorf("bucket initialization failed: %w", err)
	}

	return nil
}

func (storage *Storage) Close() error {
	if err := storage.db.Close(); err != nil {
		return fmt.Errorf("can't close the database: %w", err)
	}

	return nil
}

func (storage *Storage) initBuckets() error {
	return storage.db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte(USERS_BUCKET)); err != nil {
			return err
		}

		if _, err := tx.CreateBucketIfNotExists([]byte(EMAIL_TO_USER_ID_BUCKET)); err != nil {
			return err
		}

		if _, err := tx.CreateBucketIfNotExists([]byte(USERNAME_TO_USER_ID_BUCKET)); err != nil {
			return err
		}

		return nil
	})
}
