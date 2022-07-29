package repository

import (
	"gakujo-notification/gakujo"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserAssignment struct {
	UserID       uint
	User         User
	AssignmentID uint
	Assignment   Assignment
	Status       gakujo.AssignmentStatus
	Model
}

func (repo *Repository) FetchAllUserAssignments(userID uint, year int) ([]*UserAssignment, error) {
	userAssignments := make([]*UserAssignment, 0)
	if err := repo.db.
		Where("year = ?", year).
		Where("user_id = ?", userID).
		Find(&userAssignments).
		Error; err != nil {
		return nil, err
	}
	return userAssignments, nil
}

func (repo *Repository) UpsertUserAssignments(userAssignments []*UserAssignment) ([]*UserAssignment, error) {
	if err := repo.RunInTransaction(func(tx *gorm.DB) error {
		if err := tx.Clauses(clause.OnConflict{DoNothing: false}).Create(&userAssignments).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return userAssignments, nil
}
