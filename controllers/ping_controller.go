package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	"/ping" request is done by the AWS or any other cloud system in order to check if our program is live or not.
	so in it's response we need to pass "pong".
*/
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
