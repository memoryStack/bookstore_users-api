package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// will be called by main function to start the server.
func StartApplication() {
	mapUrls()
	router.run(":8080")
}

/*
	we are using gin-gonic HTTP framework. here we will start the http server and call the corresponding controller
	depending upon the path of the request.
	NOTE: the only place where we will be interacting with the HTTP framework is this file and the controllers package so that
	we don't have to make much changes if we decide to change the HTTP framework someday.
*/
