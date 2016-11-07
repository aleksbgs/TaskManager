package controllers

import (
	"net/http"
	"encoding/json"
	"TaskManager/common"
	"TaskManager/models"
	"TaskManager/data"

)

func Register(w http.ResponseWriter,r * http.Request){

	var dataResource UserResource

	err := json.NewDecoder(r.Body).Decode(&dataResource)

	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid User data",
			500,
		)
		return
	}
	user := &dataResource.Data

	context := NewContext()
	defer context.Close()

	col:=context.DbCollection("users")
	repo := &data.UserRepository{C:col}

	repo.CreateUser(user)

	user.HashPassword = nil

        if j,err := json.Marshal(UserResource{Data:*user});err!=nil{
		common.DisplayAppError(
			w,
			err,
			"An expected error has occurred",
			500,
		)
		return
	}else {
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}



}
func Login(w http.ResponseWriter, r *http.Request) {
	var dataResource LoginResource
	var token string
	// Decode the incoming Login json
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
		Email:    loginModel.Email,
		Password: loginModel.Password,
	}
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("users")
	repo := &data.UserRepository{C: col}
	// Authenticate the login user
	user, err := repo.Login(loginUser)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid login credentials",
			401,
		)
		return
	}
	// Generate JWT token
	token, err = common.GenerateJWT(user.Email, "member")
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Eror while generating the access token",
			500,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// Clean-up the hashpassword to eliminate it from response JSON
	user.HashPassword = nil
	authUser := AuthUserModel{
		User:  user,
		Token: token,
	}
	j, err := json.Marshal(AuthUserResource{Data: authUser})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}