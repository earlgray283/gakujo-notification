package worker

import (
	"context"
	"fmt"
	"gakujo-notification/gakujo"
	"gakujo-notification/lib"
	"gakujo-notification/repository"
	"log"
	"strings"
	"time"
)

func CrawleAssignments(repo *repository.Repository, crypto *lib.Crypto) func() {
	return func() {
		users, err := repo.FetchAllUsers()
		if err != nil {
			log.Println(err)
			return
		}

		for _, user := range users {
			gakujoAccount, err := crypto.Decrypt(user.EncryptedGakujoAccount)
			if err != nil {
				log.Println(err)
				return
			}
			tokens := strings.Split(gakujoAccount, "&")
			id, password := tokens[0], tokens[1]

			client, err := gakujo.NewClient(context.Background(), id, password)
			if err != nil {
				log.Println(err)
				return
			}
			defer client.Cancel()

			assignments, err := client.ReportAssignments()
			if err != nil {
				log.Println(err)
				return
			}

			repoAssignments, err := repo.UpsertAssignments(assignments)
			if err != nil {
				log.Println(err)
				return
			}
			asmtKeys := lib.MapSlice(repoAssignments, func(a *repository.Assignment) string {
				return fmt.Sprintf("%s_%d", a.Title, a.Year)
			})
			asmtMap := lib.NewMapFromIter(asmtKeys, repoAssignments)

			userAssignments := make([]*repository.UserAssignment, len(repoAssignments))
			for i, assignment := range assignments {
				repoAsmt := asmtMap[fmt.Sprintf("%s_%d", assignment.Title, assignment.Year)]
				userAssignments[i] = &repository.UserAssignment{
					UserID:       user.ID,
					AssignmentID: repoAsmt.ID,
					Status:       assignment.Status,
					Model: repository.Model{
						CreatedAt: time.Now(),
					},
				}
			}
			if _, err := repo.UpsertUserAssignments(userAssignments); err != nil {
				log.Println(err)
				return
			}
		}
	}
}
