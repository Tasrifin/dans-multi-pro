package services

import (
	"dans-multi-pro/models"
	"dans-multi-pro/params"
	"dans-multi-pro/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JobService struct {
	jobRepo repositories.JobRepo
}

func NewJobService(repo repositories.JobRepo) *JobService {
	return &JobService{
		jobRepo: repo,
	}
}

func (u *UserService) GetJobList(request params.CreateUser) *params.Response {
	user := models.User{
		Username: request.Username,
		Password: request.Password,
	}

	createUserData, err := u.userRepo.CreateUser(&user)

	if err != nil {
		return &params.Response{
			Status: http.StatusBadRequest,
			Payload: map[string]string{
				"error": err.Error(),
			},
		}
	}

	return &params.Response{
		Status: http.StatusCreated,
		Payload: gin.H{
			"id":       createUserData.ID,
			"username": createUserData.Username,
		},
	}
}
