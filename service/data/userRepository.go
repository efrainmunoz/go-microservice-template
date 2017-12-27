package data

import (
	"github.com/efrainmunoz/go-microservice-template/service/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserRepository struct
type UserRepository struct {
	C *mgo.Collection
}

// CreateUser on mongo's users collection
func (r *UserRepository) CreateUser(user *models.User) error {
	objID := bson.NewObjectId()
	user.ID = objID
	hpass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.HashPassword = hpass
	// clear the incoming text password
	user.Password = ""
	err = r.C.Insert(&user)
	return err
}

// Login user to the system
func (r *UserRepository) Login(user models.User) (u models.User, err error) {
	err = r.C.Find(bson.M{"email": user.Email}).One(&u)
	if err != nil {
		return
	}
	// validate password
	err = bcrypt.CompareHashAndPassword(u.HashPassword, []byte(user.Password))
	if err != nil {
		u = models.User{}
	}
	return
}
