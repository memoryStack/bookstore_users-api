package app

import (
	"github.com/backendLearningProjects/bookstore_users-api/controllers"
)

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
