package repository

import (
	"gakujo-notification/gakujo"
	"testing"
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
		Year: 2022,
	})
	assignments = append(assignments, &gakujo.Assignment{
		SubjectName: "test",
		Title:       "test assignment 2",
		Year: 2022,
	})
	if _, err := repo.UpsertAssignments(assignments); err != nil {
		t.Fatal(err)
	}
}
