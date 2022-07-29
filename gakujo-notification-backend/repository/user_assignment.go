package repository

import (
	"gakujo-notification/gakujo"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserAssignment struct {
	ID           string
	UserID       string
	User         User
	AssignmentID string
	Assignment   Assignment
	Status       gakujo.AssignmentStatus
	CreatedAt    time.Time
	UpdateAt     time.Time
}

func (repo *Repository) FetchAllUserAssignments(userID string, year int) ([]*UserAssignment, error) {
	userAssignments := make([]*UserAssignment, 0)
	if err := repo.db.
		Joins("Assignment").
		Where("user_assignments.user_id = ?", userID).
		Find(&userAssignments).
		Error; err != nil {
		return nil, err
	}
	return userAssignments, nil
}

func (repo *Repository) UpsertUserAssignments(userAssignments []*UserAssignment) ([]*UserAssignment, error) {
	if err := repo.RunInTransaction(func(tx *gorm.DB) error {
		if err := tx.Clauses(clause.OnConflict{
			DoNothing: true,
		}).Create(&userAssignments).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return userAssignments, nil
}
