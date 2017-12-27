package controllers

import (
	"github.com/efrainmunoz/go-microservice-template/service/models"
)

type (
	// For Post = /user/register
	UserResource struct {
		Data models.User `json:"data"`
	}

	// For Post - /user/login
	LoginResource struct {
		Data LoginModel `json:"data"`
	}

	// Model for authentication
	LoginModel struct {
		User  models.User `json:"user"`
		Token string      `json:token`
	}
)
