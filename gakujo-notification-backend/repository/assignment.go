package repository

import (
	"gakujo-notification/gakujo"
	"time"

	"gorm.io/gorm/clause"
)

type Assignment struct {
	Kind        gakujo.AssignmentKind `json:"kind"`
	SubjectName string                `json:"subjectName" gorm:"index:idx_asmt,unique"`
	Semester    gakujo.Semester       `json:"semester"`
	Title       string                `json:"title" gorm:"index:idx_asmt,unique"`
	Since       time.Time             `json:"since"`
	Until       time.Time             `json:"until"`
	Description string                `json:"description"`
	Message     string                `json:"message"`
	Year        int                   `json:"year" gorm:"index:idx_asmt,unique"`

	Model
}

func (repo *Repository) FetchAssignments(year int, subjectNames ...string) ([]*Assignment, error) {
	assignments := make([]*Assignment, len(subjectNames))
	if err := repo.db.
		Where("subject_name = ?", subjectNames).
		Where("year = ?", year).
		Find(&assignments).Error; err != nil {
		return nil, err
	}
	return assignments, nil
}

func (repo *Repository) UpsertAssignments(assignments []*gakujo.Assignment) ([]*Assignment, error) {
	repoAssignments := make([]*Assignment, len(assignments))
	for i, assignment := range assignments {
		repoAssignments[i] = &Assignment{
			Kind:        assignment.Kind,
			SubjectName: assignment.SubjectName,
			Semester:    assignment.Semester,
			Title:       assignment.Title,
			Since:       assignment.Since,
			Until:       assignment.Until,
			Description: assignment.Description,
			Message:     assignment.Message,
			Year:        assignment.Year,
			Model: Model{
				CreatedAt: time.Now(),
			},
		}
	}
	if err := repo.db.
		Clauses(clause.OnConflict{
			DoNothing: true,
		}).
		Create(&repoAssignments).
		Error; err != nil {
		return nil, err
	}
	return repoAssignments, nil
}
