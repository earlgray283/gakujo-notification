package server

import (
	"gakujo-notification/gakujo"
	"gakujo-notification/repository"

	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type Assignment struct {
	repository.Assignment
	Status gakujo.AssignmentStatus `json:"status"`
}

func (srv *Server) HandleGetAllAssignments(c *fiber.Ctx) error {
	year, _ := c.ParamsInt("year", 2022)
	userId := getUserIdFromJwtToken(c.Locals("user").(*jwt.Token))
	userAssignments, err := srv.repo.FetchAllUserAssignments(userId, int(year))
	if err != nil {
		srv.logger.Println(err)
		return c.SendStatus(http.StatusInternalServerError)
	}
	assignments := make([]*Assignment, len(userAssignments))
	for i, userAssignment := range userAssignments {
		assignments[i] = &Assignment{
			Assignment: userAssignment.Assignment,
			Status:     userAssignment.Status,
		}
	}
	return c.Status(http.StatusOK).JSON(assignments)
}
