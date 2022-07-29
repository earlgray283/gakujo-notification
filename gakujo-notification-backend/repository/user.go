package repository

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                     string
	Username               string `json:"username" gorm:"unique"`
	EncryptedPassword      string `json:"-"`
	EncryptedGakujoAccount string `json:"-"` // {id}&{password}
	CreatedAt              time.Time
	UpdatedAt              time.Time
}

func NewUser(username, encryptedPassword, encryptedGakujoAccount string) *User {
	return &User{
		ID:                     uuid.New().String(),
		Username:               username,
		EncryptedPassword:      encryptedPassword,
		EncryptedGakujoAccount: encryptedGakujoAccount,
		CreatedAt:              time.Now(),
	}
}

func (repo *Repository) FetchUserByUsername(username string) (*User, error) {
	var user User
	if err := repo.db.
		Where("username = ?", username).
		Find(&user).
		Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *Repository) FetchAllUsers() ([]*User, error) {
	users := make([]*User, 0)
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
