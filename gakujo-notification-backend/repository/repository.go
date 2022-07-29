package repository

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbname  = "gakujo-notification"
	port    = "5432"
	sslmode = "disable"
)

var (
	ErrSessionExpired = errors.New("session expired")
)

type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type Repository struct {
	db *gorm.DB
}

func New(host, user, password string) (*Repository, error) {
	mp := map[string]string{}
	mp["host"] = host
	mp["user"] = user
	mp["password"] = password
	mp["dbname"] = dbname
	mp["port"] = port
	mp["sslmode"] = sslmode
	dsn := joinMap(mp, "=", " ")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&Assignment{}, &UserAssignment{}, &User{}); err != nil {
		return nil, err
	}
	return &Repository{db}, nil
}

func joinMap(mp map[string]string, keyValSep, sliceSep string) string {
	tokens := []string{}
	for k, v := range mp {
		tokens = append(tokens, strings.Join([]string{k, v}, keyValSep))
	}
	return strings.Join(tokens, sliceSep)
}

func (r *Repository) RunInTransaction(f func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return r.db.Transaction(f, opts...)
}
