package data

import (
	"gopkg.in/mgo.v2"
	"TaskManager/models"
	"gopkg.in/mgo.v2/bson"
	"golang.org/x/crypto/bcrypt"


)

type UserRepository struct {
	C * mgo.Collection
}


func (r * UserRepository) CreateUser(user * models.User)error{
	obj_id := bson.NewObjectId()
	user.Id = obj_id

	hpass ,err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)

	if err != nil{
		panic(err)
	}

	user.HashPassword = hpass

	user.Password = ""
	err = r.C.Insert(&user)
	return err
}
func (r * UserRepository) Login(user models.User)(u models.User, err error){

	err = bcrypt.CompareHashAndPassword(u.HashPassword,[]byte(user.Password))

	if err != nil{
		u = models.User{}
	}
	return


}
