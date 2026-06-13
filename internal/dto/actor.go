package dto

import (
	"net/http"

	"github.com/damir/jobfinder/internal/middleware"
	"github.com/damir/jobfinder/internal/model"
)

type Actor struct {
	UserID uint
	Role   model.Role
}

func GetActor(r *http.Request) *Actor {
	return &Actor{
		UserID: middleware.GetUserID(
			r.Context(),
		),
		Role: middleware.GetRole(r.Context()),
	}
}
