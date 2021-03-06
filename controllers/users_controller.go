package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/backendLearningProjects/bookstore_users-api/domain/users"
	"github.com/backendLearningProjects/bookstore_users-api/services"
	"github.com/backendLearningProjects/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

/*
	<----- some wisdom about why and what of this file ------>

	the "input layer" of our application will be this file. all the requests will be attended by this file.
	and the whole idea of a controller is to provide the endpoints and the functionalities for the interation
	with the USERS API.

	basically a controller takes a request and reads the data from the request and passes all the needed information
	to the services where all the data related business logic is going on. And all the controllers will have same kind
	of repeated logic in them.

*/

func GetUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		err := errors.BadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

func CreateUser(c *gin.Context) {
	var user users.User
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		restErr := errors.BadRequestError("error while reading request body")
		c.JSON(restErr.Status, restErr)
		return
	}

	// TODO: how do we know that this function actually expects address of the "user" variable by looking
	// at the inputs type ??
	// maps and slices are passed by reference in go. so this is the answer to previous question above
	if err := json.Unmarshal(bytes, &user); err != nil {
		restErr := errors.BadRequestError("error while unmarshalling request body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, createError := services.CreateUser(user)

	if createError != nil {
		// TODO: handle user create error
	}
	c.JSON(http.StatusCreated, result)
}
