package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	"/ping" request is done by the AWS or any other system in ourder to check if our program is live or not.
	so it it's response we need to pass "200".
*/
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
