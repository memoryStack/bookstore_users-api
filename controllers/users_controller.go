package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	<----- some wisdom about why and what of this file ------>

	the "input layer" of our application will be this file. all the requests will be attended by this file.
	and the whole idea of a controller is to provide the endpoints and the functionalities for the interation
	with the USERS API.
	basically a controller takes a request and passes all the needed information to the services where all the data related business logic
	is going on.

*/

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
