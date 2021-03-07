package services

import (
	"github.com/backendLearningProjects/bookstore_users-api/domain/users"
	"github.com/backendLearningProjects/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	return &user, nil
}

/*
	controllers responsibility is to just extract infos from the request and then pass over the required info to the service handlers.
	and then it's the responsibility to do all the business logic for returning the response for the request.
*/
