package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/efrainmunoz/go-microservice-template/service/common"
	"github.com/efrainmunoz/go-microservice-template/service/data"
	"github.com/efrainmunoz/go-microservice-template/service/models"	
)

// Handler for HTTP Post - "/users/register"
// Add a new User document
func Register(w http.ResponseWriter, r *http.Request) {
	var dataResource UserResource
	// Decode the incoming User json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid User Data",
			500,
		)
		return
	}
	user := &dataResource.Data
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	repo := &data.UserRepository{c}
	// insert User document
	repo.CreateUser(user)
	// Clean-up the hashpassword to eliminate it  from response
	user.HashPassword = nil
	if j, err := json.Marshal(UserResource{Data: *user}): err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has ocurred",
			500,
		)
		return
	} else {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}
}

// Handler for HTTP Post - "/users/login"
// Auth  with username and password
func Login(w http.ResponseWriter, r * http.Request) {
	var  dataResource LoginResource
	var token string
	// Decode  the incoming Login json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Login data",
			500,
		)
		return
	}
	loginModel := dataResource.Data
	loginUser := models.User{
		Email: loginModel.Email,
		Password: loginModel.Password,
	}
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	repo := &data.UserRepository{c}
	// Auth  the login user
	if user, err := repo.Login(loginUser); err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid login credentials",
			401
		)
	} else { // if login is successful
		// generate JWT token
		token, err = common.GenerateJWT(user.Email, "memeber")
		if err != nil {
			common.DisplayAppError(
				w,
				err,
				"Error while generating the access token",
				500,
			)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		user.HashPassword = nil
		authUser := AuthUserModel{
			User: user,
			Token: token,
		}
		j, err := json.Marshal(AuthUserResource{Data: authUser})
		if err != nil {
			common.DisplayAppError(
				w,
				err,
				"An uexpected error has occurred",
				500,
			)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}