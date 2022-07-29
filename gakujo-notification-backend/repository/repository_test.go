package repository

import (
	"gakujo-notification/gakujo"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestAssignments(t *testing.T) {
	repo, err := New("localhost", "root", "root")
	if err != nil {
		t.Fatal(err)
	}
	assignments := make([]*gakujo.Assignment, 0)
	assignments = append(assignments, &gakujo.Assignment{
		SubjectName: "test",
		Title:       "test assignment",
		Year:        2022,
	})
	assignments = append(assignments, &gakujo.Assignment{
		SubjectName: "test",
		Title:       "test assignment 2",
		Year:        2022,
	})
	repoAssignments, err := repo.UpsertAssignments(assignments)
	if err != nil {
		t.Fatal(err)
	}
	for _, asmt := range repoAssignments {
		t.Log(asmt.ID)
	}
	userAsmts := make([]*UserAssignment, 0)
	for _, asmt := range repoAssignments {
		userAsmts = append(userAsmts, &UserAssignment{
			UserID:       uuid.New().String(),
			AssignmentID: asmt.ID,
			Status:       gakujo.AssignmentStatusOpen,
			CreatedAt:    time.Now(),
		})
	}
	if _, err := repo.UpsertUserAssignments(userAsmts); err != nil {
		t.Log(err)
	}
}
